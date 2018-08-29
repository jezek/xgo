package xgo

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

const (
	IsUnmapped   = xproto.MapStateUnmapped
	IsUnviewable = xproto.MapStateUnviewable
	IsViewable   = xproto.MapStateViewable
)

var windowLog *log.Logger = func() *log.Logger {
	writer := io.Writer(ioutil.Discard)
	if DEBUG {
		writer = os.Stderr
	}
	return log.New(writer, "window: ", log.LstdFlags)
}()

type Window struct {
	xproto.Window         //required
	s             *Screen //required

	p *Pointer  //not required
	k *Keyboard //not required
}

func (w *Window) Screen() *Screen {
	if w.s == nil {
		log.Fatalf("Window %s has no screen", w)
	}
	return w.s
}

func (w *Window) Name() string {

	reply, err := w.getProperty("_NET_WM_NAME")
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
	windowLog.Printf("Window %s atrributes reply %v, %v", w, reply, err)
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

// Maps window
func (w *Window) Map() error {
	//TODO what if window is allready mapped?
	if err := xproto.MapWindowChecked(w.Screen().Display().Conn, w.Window).Check(); err != nil {
		return err
	}
	return nil
}

func (w *Window) Pointer() *Pointer {
	//TODO possible race condition, use mutex
	if w.p == nil {
		w.p = &Pointer{w, nil}
	}
	return w.p
}

func (w *Window) Keyboard() *Keyboard {
	//TODO possible race condition, use mutex
	if w.k == nil {
		w.k = &Keyboard{w, nil}
	}
	return w.k
}

func (w *Window) getProperty(name string) (*xproto.GetPropertyReply, error) {

	// get xproto.Atom from name
	atomId, err := w.Screen().Display().atoms().GetByName(name)
	if err != nil {
		return nil, errWrap{"Window.getProperty", fmt.Errorf("error getting atom: %v", err)}
	}

	windowLog.Printf("Window.getProperty: %s = %v", name, atomId)

	// get xproto.GetProperty for atom
	reply, err := xproto.GetProperty(
		w.Screen().Display().Conn,
		false,
		w.Window,
		atomId,
		xproto.GetPropertyTypeAny,
		0, (1<<32)-1,
	).Reply()

	if err != nil {
		return nil, errWrap{"Window.getProperty", fmt.Errorf("error getting property \"%s\" for window %v: %v", name, w, err)}
	}

	return reply, nil
}

func (w *Window) ProtocolsSet(protocols []string) error {
	// convert strings to atoms, if atom does not exist, then create it
	atoms := make([]xproto.Atom, 0, len(protocols))
	for _, protocol := range protocols {
		if aid, err := w.Screen().Display().atoms().GetByName(protocol); err != nil {
			return errWrap{"Window.ProtocolsSet", fmt.Errorf("error getting atom by name \"%s\" for window %v: %v", protocol, w, err)}
		} else {
			atoms = append(atoms, aid)
		}
	}

	//TODO convert to 32 bit data in one previous cycle
	// convert atoms to 32 bit data
	atomData := make([]byte, 4)
	data := make([]byte, 0, len(atoms)*len(atomData))
	for _, atomId := range atoms {
		xgb.Put32(atomData, uint32(atomId))
		data = append(data, atomData...)
	}

	propAtom, err := w.Screen().Display().atoms().GetByName("WM_PROTOCOLS")
	if err != nil {
		return errWrap{"Window.ProtocolsSet", fmt.Errorf("error getting atom by name \"%s\" for window %v: %v", "WM_PROTOCOLS", w, err)}
	}

	typAtom, err := w.Screen().Display().atoms().GetByName("ATOM")
	if err != nil {
		return errWrap{"Window.ProtocolsSet", fmt.Errorf("error getting atom by name \"%s\" for window %v: %v", "ATOM", w, err)}
	}

	// change property of window by sending agregated data
	return xproto.ChangePropertyChecked(
		w.Screen().Display().Conn,
		xproto.PropModeReplace,
		w.Window,
		propAtom,
		typAtom,
		32,
		uint32(len(data)/4), // data length
		data,
	).Check()
}

func (w *Window) Protocols() ([]string, error) {
	// get window protocols as properties
	reply, err := w.getProperty("WM_PROTOCOLS")
	if err != nil {
		return nil, errWrap{"Window.Protocols", err}
	}

	if reply.Format == 0 {
		// there is no such property, return empty slice
		return []string{}, nil
		//return nil, errWrap{"Window.Protocols", fmt.Errorf("there is no property \"%s\" for window %v", name, w)}
	}

	windowLog.Printf("Window.Protocols: getProperty reply = %v", reply)

	// decode properties to []string
	if reply.Format != 32 {
		// sorry, can do only 32bit format for now
		return nil, errWrap{"Window.Protocols", fmt.Errorf("want 32 bit protocols property reply for window %v, got %d", w, reply.Format)}
	}

	res := make([]string, 0, reply.ValueLen)

	// result.Value is a slice of Atom identifiers, every 4B long
	bytes := reply.Value
	for len(bytes) >= 4 { // 4B=32b
		name, err := w.Screen().Display().atoms().GetById(xproto.Atom(xgb.Get32(bytes)))
		if err != nil {
			return nil, errWrap{"Window.Protocols", fmt.Errorf("want 32 bit protocols property reply for window %v, got %d", w, reply.Format)}
		}
		res = append(res, name)
		bytes = bytes[4:]
	}

	return res, nil
}

func (w *Window) CloseNotify(stop <-chan struct{}) (<-chan struct{}, error) {
	// get window protocols
	protocols, err := w.Protocols()
	if err != nil {
		return nil, errWrap{"Window.CloseNotify", fmt.Errorf("window %v get protocols error: %v", w, err)}
	}

	windowLog.Printf("window %v protocols: %v", w, protocols)
	{ // append a WM_DELETE_WINDOW to protocols, if not allready added
		isDelete := false
		for _, protocol := range protocols {
			if protocol == "WM_DELETE_WINDOW" {
				isDelete = true
				break
			}
		}
		if !isDelete {
			if err := w.ProtocolsSet(append(protocols, "WM_DELETE_WINDOW")); err != nil {
				return nil, errWrap{"Window.CloseNotify", fmt.Errorf("window %v set protocols error: %v", w, err)}
			}
			windowLog.Printf("updated window %v protocols: %v", w, protocols)
		}
	}

	// the protocol is set, now listen to it
	return w.Screen().Display().events().listenWmDeleteWindow(w, stop), nil
}

var errNilWindow error = errors.New("window is nil")

// Makes a ConfigRequest, to alter position and size. Returns error if something is wrong
func (w *Window) MoveResize(bounds image.Rectangle) error {
	if w == nil {
		return errWrap{"Window.MoveResize", errNilWindow}
	}
	//TODO if this is a top-level window in WM that supports EWMH, use WMMoveResize
	if bounds.Dx() <= 0 || bounds.Dy() <= 0 {
		return errors.New("Window.MoveResize: zero width or height " + bounds.String())
	}

	flags := uint16(xproto.ConfigWindowX | xproto.ConfigWindowY | xproto.ConfigWindowWidth | xproto.ConfigWindowHeight)
	vals := []uint32{uint32(bounds.Min.X), uint32(bounds.Min.Y), uint32(bounds.Dx()), uint32(bounds.Dy())}
	if err := xproto.ConfigureWindowChecked(
		w.Screen().Display().Conn,
		w.Window,
		flags,
		vals,
	).Check(); err != nil {
		return errWrap{"Window.MoveResize", fmt.Errorf("ConfigWindow request resulted with error: %v", err)}
	}
	//TODO cache size & position
	return nil
}

func (w *Window) Destroy() error {
	if w.Window != xproto.WindowNone {
		//TODO deatach events

		if err := xproto.DestroyWindowChecked(w.Screen().Display().Conn, w.Window).Check(); err != nil {
			return errWrap{"Window.Destroy", fmt.Errorf("window %v destroy error: %v", w, err)}
		}
		w.Window = xproto.WindowNone
	}
	return nil
}

func (w *Window) BgColorChange(col color.Color) error {
	//TODO convert color to uint32 color used by xlib
	bgColor := uint32(0)
	if err := xproto.ChangeWindowAttributesChecked(
		w.Screen().Display().Conn,
		w.Window,
		xproto.CwBackPixel,
		[]uint32{bgColor},
	).Check(); err != nil {
		return errWrap{"Window.BgColorChange", fmt.Errorf("ChangeWindowAttributes error: %v", err)}
	}

	return nil
}
