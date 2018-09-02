package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"os/signal"
	"unicode/utf16"
	"xgo"

	"github.com/jezek/xgb/xproto"
)

func main() {

	// load image
	imageFileName := "pokeslon.jpg"
	//imageFileName := "jez.jpg"
	reader, err := os.Open(imageFileName)
	if err != nil {
		fmt.Printf("Unable to open \"%s\": %v\n", imageFileName, err)
		return
	}
	defer reader.Close()
	img, format, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Unable to decode image \"%s\": %v\n", imageFileName, err)
		return
	}
	fmt.Printf("Loaded image \"%s\" as an %dx%d %s format\n", imageFileName, img.Bounds().Dx(), img.Bounds().Dy(), format)
	bounds := img.Bounds()

	// open display of x11
	d, err := xgo.OpenDisplay("")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Display opened: %#v\n", d)

	// create window
	w, err := d.CreateWindow()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Window created: %#v\n", w)
	defer w.Destroy()

	// change window size & background color & center
	screenBounds := image.Rect(0, 0, int(w.Screen().WidthInPixels), int(w.Screen().HeightInPixels))
	winBounds := screenBounds.Intersect(bounds)                                                                     // get win size
	winBounds = winBounds.Add(image.Pt((screenBounds.Dx()-winBounds.Dx())/2, (screenBounds.Dy()-winBounds.Dy())/2)) // center win
	if err := w.MoveResize(winBounds); err != nil {
		fmt.Printf("Error moving & resizing window: %v\n", err)
		return
	}
	fmt.Println("Window moved and resized to", winBounds)

	if err := w.BgColorChange(color.RGBA{}); err != nil {
		fmt.Printf("Error changing window background color: %v\n", err)
		return
	}
	fmt.Println("Window bg color changed")
	//TODO shrink image size to fit if needed

	//TODO? Set WM_STATE so it is interpreted as a top-level window.
	//TODO? Set WM_NORMAL_HINTS so the window can't be resized.
	//TODO? Set _NET_WM_NAME so it looks nice.

	// convert image to BGRA order, cause x11 works this way
	pixBGRA := make([]byte, 0, bounds.Dx()*bounds.Dy())
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			// r is uint32, but color inside is from 0 to 0xffff (16 bit), so for a 8 bit repesentation just shift right by 8 bits
			pixBGRA = append(pixBGRA, byte(b>>8), byte(g>>8), byte(r>>8), byte(a>>8))
		}
	}
	fmt.Println("image converted to BGRA bytes", pixBGRA[:60], "...")
	// allocate new pixmap in x
	// first get new pixmap id
	pixmapId, err := xproto.NewPixmapId(d.Conn)
	if err != nil {
		fmt.Printf("Unable to create pixmap id: %v\n", err)
		return
	}
	// then send create request for pixmap
	if err := xproto.CreatePixmapChecked(
		d.Conn,
		w.Screen().RootDepth,
		pixmapId,
		xproto.Drawable(w.Screen().Root),
		uint16(bounds.Dx()),
		uint16(bounds.Dy()),
	).Check(); err != nil {
		fmt.Printf("Unable to create pixmap: %v\n", err)
		return
	}
	fmt.Println("pixmap created")

	{ // draw our image to pixmap
		// create and allocate graphical content
		gc, err := xproto.NewGcontextId(d.Conn)
		if err != nil {
			fmt.Printf("Unable to allocate graphic context id: %v\n", err)
			return
		}
		if err := xproto.CreateGCChecked(
			d.Conn,
			gc,
			xproto.Drawable(w.Screen().Root),
			xproto.GcForeground,
			[]uint32{w.Screen().WhitePixel},
		).Check(); err != nil {
			fmt.Printf("Unable to create graphic context: %v\n", err)
			return
		}
		fmt.Println("gc created")

		maxBytesToSend := xgo.SizeReqestMax - (xgo.SizeRequestPutImageFixedPart + 4)
		maxAreaToSend := maxBytesToSend / 4
		if maxAreaToSend*4 != maxBytesToSend {
			fmt.Printf("maxAreaToSend(%d) * 4 (%d) != maxBytesToSend (%d)\n", maxAreaToSend, maxAreaToSend*4, maxBytesToSend)
			maxBytesToSend = maxAreaToSend * 4
		}
		rectangles := decomposeToRectanglesWithAreaMax(bounds, maxAreaToSend)

		for _, rectangle := range rectangles {
			// draw our image stored in pixBGRA to pixmap (pid) using gc
			if err := xproto.PutImageChecked(
				d.Conn,
				xproto.ImageFormatZPixmap,
				xproto.Drawable(pixmapId),
				gc,
				uint16(rectangle.Dx()),
				uint16(rectangle.Dy()),
				int16(rectangle.Min.X),
				int16(rectangle.Min.Y),
				0,  // left padding
				24, //TODO get from win. depth
				getRectangleBytes(rectangle, pixBGRA),
			).Check(); err != nil {
				fmt.Printf("Unable to write image bgra pixels to pixmap: %v\n", err)
				return
			}
			fmt.Printf("pixBGRA data for image rectangle %v inserted into drawable pixmap, using gc\n", rectangle)
		}
	}

	fontName := "*fixed*-20-*"
	//fontName := "-Misc-Fixed-Medium-R-SemiCondensed--13-120-75-75-C-60-ISO8859-1"

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

	{ // draw something to pixmap

		//if fp, err := xproto.GetFontPath(d.Conn).Reply(); err != nil {
		//	fmt.Println("unable to GetFontPath:", err)
		//	return
		//} else {
		//	fmt.Printf("GetFontPath reply: %#v\n", fp)
		//}

		//get font id, open font name to id, use foont id, close font id
		fontId, err := xproto.NewFontId(d.Conn)
		if err != nil {
			fmt.Printf("Unable to get font id: %v\n", err)
			return
		}
		fmt.Printf("Got new font id: %v\n", fontId)

		if err := xproto.OpenFontChecked(
			d.Conn,
			fontId,
			uint16(len(fontName)),
			fontName,
		).Check(); err != nil {
			fmt.Printf("Unable to bind font id: %v to font name \"%s\": %v\n", fontId, fontName, err)
			return
		}

		defer func() {
			if err := xproto.CloseFontChecked(d.Conn, fontId).Check(); err != nil {
				fmt.Printf("Unable to close font id: %v for cleanup: %v\n", fontId, err)
				return
			}
			fmt.Printf("Closed font id: %v for cleanup\n", fontId)
		}()

		// create and allocate graphical content for text write
		// use screen pexels and selected font
		gc, err := xproto.NewGcontextId(d.Conn)
		if err != nil {
			fmt.Printf("Unable to allocate graphic context id for text write: %v\n", err)
			return
		}
		if err := xproto.CreateGCChecked(
			d.Conn,
			gc,
			xproto.Drawable(w.Screen().Root),
			xproto.GcForeground|xproto.GcBackground|xproto.GcFont,
			[]uint32{w.Screen().WhitePixel, w.Screen().BlackPixel, uint32(fontId)},
			//xproto.GcForeground|xproto.GcBackground,
			//[]uint32{w.Screen().WhitePixel, w.Screen().BlackPixel},
		).Check(); err != nil {
			fmt.Printf("Unable to create graphic context for text write: %v\n", err)
			return
		}
		fmt.Println("gc for text write created")

		fry, err := xproto.QueryFont(
			d.Conn,
			xproto.Fontable(gc),
		).Reply()
		if err != nil {
			fmt.Println("unable to query for font:", err)
			return
		}
		//fmt.Printf("Font reply: %#v\n", fry)
		fmt.Println("Font reply properties:")
		for _, prop := range fry.Properties {
			fmt.Println("\t", d.Atom(prop.Name), ":", d.Atom(xproto.Atom(prop.Value)))
		}

		imageText := imageFileName

		uint16String := utf16.Encode([]rune(imageText))
		c2bString := make([]xproto.Char2b, len(uint16String))
		for i, v := range uint16String {
			c2bString[i].Byte1 = byte(v >> 8)
			c2bString[i].Byte2 = byte(v)
		}
		ter, err := xproto.QueryTextExtents(
			d.Conn,
			xproto.Fontable(fontId),
			c2bString,
			uint16(len(c2bString)),
		).Reply()

		if err := xproto.ImageText8Checked(
			d.Conn,
			byte(len(imageText)),
			xproto.Drawable(pixmapId),
			gc,
			(int16(bounds.Dx())-int16(ter.OverallWidth))/2,
			fry.FontAscent,
			imageText,
		).Check(); err != nil {
			fmt.Printf("Unable to write text to pixmap: %v\n", err)
			return
		}
		fmt.Println("Wrote some text to pixmap")

		if err := xproto.ImageText16Checked(
			d.Conn,
			byte(len(c2bString)),
			xproto.Drawable(pixmapId),
			gc,
			int16(int32(bounds.Dx())-ter.OverallWidth)/2,
			int16(bounds.Dy())-fry.FontDescent,
			c2bString,
		).Check(); err != nil {
			fmt.Printf("Unable to write text to pixmap: %v\n", err)
			return
		}
		fmt.Println("Wrote some text to pixmap")

		if err := xproto.PolyLineChecked(
			d.Conn,
			xproto.CoordModeOrigin,
			xproto.Drawable(pixmapId),
			gc,
			[]xproto.Point{xproto.Point{0, 0}, xproto.Point{100, 100}, xproto.Point{0, 100}},
		).Check(); err != nil {
			fmt.Printf("Unable draw line: %v\n", err)
			return
		}
		fmt.Println("line draw")

	}

	//tell the window, to use our pixmap as background
	if err := xproto.ChangeWindowAttributesChecked(
		w.Screen().Display().Conn,
		w.Window,
		xproto.CwBackPixmap,
		[]uint32{uint32(pixmapId)},
	).Check(); err != nil {
		fmt.Printf("Unable to change window attribute to use pixmap as background: %v\n", err)
		return
	}
	fmt.Printf("Window attribute to have pixmap as background changed\n")

	// to show what has been drawn, we need to switch/flush the buffer
	if err := xproto.ClearAreaChecked(
		d.Conn,
		false,
		w.Window,
		0, 0, 0, 0,
	).Check(); err != nil {
		fmt.Printf("Unable to clear window area: %v\n", err)
		return
	}
	fmt.Printf("window cleared\n")

	if err := w.Map(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Window mapped")

	//TODO? do some helper funcion for this
	// siplified wait for window close request or interupt signal notify
	signalNotify := make(chan os.Signal, 1)
	signal.Notify(signalNotify, os.Interrupt, os.Kill)

	stopCloseNotify := make(chan struct{})
	closeRequest, err := w.CloseNotify(stopCloseNotify)
	if err != nil {
		fmt.Printf("Unable to monitor window for close notify due to error: %v\n", err)
		//TODO an example, how to use xgo.Display.Conn.WaitForEvent() if xgo.Window.CloseNotify returns error
		return
	}
	defer close(stopCloseNotify)

	select {
	case <-signalNotify:
	case <-closeRequest:
	}

}

func decomposeToRectanglesWithAreaMax(rect image.Rectangle, maxArea int) []image.Rectangle {
	if rect.Dx()*rect.Dy() <= maxArea {
		return []image.Rectangle{rect}
	}
	// area is greater than maxArea, so we are splitting
	//TODO? optimize... but how? use cycles instead of recursion

	// split verticaly, if deeded
	if rect.Dx() > maxArea {
		return append(
			decomposeToRectanglesWithAreaMax(
				image.Rect(
					rect.Min.X, rect.Min.Y,
					rect.Min.X+maxArea, rect.Max.Y,
				),
				maxArea,
			),
			decomposeToRectanglesWithAreaMax(
				image.Rect(
					rect.Min.X+maxArea, rect.Min.Y,
					rect.Max.X, rect.Max.Y,
				),
				maxArea,
			)...,
		)
	}
	// rect.Dx() < maxArea
	// split horizontaly by rows
	maxRows := maxArea / rect.Dx()
	return append(
		[]image.Rectangle{
			image.Rect(
				rect.Min.X, rect.Min.Y,
				rect.Max.X, rect.Min.Y+maxRows,
			),
		},
		decomposeToRectanglesWithAreaMax(
			image.Rect(
				rect.Min.X, rect.Min.Y+maxRows,
				rect.Max.X, rect.Max.Y,
			),
			maxArea,
		)...,
	)

}

func getRectangleBytes(rect image.Rectangle, quadrupletBytes []byte) []byte {
	//fmt.Println("geting rextangle %v bytes\n", rect)
	res := make([]byte, 0, rect.Dx()*rect.Dy()*4)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			pos := (rect.Dx() * y * 4) + x*4
			//fmt.Println("pos for x: %d, y: %d is %d\n", x, y, pos)
			res = append(res, quadrupletBytes[pos:pos+4]...)
		}
	}
	return res
}
