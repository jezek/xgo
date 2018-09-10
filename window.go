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
	"sort"

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

func (w *Window) Display() *Display {
	return w.Screen().Display()
}

func (w *Window) Drawable() xproto.Drawable {
	return xproto.Drawable(w.Window)
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

type WindowAttributesInfo struct {
	*xproto.GetWindowAttributesReply
	w *Window
}

func (w *Window) AttributesInfo() (*WindowAttributesInfo, error) {
	reply, err := xproto.GetWindowAttributes(w.s.d.Conn, w.Window).Reply()
	windowLog.Printf("Window %s atrributes reply %v, %v", w, reply, err)
	if err != nil {
		return nil, err
	}
	return &WindowAttributesInfo{reply, w}, nil
}

func (w *Window) IsVisible() bool {
	ret, err := w.AttributesInfo()
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

// Makes a ConfigRequest, to alter size. Returns error if something is wrong
func (w *Window) Resize(size image.Point) error {
	if w == nil {
		return errWrap{"Window.Resize", errNilWindow}
	}
	if size.X <= 0 || size.Y <= 0 {
		return errors.New("Window.Resize: zero width or height " + size.String())
	}

	flags := uint16(xproto.ConfigWindowWidth | xproto.ConfigWindowHeight)
	vals := []uint32{uint32(size.X), uint32(size.Y)}
	if err := xproto.ConfigureWindowChecked(
		w.Screen().Display().Conn,
		w.Window,
		flags,
		vals,
	).Check(); err != nil {
		return errWrap{"Window.Resize", fmt.Errorf("ConfigWindow request resulted with error: %v", err)}
	}
	//TODO cache size
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

// Applies attributes to window. If duplicate attributes provided, error is produced.
// The attributes are not applied in provided order and are allways sorted by mask of the attribute.
// eg. if you use BackgroundPixelColor and BackgroundPixmap, it will not show the pixmap on background,
// because mask for BackgroundPixelColor has greater value than BackgroundPixmap and will com after. So every time the last background operation is BackgroundPixelColor.
// See WindowAttributes struct and underlying xproto.Cw... costants for mask value info
// TODO? or should I return error if background is set twice, or any other alike situations?
func (w *Window) AttributesChange(attributes ...WindowAttribute) error {
	// sort attributes by mask ascending
	sort.Slice(attributes, func(i, j int) bool {
		im, _ := attributes[i]()
		jm, _ := attributes[j]()
		return im < jm
	})

	mask := uint32(0)
	values := make([]uint32, 0) //, len(attributes))

	for _, attribute := range attributes {
		am, av := attribute()
		if mask|am == mask {
			return errWrap{"Window.AttributesChange", fmt.Errorf("duplicate attrbute: %#v", attribute)}
		}
		mask = mask | am
		values = append(values, av)
	}
	if mask == 0 {
		return nil
	}
	if err := xproto.ChangeWindowAttributesChecked(
		w.Screen().Display().Conn,
		w.Window,
		mask,
		values,
	).Check(); err != nil {
		return errWrap{"Window.AttributesChange", fmt.Errorf("ChangeWindowAttributes error: %v", err)}
	}
	return nil
}

// Paints a rectangular area in the window with the window's background pixel or pixmap.
// If the window has a defined background tile, the rectangle clipped by any children is filled with this tile.
// If the window has background None, the contents of the window are not changed.
func (w *Window) Clear() error {
	return xproto.ClearAreaChecked(
		w.Screen().Display().Conn,
		false,
		w.Window,
		0, 0, 0, 0,
	).Check()
}

// Creates an unmapped input/output window on default screen with no initialization.
// Window size is [1,1], position is [0,0], border 0, with screen root depth and none (default?) attributes.
// Then applies all window operations.
// If there is an error along somewhere, window is destroyed and the first occuring error is returned
func NewWindowOnScreen(s *Screen, operations ...WindowOperation) (*Window, error) {
	//TODO check if screen is a valid screen
	wid, err := xproto.NewWindowId(s.Display().Conn)
	if err != nil {
		return nil, err
	}

	parentWindow := s.Root

	//TODO some WMs dont care about coords, so handle position with respect to WM
	if err := xproto.CreateWindowChecked(
		s.Display().Conn,
		s.RootDepth,
		wid,
		parentWindow,
		0, 0, //x, y
		1, 1, //width, height
		0, //border width
		xproto.WindowClassInputOutput, //window class
		s.RootVisual,                  //VisualId
		0,                             //window attributes mask
		[]uint32{},                    //window attributes to set for masks
	).Check(); err != nil {
		return nil, err
	}

	//TODO remember created windows and hadle them on display close
	w := &Window{wid, s, nil, nil}

	for i, operation := range operations {
		if err := operation(w); err != nil {
			w.Destroy()
			return nil, errWrap{"NewWindowOnScreen", fmt.Errorf("operation no %d error: %v", i, err)}
		}
	}
	return w, nil
}

type WindowOperation func(*Window) error
type WindowOperations struct{}

//TODO ConfigWindowX, ConfigWindowY, ConfigWindowBorderWidth, ConfigWindowSibling, ConfigWindowStackMode,
func (_ WindowOperations) Size(size image.Point) WindowOperation {
	return func(w *Window) error {
		return w.Resize(size)
	}
}

func (_ WindowOperations) Attributes(attributes ...WindowAttribute) WindowOperation {
	return func(w *Window) error {
		return w.AttributesChange(attributes...)
	}
}
func (_ WindowOperations) Clear() WindowOperation {
	return func(w *Window) error {
		return w.Clear()
	}
}
func (_ WindowOperations) Map() WindowOperation {
	return func(w *Window) error {
		return w.Map()
	}
}

type WindowAttribute func() (mask uint32, value uint32)

type WindowAttributes struct{}

//TODO CwBorderPixmap, CwBorderPixel, CwBitGravity, CwWinGravity, CwBackingStore, CwBackingPlanes, CwBackingPixel, CwOverrideRedirect , CwSaveUnder, CwEventMask, CwDontPropagate, CwColormap, CwCursor
func (_ WindowAttributes) BackgroundPixelColor(col color.Color) WindowAttribute {
	r, g, b, _ := col.RGBA()
	return func() (uint32, uint32) {
		return xproto.CwBackPixel, uint32(r>>8)<<16 + uint32(g>>8)<<8 + uint32(b>>8)
	}
}
func (_ WindowAttributes) BackgroundPixmap(pixmap *Pixmap) WindowAttribute {
	return func() (uint32, uint32) {
		return xproto.CwBackPixmap, uint32(pixmap.Pixmap)
	}
}
