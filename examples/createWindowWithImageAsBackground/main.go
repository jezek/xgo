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

var (
	gcc xgo.GraphicsContextComponents
	wo  xgo.WindowOperations
	wa  xgo.WindowAttributes
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

	//TODO with font info, so i don't have to fetch new and check for error
	textGc, err := screen.NewGraphicsContext(
		gcc.BackgroundPixel(screen.WhitePixel),
		gcc.ForegroundPixel(screen.BlackPixel),
		//gcc.Font(font),
		gcc.NewFont("*fixed*-20-*"),
		//TODO gcc.NewFontOptional(fontPattern),
		//TODO gcc.FontInfoQuery
	)
	if err != nil {
		fmt.Printf("Error creating graphics context for image text: %v\n", err)
		return
	}
	fmt.Printf("Graphics context for image text crated: %v\n", textGc)
	defer func() {
		if err := textGc.Free(); err != nil {
			fmt.Printf("Freeing graphics context for image text error: %v\n", err)
			return
		}
		fmt.Printf("Graphics context for image text freed: %v\n", textGc)
	}()

	textGcFontInfo, err := textGc.FontInfo()
	if err != nil {
		fmt.Printf("Error getting font informations from graphics context for image text \"%v\": %v\n", textGc, err)
		return
	}
	fmt.Printf("Font info properties from graphics context for image text: %v\n", textGcFontInfo.Properties())

	// get screen bounds
	screenBounds := image.Rect(0, 0, int(screen.WidthInPixels), int(screen.HeightInPixels))
	pixmapSize := image.Pt(img.Bounds().Dx(), img.Bounds().Dy()+int(textGcFontInfo.FontAscent)+int(textGcFontInfo.FontDescent))
	//TODO center text
	pixmapTextPosition := image.Pt(0, img.Bounds().Dy()+int(textGcFontInfo.FontAscent))

	//TODO? is it ok, to get the screen bounds? exists there a methot, that can rovide us with max window size?

	// get window size to fit image and text
	winBounds := screenBounds.Intersect(image.Rectangle{image.Pt(0, 0), pixmapSize})
	// TODO? what about window borders? do they count?
	if !winBounds.In(screenBounds) {
		fmt.Printf("Image to large for screen\n")
		return
	}

	// center window bounds in screen bounds
	winBounds = winBounds.Add(image.Pt(
		(screenBounds.Dx()-winBounds.Dx())/2,
		(screenBounds.Dy()-winBounds.Dy())/2,
	))

	pixmap, err := screen.NewPixmap(
		pixmapSize,
		xgo.PixmapOperations{}.Draw(
			xgo.PixmapDrawers{}.Image(img),
			xgo.PixmapDrawers{}.Text(imageFileName, pixmapTextPosition, textGc),
		),
	)
	if err != nil {
		fmt.Printf("Unable to create pixmap from image: %v\n", err)
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

	// create window
	w, err := screen.NewWindow(
		wo.Size(image.Pt(winBounds.Dx(), winBounds.Dy())),
		wo.Attributes(
			wa.BackgroundPixmap(pixmap),
		),
		wo.Clear(),
		wo.Map(),
	)
	//TODO make window NOT resizeable
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

//TODO use this in another example
/*
	fontPattern := "*fixed*-20-*"
	font, err := d.FontOpen(fontPattern)
	if err != nil {
		fmt.Printf("Error getting font for pattern \"%s\": %v\n", fontPattern, err)
		return
	}
	fmt.Printf("Font opened: %v\n", font)
	defer func() {
		if err := font.Close(); err != nil {
			fmt.Printf("Closing font error: %v\n", err)
			return
		}
		fmt.Printf("Font closed: %v\n", font)
	}()

	fontInfo, err := font.Info()
	if err != nil {
		fmt.Printf("Error getting font informations for font \"%v\": %v\n", font, err)
		return
	}
	fmt.Printf("Font info properties: %v\n", fontInfo.Properties())
*/
/*

	{ // draw something to pixmap



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
