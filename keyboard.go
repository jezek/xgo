package xgo

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgb/xtest"
)

type Keyboard struct {
	w *Window
	c *KeyboardControll
}

func (k *Keyboard) Window() *Window {
	if k.w == nil {
		log.Fatalf("Poiner %s has no window", k)
	}
	return k.w
}

func (k *Keyboard) Screen() *Screen {
	return k.Window().Screen()
}

func (k *Keyboard) Display() *Display {
	return k.Screen().Display()
}

func (k *Keyboard) String() string {
	return fmt.Sprintf("k%s", k.Window())
}

func (k *Keyboard) Control() *KeyboardControll {
	if k.c == nil {
		k.c = &KeyboardControll{k}
	}
	return k.c
}

type KeyboardControll struct {
	k *Keyboard
}

func (c *KeyboardControll) Screen() *Screen {
	return c.k.Screen()
}

func (c *KeyboardControll) Display() *Display {
	return c.Screen().Display()
}

func (c *KeyboardControll) keyCodesFromKeySym(ks xproto.Keysym) []xproto.Keycode {

	min, max := c.Display().Setup().MinKeycode, c.Display().Setup().MaxKeycode
	keyMap, err := xproto.GetKeyboardMapping(c.Display().Conn, min, byte(max-min+1)).Reply()
	if err != nil {
		log.Fatal(err)
	}

	keyCodes := make([]xproto.Keycode, 0)
	set := make(map[xproto.Keycode]bool, 0)

	for kc := int(min); kc <= int(max); kc++ {
		keyCode := xproto.Keycode(kc)
		for c := 0; c < int(keyMap.KeysymsPerKeycode); c++ {
			if set[keyCode] {
				continue
			}
			i := (kc-int(min))*int(keyMap.KeysymsPerKeycode) + c
			if ks == keyMap.Keysyms[i] {
				keyCodes = append(keyCodes, keyCode)
				set[keyCode] = true
			}
		}
	}
	return keyCodes
}

func (c *KeyboardControll) keyCodeInput(kc xproto.Keycode, up bool) error {

	if err := c.Display().extension("xtest"); err != nil {
		return err
	}
	log.Printf("send keycode: %v", kc)
	t := byte(xproto.KeyPress)
	if up {
		t = xproto.KeyRelease
	}
	return xtest.FakeInputChecked(
		c.Display().Conn,
		t,
		byte(kc),
		xproto.TimeCurrentTime,
		c.k.Window().Window,
		0, 0, 0,
	).Check()
}

type errKeySymInput struct {
	s        xproto.Keysym
	kcl, kcp int
	kc       xproto.Keycode
	up       bool
	err      error
}

func (e errKeySymInput) Error() string {
	action := "press"
	if e.up {
		action = "relase"
	}
	sym := string(e.s)
	codepos := ""
	if e.kcl > 1 {
		codepos = string(e.kcp) + "th "
	}
	code := string(e.kc)

	return fmt.Sprintf("Can't %s symbol's %s %skeycode %d due to error \"%s\"", action, sym, codepos, code, e.err)
}

func (c *KeyboardControll) getUnusedKeyCode() (xproto.Keycode, error) {
	min, max := c.Display().Setup().MinKeycode, c.Display().Setup().MaxKeycode
	keyMap, err := xproto.GetKeyboardMapping(c.Display().Conn, min, byte(max-min+1)).Reply()
	if err != nil {
		log.Fatal(err)
	}

	for kc := int(min); kc <= int(max); kc++ {
		empty := true
		for c := 0; c < int(keyMap.KeysymsPerKeycode); c++ {
			i := (kc-int(min))*int(keyMap.KeysymsPerKeycode) + c
			if keyMap.Keysyms[i] != 0 {
				empty = false
				break
			}
		}
		if empty {
			return xproto.Keycode(kc), nil
		}
	}
	// not empty, use last
	return 0, fmt.Errorf("no free keycode")
}
func (c *KeyboardControll) keySymInput(ks xproto.Keysym, up bool) (err error) {
	log.Printf("keySymInput: %v, %v", ks, up)
	kcs := c.keyCodesFromKeySym(ks)
	log.Printf("keySymInput: codes: %v", kcs)

	if len(kcs) == 0 {
		// no keycode, need to do temporary mapping
		tmpKeyCode, uerr := c.getUnusedKeyCode()
		if uerr != nil {
			return fmt.Errorf("get unused key code: %s", uerr)
		}
		tmpCodes := []xproto.Keysym{ks}
		log.Printf("keySymInput: change mapping: %v, %v", tmpKeyCode, tmpCodes)
		if err = xproto.ChangeKeyboardMappingChecked(c.Display().Conn, 1, tmpKeyCode, 1, tmpCodes).Check(); err != nil {
			return fmt.Errorf("change keyboard mapping: %s", err)
		}
		c.Display().Conn.Sync()
		kcs = []xproto.Keycode{tmpKeyCode}
		defer func() {
			tmpCodes := []xproto.Keysym{0}
			log.Printf("keySymInput: change mapping back: %v, %v", tmpKeyCode, tmpCodes)
			if merr := xproto.ChangeKeyboardMappingChecked(c.Display().Conn, 1, tmpKeyCode, 1, tmpCodes).Check(); merr != nil {
				if err != nil {
					merr = fmt.Errorf("previous error: %s", err)
				}
				err = fmt.Errorf("change keyboard mapping back: %s", merr)
				return
			}
			c.Display().Conn.Sync()
		}()

	}

	if len(kcs) > 1 {
		log.Printf("no")
		return nil
	}
	for i, kc := range kcs {
		if err = c.keyCodeInput(kc, up); err != nil {
			log.Printf("keySymInput: code input error %s", err)
			err = errKeySymInput{ks, len(kcs), i, kc, up, err}
			return
		}
	}
	return
}

func (c *KeyboardControll) press(ks xproto.Keysym) error {
	return c.keySymInput(ks, false)
}

func (c *KeyboardControll) relase(ks xproto.Keysym) error {
	return c.keySymInput(ks, true)
}

type keyboarder interface {
	press(xproto.Keysym) error
	relase(xproto.Keysym) error
}

type keyAction interface {
	action(keyboarder) error
}

type keySymDown struct {
	ks xproto.Keysym
}

func (a keySymDown) action(k keyboarder) error {
	return k.press(a.ks)
}

func (a keySymDown) String() string {
	return fmt.Sprintf("d{0x%X}", a.ks)
}

type keySymUp struct {
	ks xproto.Keysym
}

func (a keySymUp) action(k keyboarder) error {
	return k.relase(a.ks)
}

func (a keySymUp) String() string {
	return fmt.Sprintf("u{0x%X}", a.ks)
}

type errInvalidAction struct {
	a byte
}

func (e errInvalidAction) Error() string {
	return fmt.Sprintf("Invalid action %%%s", string(e.a))
}

type errActionRead struct {
	err error
}

func (e errActionRead) Error() string {
	return fmt.Sprintf("Can't read action: %s", e.err)
}

type errRuneRead struct {
	err error
}

func (e errRuneRead) Error() string {
	return fmt.Sprintf("Can't read rune: %s", e.err)
}

type errInvalidKeySymString struct {
	s string
}

func (e errInvalidKeySymString) Error() string {
	return fmt.Sprintf("Invalid keysym string: %s", e.s)
}

func keyActionsFromString(s string) ([]keyAction, error) {
	if s == "" {
		return nil, nil
	}
	res := []keyAction{}
	buf := strings.NewReader(s)

	readRune := func(b *strings.Reader) (rune, error) {
		r, _, err := buf.ReadRune()
		if err != nil {
			return 0, err
		}
		if r == '%' {
			a, err := buf.ReadByte()
			log.Printf("readRune: %%: action: \"%v\", %s", string(a), err)
			if err != nil {
				return 0, errRuneRead{errActionRead{err}}
			}
			if a == '"' {
				str := ""
				for {
					c, err := buf.ReadByte()
					if err != nil {
						return 0, errRuneRead{err}
					}
					if c == '"' {
						break
					}
					str += string(c)
				}
				ks, ok := keysym[str]
				if !ok {
					return 0, errInvalidKeySymString{str}
				}
				res = append(res, keySymDown{xproto.Keysym(ks)}, keySymUp{xproto.Keysym(ks)})
			}
			buf.UnreadByte()
		}
		return r, nil
	}

	for {
		r, err := readRune(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errRuneRead{err}
		}
		log.Printf("kafs: rune: \"%v\", %v, %x", string(r), r, r)

		if r == '%' {
			a, err := buf.ReadByte()
			log.Printf("kafs: %%: action: \"%s\", %s", string(a), err)
			if err != nil {
				return nil, errActionRead{err}
			}
			switch a {
			case '+', '-':
				r, err := readRune(buf)
				log.Printf("kafs: %%+- rune: %s, %s", string(r), err)
				if err != nil {
					return nil, err
				}
				ks, ok := utf[r]
				if !ok {
					log.Printf("kafs: %%+- rune %s is not in convert table", string(r))
					break
				}
				var act keyAction
				if a == 'i' {
					act = keySymUp{xproto.Keysym(ks)}
				} else {
					act = keySymDown{xproto.Keysym(ks)}
				}
				res = append(res, act)
			case '%':
				res = append(res, keySymDown{xproto.Keysym('%')}, keySymUp{xproto.Keysym('%')})
			default:
				return nil, errInvalidAction{a}
			}
			continue
		}

		ks, ok := utf[r]
		if !ok {
			log.Printf("kafs: rune %s is not in convert table", r)
			continue
		}
		res = append(res, keySymDown{xproto.Keysym(ks)}, keySymUp{xproto.Keysym(ks)})
	}
	return res, nil
}

func (c *KeyboardControll) Write(s string) error {
	actions, err := keyActionsFromString(s)
	if err != nil {
		return err
	}
	for _, a := range actions {
		if err := a.action(c); err != nil {
			return err
		}
	}
	return nil
}
