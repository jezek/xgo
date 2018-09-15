package xgo

import (
	"fmt"
	"log"
	"sort"

	"github.com/jezek/xgb/xproto"
)

type Drawable interface {
	Drawable() xproto.Drawable
}

type GraphicsContext struct {
	xproto.Gcontext
	d        *Display
	drawable Drawable

	fontCloseOnFree *Font
}

type componentInfo struct {
	componentNo int
	mask, value uint32
}

func (gc *GraphicsContext) Display() *Display {
	if gc.d == nil {
		log.Fatalf("GraphicsContext %v has no display", gc)
	}
	return gc.d
}

func (gc *GraphicsContext) TextExtents(text string) (*xproto.QueryTextExtentsReply, error) {
	return TextExtents(text, gc)
}

// Create and allocate graphics context on display with drawable root and depth using provided components
// Returned graphical context can be used with any destination drawable having the same root and depthe.
// Use with other drawables results in a BadMatch error.
func NewGraphicsContextOnDisplay(d *Display, drawable Drawable, components ...GraphicsContextComponent) (resultGc *GraphicsContext, resultErr error) {
	resultGc = &GraphicsContext{0, d, drawable, nil}

	defer func() {
		// clean up some potentialy created font if some error is returned
		if resultErr != nil {
			if resultGc != nil && resultGc.fontCloseOnFree != nil {
				if errClose := resultGc.fontCloseOnFree.Close(); errClose != nil {
					resultErr = errWrap{fmt.Sprintf("additional error: \"NewGraphicsContextOnDisplay: unable to free font %v: %v\", original error", resultGc.fontCloseOnFree, errClose), resultErr}
				} else {
					resultGc.fontCloseOnFree = nil
				}
			}
			resultGc = nil
		}
	}()
	// get all graphics context components to one array for further sorting and error checking
	componentsInfo := []componentInfo{}

	for i, component := range components {
		mvMap, err := component(resultGc)
		if err != nil {
			return resultGc, errWrap{"NewGraphicsContextOnDisplay", errWrap{fmt.Sprintf("error getting mask:value map from component no %d", i), err}}
		}

		for cMask, cValue := range mvMap {
			componentsInfo = append(componentsInfo, componentInfo{i, cMask, cValue})
		}
	}

	// sort componentsInfo by mask ascending
	sort.SliceStable(componentsInfo, func(i, j int) bool {
		return componentsInfo[i].mask < componentsInfo[j].mask
	})

	mask := uint32(0)
	values := make([]uint32, 0)

	for _, c := range componentsInfo {
		if mask|c.mask == mask {
			return resultGc, errWrap{"NewGraphicsContextOnDisplay", fmt.Errorf("component %d returned duplicate mask: %d", c.componentNo, c.mask)}
		}
		mask = mask | c.mask
		values = append(values, c.value)
	}

	// use screen pexels and selected font
	gc, err := xproto.NewGcontextId(d.Conn)
	if err != nil {
		return resultGc, errWrap{"NewGraphicsContextOnDisplay", errWrap{"unable to allocate graphics context id", err}}
	}

	if err := xproto.CreateGCChecked(
		d.Conn,
		gc,
		drawable.Drawable(),
		mask,
		values,
	).Check(); err != nil {

		return resultGc, errWrap{"NewGraphicsContextOnDisplay", errWrap{"unable to create graphic context", err}}
	}

	resultGc.Gcontext = gc

	return resultGc, nil
}

func (gc *GraphicsContext) Free() error {
	if gc.Gcontext == 0 {
		return nil
	}

	if gc.fontCloseOnFree != nil {
		if err := gc.fontCloseOnFree.Close(); err != nil {
			return errWrap{"GraphicsContext.Free", errWrap{fmt.Sprintf("unable to close graphics contexts opened font %v", gc.fontCloseOnFree), err}}
		}
		gc.fontCloseOnFree = nil
	}

	if err := xproto.FreeGCChecked(
		gc.Display().Conn,
		gc.Gcontext,
	).Check(); err != nil {
		return errWrap{"GraphicsContext.Free", errWrap{fmt.Sprintf("unable to free graphics context %v", gc), err}}
	}
	gc.Gcontext = 0
	return nil
}

func (gc *GraphicsContext) Fontable() xproto.Fontable {
	return xproto.Fontable(gc.Gcontext)
}

func (gc *GraphicsContext) FontInfo() (*FontInfo, error) {
	fi, err := FontQuery(gc)
	if err != nil {
		err = errWrap{"GraphicsContext.FontInfo", err}
	}
	return fi, err
}

type GraphicsContextComponent func(gc *GraphicsContext) (map[uint32]uint32, error)

type GraphicsContextComponents struct{}

func (_ GraphicsContextComponents) BackgroundPixel(val uint32) GraphicsContextComponent {
	return func(_ *GraphicsContext) (map[uint32]uint32, error) {
		return map[uint32]uint32{xproto.GcBackground: val}, nil
	}
}

func (_ GraphicsContextComponents) ForegroundPixel(val uint32) GraphicsContextComponent {
	return func(_ *GraphicsContext) (map[uint32]uint32, error) {
		return map[uint32]uint32{xproto.GcForeground: val}, nil
	}
}

// Opens first font matching pattern, if there is any and assigns it to graphics context,
// If font does not exist, previous is kept, or if there is no previous, default is used.
// if used in graphics context creation or editing. Font is closed upon graphics context freeing,
// or graphics contexts font beeing updated (eg. with this component again).
func (_ GraphicsContextComponents) NewFontIfMatch(pattern string) GraphicsContextComponent {
	return func(gc *GraphicsContext) (map[uint32]uint32, error) {
		font, err := gc.Display().FontOpen(pattern)
		if err != nil {
			return map[uint32]uint32{}, nil
		}

		if gc.fontCloseOnFree != nil {
			if err := gc.fontCloseOnFree.Close(); err != nil {
				return nil, errWrap{"GraphicsContextComponents.Font", errWrap{"GraphicsContextComponent", errWrap{fmt.Sprintf("unable to close graphics contexts opened font %v before assigning new", gc.fontCloseOnFree), err}}}
			}
		}

		gc.fontCloseOnFree = font
		return map[uint32]uint32{xproto.GcFont: uint32(font.Font)}, nil
	}
}

// Opens first font matching pattern an assigns to graphics context,
// if used in graphics context creation or editing. Font is closed upon graphics context freeing,
// or graphics contexts font beeing updated (eg. with this component again).
func (_ GraphicsContextComponents) NewFont(pattern string) GraphicsContextComponent {
	return func(gc *GraphicsContext) (map[uint32]uint32, error) {
		font, err := gc.Display().FontOpen(pattern)
		if err != nil {
			return nil, errWrap{"GraphicsContextComponents.NewFont", errWrap{"GraphicsContextComponent", err}}
		}

		if gc.fontCloseOnFree != nil {
			if err := gc.fontCloseOnFree.Close(); err != nil {
				return nil, errWrap{"GraphicsContextComponents.Font", errWrap{"GraphicsContextComponent", errWrap{fmt.Sprintf("unable to close graphics contexts opened font %v before assigning new", gc.fontCloseOnFree), err}}}
			}
		}

		gc.fontCloseOnFree = font
		return map[uint32]uint32{xproto.GcFont: uint32(font.Font)}, nil
	}
}

func (_ GraphicsContextComponents) Font(font *Font) GraphicsContextComponent {
	return func(gc *GraphicsContext) (map[uint32]uint32, error) {
		if gc.fontCloseOnFree != nil {
			if err := gc.fontCloseOnFree.Close(); err != nil {
				return nil, errWrap{"GraphicsContextComponents.Font", errWrap{"GraphicsContextComponent", errWrap{fmt.Sprintf("unable to close graphics contexts opened font %v before assigning new", gc.fontCloseOnFree), err}}}
			}
		}
		gc.fontCloseOnFree = nil
		return map[uint32]uint32{xproto.GcFont: uint32(font.Font)}, nil
	}
}

//TODO GcFunction, GcPlaneMask, GcLineWidth, GcLineStyle, GcCapStyle, GcJoinStyle, GcFillStyle, GcFillRule, GcTile, GcStipple, GcTileStippleOriginX, GcTileStippleOriginY, GcSubwindowMode, GcGraphicsExposures, GcClipOriginX, GcClipOriginY, GcClipMask, GcDashOffset, GcDashList, GcArcMode,
