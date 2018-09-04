package xgo

import (
	"fmt"
	"image"
	"log"

	"github.com/jezek/xgb/xproto"
)

// Screen instance
type Screen struct {
	*xproto.ScreenInfo
	d   *Display
	id  int
	def bool
	w   *Window
}

func (s *Screen) Display() *Display {
	if s.d == nil {
		log.Fatalf("Screen %d has no display", s.id)
	}
	return s.d
}

func (s *Screen) Default() bool {
	return s.def
}

func (s *Screen) Id() int {
	return s.id
}

func (s *Screen) Window() *Window {
	if s.w == nil {
		s.w = &Window{
			s.ScreenInfo.Root, s,
			nil, nil,
		}
	}
	return s.w
}

// Allocate new pixmap in X on screen, with screen default depth
func (s *Screen) NewPixmap(size image.Point) (*Pixmap, error) {
	// first get new pixmap id
	pixmapId, err := xproto.NewPixmapId(s.Display().Conn)
	if err != nil {
		return nil, errWrap{"Screen.NewPixmap", fmt.Errorf("unable to obtain pixmap id: %v", err)}
	}
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
	return &Pixmap{
		pixmapId,
		s,
		size,
		s.RootDepth,
	}, nil
}
