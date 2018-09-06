package xgo

import (
	"fmt"
	"image"

	"github.com/jezek/xgb/xproto"
)

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
			[]PixmapOperation{PixmapOperations{}.DrawImage(img)},
			operations...,
		)...,
	)
}

type PixmapOperation func(*Pixmap) error
type PixmapOperations struct{}

func (_ PixmapOperations) DrawImage(img image.Image) PixmapOperation {
	return func(p *Pixmap) error {
		return p.DrawImage(img)
	}
}
