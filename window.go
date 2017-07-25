package xgo

import (
	"fmt"
	"log"

	"github.com/BurntSushi/xgb/xproto"
)

const (
	IsUnmapped = iota
	IsUnviewable
	IsViewable
)

type Window struct {
	xproto.Window
	s *Screen
	p *Pointer
	k *Keyboard
}

func (w *Window) Screen() *Screen {
	if w.s == nil {
		log.Fatalf("Window %s has no screen", w)
	}
	return w.s
}

func (w *Window) Name() string {
	aname := "_NET_WM_NAME"
	req, err := xproto.InternAtom(w.s.d.Conn, true, uint16(len(aname)), aname).Reply()
	if err != nil {
		log.Fatal(err)
	}
	if req.Atom == xproto.AtomNone {
		log.Fatalf("No %s intern atom found", aname)
	}

	reply, err := xproto.GetProperty(w.s.d.Conn, false, w.Window, req.Atom, xproto.GetPropertyTypeAny, 0, (1<<32)-1).Reply()
	if err != nil {
		log.Fatal(err)
	}

	return string(reply.Value)
}

func (w *Window) IsRoot() bool {
	return w.Window == w.s.Window().Window
}

func (w *Window) Parent() (*Window, error) {
	if w.IsRoot() {
		return nil, nil
	}
	t, err := xproto.QueryTree(w.s.d.Conn, w.Window).Reply()
	if err != nil {
		return nil, err
	}
	return &Window{
		t.Parent, w.Screen(),
		nil, nil,
	}, nil
}

func (w *Window) Parents() ([]*Window, error) {
	if w.IsRoot() {
		return []*Window{}, nil
	}
	res := []*Window{}
	var err error
	p := w
	for p != nil {
		p, err = p.Parent()
		if err != nil {
			return res, err
		}
		if p == nil {
			break
		}
		res = append(res, p)
	}
	return res, nil
}

func (w *Window) Children() ([]*Window, error) {
	t, err := xproto.QueryTree(w.s.d.Conn, w.Window).Reply()
	if err != nil {
		return nil, err
	}
	ch := make([]*Window, t.ChildrenLen)
	for i := range ch {
		chw := &Window{
			t.Children[i], w.Screen(),
			nil, nil,
		}
		ch[i] = chw
	}
	return ch, nil
}

func (w *Window) String() string {
	return fmt.Sprintf("wid:%d", w.Window)
}

type WindowAttributes struct {
	*xproto.GetWindowAttributesReply
	w *Window
}

func (w *Window) Attributes() (*WindowAttributes, error) {
	reply, err := xproto.GetWindowAttributes(w.s.d.Conn, w.Window).Reply()
	debugf("Window %s atrributes reply %v, %v", w, reply, err)
	if err != nil {
		return nil, err
	}
	return &WindowAttributes{reply, w}, nil
}

func (w *Window) IsVisible() bool {
	ret, err := w.Attributes()
	if err != nil {
		log.Printf("Can't obtain attributes for window %s due to error: %s", w, err)
		return false
	}
	return ret.MapState == IsViewable
}

func (w *Window) Pointer() *Pointer {
	if w.p == nil {
		w.p = &Pointer{w, nil}
	}
	return w.p
}

func (w *Window) Keyboard() *Keyboard {
	if w.k == nil {
		w.k = &Keyboard{w, nil}
	}
	return w.k
}
