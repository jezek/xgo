package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"os/signal"
	"xgo"
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

	// open display of x11
	d, err := xgo.OpenDisplay("")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Display opened: %#v\n", d)

	screen := d.DefaultScreen()

	// get screen bounds
	screenBounds := image.Rect(0, 0, int(screen.WidthInPixels), int(screen.HeightInPixels))
	imageBounds := img.Bounds()
	//TODO? is it ok, to get the screen bounds? exists there a methot, that can rovide us with max window size?
	// get window size
	winBounds := screenBounds.Intersect(imageBounds)
	// TODO? what about window borders? do they count?
	if !winBounds.In(screenBounds) {
		//TODO shrink image size to fit if needed
		fmt.Printf("Image to large for screen\n")
		return
	}

	// center window bounds in screen bounds
	winBounds = winBounds.Add(image.Pt(
		(screenBounds.Dx()-winBounds.Dx())/2,
		(screenBounds.Dy()-winBounds.Dy())/2,
	))

	pixmap, err := screen.NewPixmap(image.Pt(imageBounds.Dx(), imageBounds.Dy()))
	if err != nil {
		fmt.Printf("Unable to create pixmap: %v\n", err)
		return
	}
	fmt.Printf("Pixmap created: %v\n", pixmap)

	defer func() {
		if err := pixmap.Destroy(); err != nil {
			fmt.Printf("Freeing pixmap error: %v\n", err)
			return
		}
		fmt.Printf("Pixmap destroyed: %v\n", pixmap)
	}()

	if err := pixmap.DrawImage(img); err != nil {
		fmt.Printf("Error drawing image to pixmap: %v\n", err)
		return
	}

	// create window
	w, err := screen.NewWindow(
		xgo.WindowOperations{}.Size(image.Pt(winBounds.Dx(), winBounds.Dy())),
		xgo.WindowOperations{}.Attributes(
			xgo.BackgroundPixmap(pixmap.Pixmap),
		),
		xgo.WindowOperations{}.Clear(),
		xgo.WindowOperations{}.Map(),
	)
	if err != nil {
		fmt.Printf("Default screen window creation error: %v\n", err)
		return
	}
	fmt.Printf("Window created: %#v\n", w)
	defer func() {
		if err := w.Destroy(); err != nil {
			fmt.Printf("Window destroy error: %v\n", err)
			return
		}
		fmt.Printf("Window destroyed: %v\n", w)
	}()

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

/*

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
			(int16(imageBounds.Dx())-int16(ter.OverallWidth))/2,
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
			int16(int32(imageBounds.Dx())-ter.OverallWidth)/2,
			int16(imageBounds.Dy())-fry.FontDescent,
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

*/
