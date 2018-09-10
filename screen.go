package xgo

import (
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

func (s *Screen) NewWindow(operations ...WindowOperation) (*Window, error) {
	return NewWindowOnScreen(s, operations...)
}

// Calls NewPixmapOnScreen on current screen
func (s *Screen) NewPixmap(size image.Point, operations ...PixmapOperation) (*Pixmap, error) {
	return NewPixmapOnScreen(s, size, operations...)
}

// Calls NewPixmapFromImageOnScreen on current screen
func (s *Screen) NewPixmapFromImage(image image.Image, operations ...PixmapOperation) (*Pixmap, error) {
	return NewPixmapFromImageOnScreen(s, image, operations...)
}

// Creates GraphicsContext with this creens root window depth.
func (s *Screen) NewGraphicsContext(components ...GraphicsContextComponent) (*GraphicsContext, error) {
	return NewGraphicsContextOnDisplay(s.Display(), s.Window(), components...)
}
