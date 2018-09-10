package xgo

import (
	"fmt"
	"image"
	"log"
	"unicode/utf16"

	"github.com/jezek/xgb/xproto"
)

type Pixmap struct {
	xproto.Pixmap
	s     *Screen
	Size  image.Point
	Depth byte
}

func (p *Pixmap) Display() *Display {
	return p.Screen().Display()
}
func (p *Pixmap) Screen() *Screen {
	if p.s == nil {
		log.Fatalf("Pixmap %v has no screen", p)
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

type PixmapDrawer func(*Pixmap) error
type PixmapDrawers struct{}

var pd PixmapDrawers

func (_ PixmapDrawers) Image(img image.Image) PixmapDrawer {
	// precompute everything what we can

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

	// we cand send limited numbers of bytes per request, so calculate batches for drawing
	maxBytesToSend := SizeReqestMax - (SizeRequestPutImageFixedPart + 4)
	maxAreaToSend := maxBytesToSend / 4
	if maxAreaToSend*4 != maxBytesToSend {
		//fmt.Printf("maxAreaToSend(%d) * 4 (%d) != maxBytesToSend (%d)\n", maxAreaToSend, maxAreaToSend*4, maxBytesToSend)
		maxBytesToSend = maxAreaToSend * 4
	}

	return func(p *Pixmap) error {
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

		//TODO the rectangles can be precomputed!!!
		rectangles := p.decomposeToRectanglesWithAreaMax(imageBounds, maxAreaToSend)

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
				p.getRectangleBytes(rectangle, pixBGRA),
			).Check(); err != nil {
				return errWrap{"Pixmap.DrawImage", fmt.Errorf("unable to write image bgra pixels to pixmap: %v\n", err)}
			}
			//fmt.Printf("pixBGRA data for image rectangle %v inserted into drawable pixmap, using gc\n", rectangle)
		}
		return nil
	}
}

func (_ PixmapDrawers) Text(text string, position image.Point, gc *GraphicsContext) PixmapDrawer {
	uint16String := utf16.Encode([]rune(text))
	c2bString := make([]xproto.Char2b, len(uint16String))
	for i, v := range uint16String {
		c2bString[i].Byte1 = byte(v >> 8)
		c2bString[i].Byte2 = byte(v)
	}
	return func(p *Pixmap) error {
		return xproto.ImageText16Checked(
			p.Display().Conn,
			byte(len(c2bString)),
			xproto.Drawable(p.Pixmap),
			gc.Gcontext,
			int16(position.X),
			int16(position.Y),
			c2bString,
		).Check()
	}
	/*
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

		fmt.Println("Wrote some text to pixmap")
	*/
}

// Uses all draw functions in order. Stops on first encountered error and returns it.
func (p *Pixmap) Draw(drawers ...PixmapDrawer) error {
	for i, drawer := range drawers {
		if err := drawer(p); err != nil {
			return errWrap{"Pixmap.Draw", fmt.Errorf("error drawing with drawer no %d: %v", i, err)}
		}
	}
	return nil
}

func (p *Pixmap) decomposeToRectanglesWithAreaMax(rect image.Rectangle, maxArea int) []image.Rectangle {
	if rect.Dx()*rect.Dy() <= maxArea {
		return []image.Rectangle{rect}
	}
	// area is greater than maxArea, so we are splitting
	//TODO? optimize... but how? use cycles instead of recursion

	// split verticaly, if deeded
	if rect.Dx() > maxArea {
		return append(
			p.decomposeToRectanglesWithAreaMax(
				image.Rect(
					rect.Min.X, rect.Min.Y,
					rect.Min.X+maxArea, rect.Max.Y,
				),
				maxArea,
			),
			p.decomposeToRectanglesWithAreaMax(
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
		p.decomposeToRectanglesWithAreaMax(
			image.Rect(
				rect.Min.X, rect.Min.Y+maxRows,
				rect.Max.X, rect.Max.Y,
			),
			maxArea,
		)...,
	)

}

func (_ *Pixmap) getRectangleBytes(rect image.Rectangle, quadrupletBytes []byte) []byte {
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

// Allocate new pixmap in X on screen, with screen default depth and provided size
// Then applies all pixmap operations.
// If there is an error along somewhere, pixmap is destroyed and the first occuring error is returned
func NewPixmapOnScreen(s *Screen, size image.Point, operations ...PixmapOperation) (*Pixmap, error) {
	// first get new pixmap id
	pixmapId, err := xproto.NewPixmapId(s.Display().Conn)
	if err != nil {
		return nil, errWrap{"Screen.NewPixmap", fmt.Errorf("unable to obtain pixmap id: %v", err)}
	}
	//TODO? what about depth?

	// then send create request for pixmap
	if err := xproto.CreatePixmapChecked(
		s.Display().Conn,
		s.RootDepth,
		pixmapId,
		xproto.Drawable(s.Root),
		uint16(size.X),
		uint16(size.Y),
	).Check(); err != nil {
		return nil, errWrap{"Screen.NewPixmap", fmt.Errorf("unable to create pixmap: %v\n", err)}

	}
	p := &Pixmap{
		pixmapId,
		s,
		size,
		s.RootDepth,
	}

	for i, operation := range operations {
		if err := operation(p); err != nil {
			p.Destroy()
			return nil, errWrap{"NewPixmapOnScreen", fmt.Errorf("operation no %d error: %v", i, err)}
		}
	}

	return p, nil
}

// Allocate new pixmap in X on screen, with screen default depth and size of the image
// Then draws the image to pixmap and applies all pixmap operations.
// If there is an error along somewhere, pixmap is destroyed and the first occuring error is returned
// This is an shortcut for:
// NewPixmapOnScreen(
// 	scr,
// 	image.Pt(img.Bounds().Dx(),img.Bounds().Dy()),
// 	PixmapOperations{}.DrawImage(img),
// 	options...,
// 	)
func NewPixmapFromImageOnScreen(scr *Screen, img image.Image, operations ...PixmapOperation) (*Pixmap, error) {
	return NewPixmapOnScreen(
		scr,
		image.Pt(img.Bounds().Dx(), img.Bounds().Dy()),
		append(
			[]PixmapOperation{po.Draw(pd.Image(img))},
			operations...,
		)...,
	)
}

type PixmapOperation func(*Pixmap) error
type PixmapOperations struct{}

var po PixmapOperations

func (_ PixmapOperations) Draw(drawers ...PixmapDrawer) PixmapOperation {
	return func(p *Pixmap) error {
		return p.Draw(drawers...)
	}
}
