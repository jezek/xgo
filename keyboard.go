package xgo

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"time"

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
	//log.Printf("xproto.GetKeyboardMapping")
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
	log.Printf("send keycode: %v, up: %v", kc, up)
	c.printKeyCode(int(kc))
	t := byte(xproto.KeyPress)
	if up {
		t = xproto.KeyRelease
	}
	log.Printf("xtest.FakeInputChecked")
	if err := xtest.FakeInputChecked(
		c.Display().Conn,
		t,
		byte(kc),
		xproto.TimeCurrentTime,
		c.k.Window().Window,
		0, 0, 0,
	).Check(); err != nil {
		return err
	}
	log.Printf("sent keycode: %v, up: %v", kc, up)
	return nil
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

func (c *KeyboardControll) printKeyCode(kc int) {
	min, max := c.Display().Setup().MinKeycode, c.Display().Setup().MaxKeycode
	keyMap, _ := xproto.GetKeyboardMapping(c.Display().Conn, min, byte(max-min+1)).Reply()
	start := (kc - int(min)) * int(keyMap.KeysymsPerKeycode)
	kckss := []string{}
	for c := 0; c < int(keyMap.KeysymsPerKeycode); c++ {
		i := start + c
		kckss = append(kckss, ksString[keyMap.Keysyms[i]])
	}
	log.Printf("kc:%d = %v", kc, kckss)
}

func (c *KeyboardControll) printKeyCodes() {
	min, max := c.Display().Setup().MinKeycode, c.Display().Setup().MaxKeycode
	keyMap, _ := xproto.GetKeyboardMapping(c.Display().Conn, min, byte(max-min+1)).Reply()
	log.Printf("min: %d, max: %d, ksPkc: %d", min, max, keyMap.KeysymsPerKeycode)
	for kc := int(min); kc <= int(max); kc++ {
		c.printKeyCode(kc)
	}
}

func (c *KeyboardControll) waitForMappingNotifyWithSequence(seq uint16, mnch <-chan xproto.MappingNotifyEvent) bool {
	log.Printf("waitForMappingNotifyWithSequence: %v", seq)
	for {
		select {
		case <-time.After(time.Second):
			log.Printf("waitForMappingNotifyWithSequence: no mapping notify event for second")
			return false
		case mne, ok := <-mnch:
			if !ok {
				log.Printf("waitForMappingNotifyWithSequence: mapping notify channel closed")
				return false
			}
			if mne.Sequence == seq {
				log.Printf("waitForMappingNotifyWithSequence: got seq")
				return true
			}
			log.Printf("waitForMappingNotifyWithSequence: got !seq")
		}
	}
	return false
}

func (c *KeyboardControll) keySymInput(ks xproto.Keysym, up bool) (err error) {
	//log.Printf("keySymInput: %x, %v", ks, up)
	kcs := c.keyCodesFromKeySym(ks)
	//log.Printf("keySymInput: codes: %v", kcs)

	if len(kcs) == 0 {
		// no keycode, need to do temporary mapping
		tmpKeyCode, uerr := c.getUnusedKeyCode()
		if uerr != nil {
			return fmt.Errorf("get unused key code: %s", uerr)
		}

		tmpCodes := []xproto.Keysym{ks}

		stop := make(chan struct{})
		defer close(stop)
		mappingNotifyChannel := c.Display().Events().listenMappingNotify(stop)

		log.Printf("keySymInput: change mapping: %v, %v", tmpKeyCode, tmpCodes)
		ckmCookie := xproto.ChangeKeyboardMappingChecked(c.Display().Conn, 1, tmpKeyCode, 1, tmpCodes)
		if err = ckmCookie.Check(); err != nil {
			return fmt.Errorf("change keyboard mapping: %s", err)
		}

		if ok := c.waitForMappingNotifyWithSequence(ckmCookie.Cookie.Sequence, mappingNotifyChannel); !ok {
			return fmt.Errorf("change keyboard mapping: no mapping notify after")
		}

		defer func() {
			//TODO don't know why, but I have to wait before mapping back
			//otherwise the key will mostly not be pressed
			//can I do this better?
			time.Sleep(50 * time.Millisecond)
			tmpCodes := []xproto.Keysym{0}

			log.Printf("keySymInput: change mapping back: %v, %v", tmpKeyCode, tmpCodes)
			ckmCookie := xproto.ChangeKeyboardMappingChecked(c.Display().Conn, 1, tmpKeyCode, 1, tmpCodes)
			if merr := ckmCookie.Check(); merr != nil {
				if err != nil {
					merr = fmt.Errorf("previous error: %s", err)
				}
				err = fmt.Errorf("change keyboard mapping back: %s", merr)
				return
			}

			if ok := c.waitForMappingNotifyWithSequence(ckmCookie.Cookie.Sequence, mappingNotifyChannel); !ok {
				if err != nil {
					err = fmt.Errorf("change keyboard mapping: no mapping notify after: previous error: %s", err)
				} else {
					err = fmt.Errorf("change keyboard mapping: no mapping notify after")
				}
			}
		}()

		kcs = []xproto.Keycode{tmpKeyCode}

	}

	if len(kcs) > 1 {
		log.Printf("no")
		return nil
	}
	//TODO modify keys
	for i, kc := range kcs {
		if err = c.keyCodeInput(kc, up); err != nil {
			//log.Printf("keySymInput: code input error %s", err)
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

type keySymStroke struct {
	ks xproto.Keysym
}

func (a keySymStroke) action(k keyboarder) error {
	if err := k.press(a.ks); err != nil {
		return err
	}
	return k.relase(a.ks)
}

func (a keySymStroke) String() string {
	return fmt.Sprintf("du{0x%X}", a.ks)
}

type errNoKeysymForUtf struct {
	r rune
}

func (e errNoKeysymForUtf) Error() string {
	return fmt.Sprintf("no keysym for rune: %v", e.r)
}

type errInvalidAction struct {
	a byte
}

func (e errInvalidAction) Error() string {
	return fmt.Sprintf("invalid action: %%%s", string(e.a))
}

type errInvalidKeyAction struct {
	a keyAction
}

func (e errInvalidKeyAction) Error() string {
	return fmt.Sprintf("invalid key action after %%+/-: %v", e.a)
}

type errActionRead struct {
	err error
}

func (e errActionRead) Error() string {
	return fmt.Sprintf("can't read action: %s", e.err)
}

type errInvalidKeySymString struct {
	s string
}

func (e errInvalidKeySymString) Error() string {
	return fmt.Sprintf("invalid keysym string: %s", e.s)
}

func readAction(b *bytes.Buffer) (keyAction, error) {
	r, _, err := b.ReadRune()
	if err != nil {
		return nil, err
	}
	if r == '%' {
		a, err := b.ReadByte()
		if err != nil {
			return nil, errActionRead{err}
		}
		switch a {
		case '"':
			ksb, err := b.ReadBytes('"')
			if err != nil {
				return nil, err
			}
			ksb = ksb[:len(ksb)-1]
			ks, ok := keysym[string(ksb)]
			if !ok {
				return nil, errInvalidKeySymString{string(ksb)}
			}
			return keySymStroke{ks}, nil
		case '+', '-':
			ka, err := readAction(b)
			if err != nil {
				return nil, err
			}

			var ks xproto.Keysym
			switch t := ka.(type) {
			case keySymStroke:
				ks = t.ks
			default:
				return nil, errInvalidKeyAction{ka}
			}

			if a == 'i' {
				return keySymUp{ks}, nil
			}
			return keySymDown{ks}, nil
		case '%':
		default:
			return nil, errInvalidAction{a}
		}
	}
	ks, ok := utf[r]
	if !ok {
		return nil, errNoKeysymForUtf{r}
	}
	return keySymStroke{ks}, nil
}

func keyActionsFromString(s string) ([]keyAction, error) {
	if s == "" {
		return nil, nil
	}
	res := []keyAction{}
	buf := bytes.NewBuffer([]byte(s))

	for {
		a, err := readAction(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		res = append(res, a)
	}

	return res, nil
}

func (c *KeyboardControll) Write(s string) error {
	actions, err := keyActionsFromString(s)
	if err != nil {
		return err
	}
	log.Printf("actions from \"%s\": %v", s, actions)
	for _, a := range actions {
		if err := a.action(c); err != nil {
			return err
		}
	}
	return nil
}
