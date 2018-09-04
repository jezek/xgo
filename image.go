package xgo

import (
	"fmt"
	"image"
	"log"

	"github.com/jezek/xgb/xproto"
)

type Pixmap struct {
	xproto.Pixmap
	s     *Screen
	Size  image.Point
	Depth byte
}

func (p *Pixmap) Screen() *Screen {
	if p.s == nil {
		log.Fatalf("Pixmap %s has no screen", p)
	}
	return p.s
}

func (p *Pixmap) Destroy() error {
	if p.Pixmap == 0 {
		return nil
	}
	if err := xproto.FreePixmapChecked(
		p.Screen().Display().Conn,
		p.Pixmap,
	).Check(); err != nil {
		return errWrap{"Pixmap.Destroy", fmt.Errorf("pixmap %v free error: %v", p, err)}
	}
	p.Pixmap = 0
	return nil
}

// Draws image to pixmap
func (p *Pixmap) DrawImage(img image.Image) error {

	imageBounds := img.Bounds()

	// convert image to BGRA order, cause x11 works this way
	pixBGRA := make([]byte, 0, imageBounds.Dx()*imageBounds.Dy())
	for y := imageBounds.Min.Y; y < imageBounds.Max.Y; y++ {
		for x := imageBounds.Min.X; x < imageBounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			// r is uint32, but color inside is from 0 to 0xffff (16 bit), so for a 8 bit repesentation just shift right by 8 bits
			pixBGRA = append(pixBGRA, byte(b>>8), byte(g>>8), byte(r>>8), byte(a>>8))
		}
	}
	//fmt.Println("image converted to BGRA bytes", pixBGRA[:60], "...")

	// create and allocate graphical content
	gc, err := xproto.NewGcontextId(p.Screen().Display().Conn)
	if err != nil {
		return errWrap{"Pixmap.DrawImage", fmt.Errorf("unable to obtain graphic context id: %v\n", err)}
	}
	if err := xproto.CreateGCChecked(
		p.Screen().Display().Conn,
		gc,
		xproto.Drawable(p.Screen().Root),
		xproto.GcForeground|xproto.GcBackground,
		[]uint32{p.Screen().WhitePixel, p.Screen().BlackPixel},
	).Check(); err != nil {
		return errWrap{"Pixmap.DrawImage", fmt.Errorf("unable to create graphic context: %v\n", err)}
	}
	//fmt.Println("gc created")

	// free graphic context after image drawing
	defer func() {
		if err := xproto.FreeGCChecked(
			p.Screen().Display().Conn,
			gc,
		).Check(); err != nil {
			log.Printf("Pixmap.DrawImage: freeing graphic context error: %v", err)
		}
	}()

	// we cand send limited numbers of bytes per request, so calculate batches for drawing
	maxBytesToSend := SizeReqestMax - (SizeRequestPutImageFixedPart + 4)
	maxAreaToSend := maxBytesToSend / 4
	if maxAreaToSend*4 != maxBytesToSend {
		//fmt.Printf("maxAreaToSend(%d) * 4 (%d) != maxBytesToSend (%d)\n", maxAreaToSend, maxAreaToSend*4, maxBytesToSend)
		maxBytesToSend = maxAreaToSend * 4
	}
	rectangles := decomposeToRectanglesWithAreaMax(imageBounds, maxAreaToSend)

	// draw our image to pixmap
	for _, rectangle := range rectangles {
		// draw our image stored in pixBGRA to pixmap (pid) using gc
		if err := xproto.PutImageChecked(
			p.Screen().Display().Conn,
			xproto.ImageFormatZPixmap,
			xproto.Drawable(p.Pixmap),
			gc,
			uint16(rectangle.Dx()),
			uint16(rectangle.Dy()),
			int16(rectangle.Min.X),
			int16(rectangle.Min.Y),
			0, // left padding
			p.Depth,
			getRectangleBytes(rectangle, pixBGRA),
		).Check(); err != nil {
			return errWrap{"Pixmap.DrawImage", fmt.Errorf("unable to write image bgra pixels to pixmap: %v\n", err)}
		}
		//fmt.Printf("pixBGRA data for image rectangle %v inserted into drawable pixmap, using gc\n", rectangle)
	}
	return nil
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
