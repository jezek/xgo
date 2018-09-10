package xgo

import (
	"fmt"
	"log"

	"github.com/jezek/xgb/xproto"
)

// Font allocated on display
type Font struct {
	xproto.Font
	d       *Display
	pattern string

	fi *FontInfo
}

func (f *Font) Display() *Display {
	if f.d == nil {
		log.Fatalf("Font %v has no display", f)
	}
	return f.d
}

// Closes font allocation in X. Font has not be used anymore after close.
func (f *Font) Close() error {
	if f.Font != xproto.FontNone {
		if err := xproto.CloseFontChecked(f.Display().Conn, f.Font).Check(); err != nil {
			return errWrap{fmt.Sprintf("unable to close font %v", f), err}
		}
		f.Font = xproto.FontNone
	}
	return nil
}

// Opens font with pattern on provided display.
// If you don't want to use the font anymore, use Close method on it.
// On any error, the error is returned
func OpenFontOnDisplay(d *Display, pattern string) (*Font, error) {
	//get font id, open font name to id, use foont id, close font id
	fontId, err := xproto.NewFontId(d.Conn)
	if err != nil {
		return nil, errWrap{"OpenFontOnDisplay", errWrap{"unable to get font id from X", err}}
	}

	if err := xproto.OpenFontChecked(
		d.Conn,
		fontId,
		uint16(len(pattern)),
		pattern,
	).Check(); err != nil {
		if closeErr := xproto.CloseFontChecked(d.Conn, fontId).Check(); closeErr != nil {
			err = errWrap{fmt.Sprintf("additional error: \"%v\", original error", errWrap{fmt.Sprintf("unable to close font #%d", fontId), closeErr}), err}
		}
		return nil, errWrap{"OpenFontOnDisplay", errWrap{fmt.Sprintf("unable to bind font #%d to font with pattern \"%s\"", fontId, pattern), err}}

	}

	return &Font{fontId, d, pattern, nil}, nil
}

func (f *Font) Info() (*FontInfo, error) {
	if f.fi != nil {
		return f.fi, nil
	}

	fi, err := FontQuery(f)
	if err != nil {
		err = errWrap{"Font.Info", err}
	}
	return fi, err
}

func (f *Font) Fontable() xproto.Fontable {
	return xproto.Fontable(f.Font)
}

//TODO integrate
/*
	// load font from x11 and draw text to our pixmap
	if fonts, err := xproto.ListFonts(d.Conn, 10, uint16(len(fontName)), fontName).Reply(); err != nil {
		fmt.Println("List fonts error:", err)
		return
	} else {
		fmt.Printf("fonts: %#v\n", fonts)
		if fonts.NamesLen == 0 {
			fmt.Printf("no font found\n")
			//return
		}
	}

	//TODO not working, just waiting for something
	//if fonts, err := xproto.ListFontsWithInfo(d.Conn, 1, 6, "*mono*").Reply(); err != nil {
	//	fmt.Println("List fonts with info error:", err)
	//	return
	//} else {
	//	fmt.Printf("fonts with info: %#v\n", fonts)
	//}

	if fp, err := xproto.GetFontPath(d.Conn).Reply(); err != nil {
		fmt.Println("unable to GetFontPath:", err)
		return
	} else {
		fmt.Printf("GetFontPath reply: %#v\n", fp)
	}
*/

type Fontable interface {
	Fontable() xproto.Fontable
	Display() *Display
}

//TODO timestamps for all X replies
type FontInfo struct {
	*xproto.QueryFontReply
	f Fontable
}

func (fi *FontInfo) Fontable() Fontable {
	if fi.f == nil {
		log.Fatalf("Font info %v has no fontable", fi)
	}
	return fi.f
}

func (fi *FontInfo) Display() *Display {
	return fi.Fontable().Display()
}

func (fi *FontInfo) Properties() map[string]string {
	//TODO cache
	res := map[string]string{}
	for _, prop := range fi.QueryFontReply.Properties {
		res[fi.Display().Atom(prop.Name)] = fi.Display().Atom(xproto.Atom(prop.Value))
	}
	return res
}

// Returns font information for Fontable structures, or error if some troubles.
// Current Fontables are: *Font, *GraphicsContext
func FontQuery(f Fontable) (*FontInfo, error) {
	reply, err := xproto.QueryFont(
		f.Display().Conn,
		f.Fontable(),
	).Reply()
	if err != nil {
		return nil, errWrap{"FontQuery", errWrap{fmt.Sprintf("unable to query for fontable %v", f), err}}
	}
	return &FontInfo{reply, f}, nil
}
