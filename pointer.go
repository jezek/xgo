package xgo

import (
	"fmt"
	"log"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgb/xtest"
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
	c *PointerControll
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

func (p *Pointer) Screen() *Screen {
	return p.Window().Screen()
}

func (p *Pointer) Display() *Display {
	return p.Screen().Display()
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

func (p *Pointer) Control() *PointerControll {
	if p.c == nil {
		p.c = &PointerControll{p}
	}
	return p.c
}

type PointerControll struct {
	p *Pointer
}

func (c *PointerControll) Screen() *Screen {
	return c.p.Screen()
}

func (c *PointerControll) Display() *Display {
	return c.Screen().Display()
}

func (c *PointerControll) Move(x, y int) error {
	if err := c.Display().extension("xtest"); err != nil {
		return err
	}
	if err := xtest.FakeInputChecked(
		c.Display().Conn,
		xproto.MotionNotify,
		0,
		xproto.TimeCurrentTime,
		c.Screen().Window().Window,
		int16(x), int16(y),
		0,
	).Check(); err != nil {
		return err
	}
	return nil
}

func (c *PointerControll) MoveRelative(x, y int) error {
	if err := c.Display().extension("xtest"); err != nil {
		return err
	}
	if err := xtest.FakeInputChecked(
		c.Display().Conn,
		xproto.MotionNotify,
		1,
		xproto.TimeCurrentTime,
		c.Screen().Window().Window,
		int16(x), int16(y),
		0,
	).Check(); err != nil {
		return err
	}
	return nil
}

func (c *PointerControll) click(bi byte) error {
	if err := c.Display().extension("xtest"); err != nil {
		return err
	}
	if err := xtest.FakeInputChecked(
		c.Display().Conn,
		xproto.ButtonPress,
		bi,
		xproto.TimeCurrentTime,
		c.Screen().Window().Window,
		int16(0), int16(0),
		0,
	).Check(); err != nil {
		return err
	}
	if err := xtest.FakeInputChecked(
		c.Display().Conn,
		xproto.ButtonRelease,
		bi,
		xproto.TimeCurrentTime,
		c.Screen().Window().Window,
		int16(0), int16(0),
		0,
	).Check(); err != nil {
		return err
	}
	return nil
}
func (c *PointerControll) ClickLeft() error {
	return c.click(xproto.ButtonIndex1)
}
func (c *PointerControll) ClickRight() error {
	return c.click(xproto.ButtonIndex3)
}
func (c *PointerControll) ScrollUp() error {
	return c.click(xproto.ButtonIndex4)
}
func (c *PointerControll) ScrollDown() error {
	return c.click(xproto.ButtonIndex5)
}
