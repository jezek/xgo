package xgo

import (
	"log"
	"os"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

// Display instance.
type Display struct {
	*xgb.Conn
	name  string
	setup *Setup
	ss    []*Screen
	m     *Mouse
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
	ret := &Display{conn, d, nil, nil, nil}
	debugf("Opened display: %v", ret)

	return ret, nil
}

func (d *Display) Name() string {
	return d.name
}

func (d *Display) Setup() *Setup {
	if d.setup == nil {
		d.setup = &Setup{d, xproto.Setup(d.Conn)}
	}
	return d.setup
}

func (d *Display) Screens() []*Screen {
	if d.ss == nil {
		sis := d.Setup().i.Roots
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
		nil, nil,
	}
}

func (d *Display) DefaultScreen() *Screen {
	ds := d.Conn.DefaultScreen
	return &Screen{&d.Setup().i.Roots[ds], d, ds, true, nil}
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
	return len(d.Setup().i.Roots)
}

func (d *Display) Mouse() *Mouse {
	if d.m == nil {
		d.m = &Mouse{d}
	}
	return d.m
}

// Setup instance
type Setup struct {
	d *Display
	i *xproto.SetupInfo
}
