package xgo

import (
	"fmt"

	"github.com/BurntSushi/xgb/xproto"
)

type Mouse struct {
	d *Display
}

type MouseStatus struct {
	Screen *Screen
	X, Y   int16
	Button [6]bool
}

func (m *Mouse) Status() (MouseStatus, error) {
	for _, s := range m.d.Screens() {
		reply, err := xproto.QueryPointer(m.d.Conn, s.Window().Window).Reply()
		if err != nil {
			return MouseStatus{}, err
		}
		if reply.SameScreen {
			return MouseStatus{
				s,
				reply.RootX, reply.RootY,
				[6]bool{
					xproto.ButtonIndexAny: (reply.Mask & xproto.ButtonMaskAny) != 0,
					xproto.ButtonIndex1:   (reply.Mask & xproto.ButtonMask1) != 0,
					xproto.ButtonIndex2:   (reply.Mask & xproto.ButtonMask2) != 0,
					xproto.ButtonIndex3:   (reply.Mask & xproto.ButtonMask3) != 0,
					xproto.ButtonIndex4:   (reply.Mask & xproto.ButtonMask4) != 0,
					xproto.ButtonIndex5:   (reply.Mask & xproto.ButtonMask5) != 0,
				},
			}, nil
		}
	}
	return MouseStatus{}, fmt.Errorf("Mouse pointer not found")
}
