package xgo

import "github.com/BurntSushi/xgb/xproto"

// Screen instance
type Screen struct {
	*xproto.ScreenInfo
	d   *Display
	id  int
	def bool
	w   *Window
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
