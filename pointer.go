package xgo

import (
	"fmt"

	"github.com/BurntSushi/xgb/xproto"
)

type Pointer struct {
	d *Display
}

type PointerStatus struct {
	Screen *Screen
	X, Y   int16
	Button [6]bool
}

func (p *Pointer) statusFromQueryPointerReply(s *Screen, reply *xproto.QueryPointerReply) PointerStatus {
	return PointerStatus{
		s,
		reply.RootX, reply.RootY,
		p.buttonsFromMask(reply.Mask),
	}
}

func (p *Pointer) Status() (PointerStatus, error) {
	for _, s := range p.d.Screens() {
		reply, err := xproto.QueryPointer(p.d.Conn, s.Window().Window).Reply()
		if err != nil {
			return PointerStatus{}, err
		}
		if reply.SameScreen {
			return p.statusFromQueryPointerReply(s, reply), nil
		}
	}
	return PointerStatus{}, fmt.Errorf("Mouse pointer not found")
}

func (p *Pointer) buttonsFromMask(m uint16) [6]bool {
	return [6]bool{
		xproto.ButtonIndexAny: (m & xproto.ButtonMaskAny) != 0,
		xproto.ButtonIndex1:   (m & xproto.ButtonMask1) != 0,
		xproto.ButtonIndex2:   (m & xproto.ButtonMask2) != 0,
		xproto.ButtonIndex3:   (m & xproto.ButtonMask3) != 0,
		xproto.ButtonIndex4:   (m & xproto.ButtonMask4) != 0,
		xproto.ButtonIndex5:   (m & xproto.ButtonMask5) != 0,
	}
}

type WindowPointer struct {
	*Pointer
	w *Window
}

type WindowPointerStatus struct {
	PointerStatus
	Window   *Window
	OnScreen bool
	X, Y     int16
	Child    *Window
}

func (p *WindowPointer) Status() (WindowPointerStatus, error) {
	reply, err := xproto.QueryPointer(p.d.Conn, p.w.Window).Reply()
	if err != nil {
		return WindowPointerStatus{}, err
	}
	var child *Window
	if reply.Child != 0 {
		child, _ = p.w.Screen().Display().FindWindow(uint32(reply.Child))
	}
	return WindowPointerStatus{
		p.Pointer.statusFromQueryPointerReply(p.w.s, reply),
		p.w,
		reply.SameScreen,
		reply.WinX, reply.WinY,
		child,
	}, nil
}
