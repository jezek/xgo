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
	s  *Screen
	p  *Window
	ch []*Window
	wp *WindowPointer
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
	if w.p == nil {
		t, err := xproto.QueryTree(w.s.d.Conn, w.Window).Reply()
		if err != nil {
			return nil, err
		}
		w.p = &Window{
			t.Parent, w.s,
			nil, nil, nil,
		}
	}
	return w.p, nil
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
	if w.ch == nil {
		t, err := xproto.QueryTree(w.s.d.Conn, w.Window).Reply()
		if err != nil {
			return nil, err
		}
		w.ch = make([]*Window, t.ChildrenLen)
		for i := range w.ch {
			chw := &Window{
				t.Children[i], w.s,
				w, nil, nil,
			}
			w.ch[i] = chw
		}
	}
	return w.ch, nil
}

func (w *Window) String() string {
	return fmt.Sprintf("wid: %d", w.Window)
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

func (w *Window) Pointer() *WindowPointer {
	if w.wp == nil {
		w.wp = &WindowPointer{
			w.s.d.Pointer(),
			w,
		}
	}
	return w.wp
}
