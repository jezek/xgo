package xgo

import (
	"fmt"
	"log"

	"github.com/BurntSushi/xgb/xproto"
)

type DisplayPointer struct {
	d *Display
}

func (p *DisplayPointer) Status() (PointerStatus, error) {
	for _, s := range p.d.Screens() {
		ps, err := s.Window().Pointer().Status()
		if err != nil {
			return PointerStatus{}, err
		}
		if ps.SameScreen {
			return ps, nil
		}
	}
	return PointerStatus{}, fmt.Errorf("Mouse pointer not found")
}

type Pointer struct {
	w *Window
}

type PointerStatus struct {
	*xproto.QueryPointerReply
	Pointer *Pointer
	Root    *Window
	Child   *Window
	Button  [6]bool
}

func (p *Pointer) Window() *Window {
	if p.w == nil {
		log.Fatalf("Poiner %s has no window", p)
	}
	return p.w
}

func (p *Pointer) String() string {
	return fmt.Sprintf("p%s", p.Window())
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

func (p *Pointer) Status() (PointerStatus, error) {
	reply, err := xproto.QueryPointer(
		p.Window().Screen().Display().Conn,
		p.Window().Window,
	).Reply()
	if err != nil {
		return PointerStatus{}, err
	}
	var root, child *Window
	if reply.Root != xproto.WindowNone {
		root, _ = p.Window().Screen().Display().FindWindow(uint32(reply.Root))
	}
	if reply.Child != xproto.WindowNone {
		child, _ = p.Window().Screen().Display().FindWindow(uint32(reply.Child))
	}
	return PointerStatus{
		reply,
		p,
		root, child,
		p.buttonsFromMask(reply.Mask),
	}, nil
}

func (p *Pointer) MotionNotify(stop <-chan struct{}) <-chan xproto.MotionNotifyEvent {
	return p.Window().Screen().Display().Events().listenMotionNotify(p.Window(), stop)
}
