package xgo

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgb/xtest"
)

// Display instance.
type Display struct {
	*xgb.Conn
	name string

	mx    *sync.Mutex
	setup *Setup
	ss    []*Screen
	p     *DisplayPointer
	e     *Events
	ext   map[string]error
}

// Creates a new Display instance.
//
// Parameter d contains display name, such as ":0".
// If empty, it uses the environment variable DISPLAY.
//
// returns nil on failure
func OpenDisplay(d string) (*Display, error) {
	if d == "" {
		d = os.Getenv("DISPLAY")
	}

	conn, err := xgb.NewConnDisplay(d)
	if err != nil {
		return nil, err
	}
	ret := &Display{conn, d, &sync.Mutex{}, nil, nil, nil, nil, make(map[string]error)}
	debugf("Opened display: %v", ret)

	return ret, nil
}

func (d *Display) Name() string {
	return d.name
}

func (d *Display) Setup() *Setup {
	d.mx.Lock()
	defer d.mx.Unlock()
	if d.setup == nil {
		d.setup = &Setup{xproto.Setup(d.Conn), d}
	}
	return d.setup
}

func (d *Display) Screens() []*Screen {
	d.mx.Lock()
	defer d.mx.Unlock()
	if d.ss == nil {
		sis := d.Setup().Roots
		d.ss = make([]*Screen, len(sis))
		ds := d.Conn.DefaultScreen
		for i := range d.ss {
			d.ss[i] = &Screen{&sis[i], d, i, i == ds, nil}
		}
	}
	return d.ss
}

func (d *Display) ActiveWindow() *Window {
	aname := "_NET_ACTIVE_WINDOW"
	activeAtom, err := xproto.InternAtom(d.Conn, true, uint16(len(aname)), aname).Reply()
	if err != nil {
		log.Fatal(err)
	}
	if activeAtom.Atom == xproto.AtomNone {
		log.Fatalf("No %s intern atom found", aname)
	}

	root := d.DefaultScreen().Window()

	reply, err := xproto.GetProperty(d.Conn, false, root.Window, activeAtom.Atom,
		xproto.GetPropertyTypeAny, 0, (1<<32)-1).Reply()
	if err != nil {
		log.Fatal(err)
	}
	return &Window{
		xproto.Window(xgb.Get32(reply.Value)), root.s,
		nil,
	}
}

func (d *Display) DefaultScreen() *Screen {
	ds := d.Conn.DefaultScreen
	return &Screen{&d.Setup().Roots[ds], d, ds, true, nil}
}

func (d *Display) NumberOfDesktops() uint32 {
	aname := "_NET_NUMBER_OF_DESKTOPS"
	req, err := xproto.InternAtom(d.Conn, true, uint16(len(aname)), aname).Reply()
	if err != nil {
		log.Fatal(err)
	}
	if req.Atom == xproto.AtomNone {
		log.Fatalf("No %s intern atom found", aname)
	}

	root := d.DefaultScreen().Window()

	reply, err := xproto.GetProperty(d.Conn, false, root.Window, req.Atom, xproto.GetPropertyTypeAny, 0, (1<<32)-1).Reply()
	if err != nil {
		log.Fatal(err)
	}
	if reply.ValueLen == 0 {
		return 0
	}

	return xgb.Get32(reply.Value)
}

func (d *Display) NumberOfScreens() int {
	return len(d.Setup().Roots)
}

func (d *Display) Pointer() *DisplayPointer {
	d.mx.Lock()
	defer d.mx.Unlock()
	if d.p == nil {
		d.p = &DisplayPointer{d}
	}
	return d.p
}

func (d *Display) FindWindow(wid uint32) (*Window, error) {
	xwid := xproto.Window(wid)
	reply, err := xproto.QueryTree(d.Conn, xwid).Reply()
	if err != nil {
		return nil, err
	}
	var ws *Screen
	for _, s := range d.Screens() {
		if s.Window().Window == reply.Root {
			ws = s
		}
	}
	if ws == nil {
		return nil, fmt.Errorf("Can't find screen for root window: %d", reply.Root)
	}
	return &Window{
		xwid, ws,
		nil,
	}, nil
}

func (d *Display) Events() *Events {
	d.mx.Lock()
	defer d.mx.Unlock()
	if d.e == nil {
		d.e = &Events{
			d.Conn,
			&sync.Mutex{},
			nil,
			nil,
			map[byte]map[xproto.Window]map[chan<- xgb.Event]xgb.Event{},
		}
	}
	return d.e
}

func (d *Display) extension(s string) error {
	d.mx.Lock()
	defer d.mx.Unlock()
	switch s {
	case "xtest":
		if err, ok := d.ext[s]; ok {
			return err
		}
		err := xtest.Init(d.Conn)
		d.ext[s] = err
		return err
	default:
		return fmt.Errorf("Unknown extension '%s'", s)
	}
	return nil
}

// Setup instance
type Setup struct {
	*xproto.SetupInfo
	d *Display
}
