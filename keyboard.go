package xgo

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgb/xtest"
)

//var keyLog *log.Logger = log.New(ioutil.Discard, "keys: ", log.LstdFlags)
var keyLog *log.Logger = log.New(os.Stderr, "keys: ", log.LstdFlags)

var keyWriteMx *sync.Mutex = &sync.Mutex{}

type Keyboard struct {
	w *Window
	c *KeyboardControll
}

func (k *Keyboard) Window() *Window {
	if k.w == nil {
		keyLog.Fatalf("Keyboard %s has no window", k)
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
	return fmt.Sprintf("k%s", k.w)
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

func (c *KeyboardControll) String() string {
	return fmt.Sprintf("c%s", c.k)
}

func (c *KeyboardControll) Screen() *Screen {
	if c.k == nil {
		keyLog.Fatalf("KeyboardControll: Screen: no Keyboard")
	}
	return c.k.Screen()
}

func (c *KeyboardControll) Display() *Display {
	return c.Screen().Display()
}

func (c *KeyboardControll) keyCode(kc xproto.Keycode) []xproto.Keysym {
	keyMap, err := xproto.GetKeyboardMapping(c.Display().Conn, kc, 1).Reply()
	if err != nil {
		keyLog.Fatalf("KeyboardControll: keyCode error: %s", err)
	}
	return keyMap.Keysyms
}

func (c *KeyboardControll) keyCodes(ks xproto.Keysym) map[xproto.Keycode][]xproto.Keysym {

	min, max := c.Display().Setup().MinKeycode, c.Display().Setup().MaxKeycode
	//keyLog.Printf("xproto.GetKeyboardMapping")
	keyMap, err := xproto.GetKeyboardMapping(c.Display().Conn, min, byte(max-min+1)).Reply()
	if err != nil {
		keyLog.Fatalf("KeyboardControll: keyCodes error: %s", err)
	}

	res := map[xproto.Keycode][]xproto.Keysym{}

	for kc := int(min); kc <= int(max); kc++ {
		keyCode := xproto.Keycode(kc)
		start := (kc - int(min)) * int(keyMap.KeysymsPerKeycode)
		for k := 0; k < int(keyMap.KeysymsPerKeycode); k++ {
			i := start + k
			if ks == keyMap.Keysyms[i] {
				res[keyCode] = keyMap.Keysyms[start : start+int(keyMap.KeysymsPerKeycode)]
				break
			}
		}
	}
	return res
}

func (c *KeyboardControll) keyCodeInput(kc xproto.Keycode, up bool) error {
	keyLog.Printf("KeyboardControll: keyCodeInput: start %v, %v", kc, up)
	defer keyLog.Printf("KeyboardControll: keyCodeInput: end")

	if err := c.Display().extension("xtest"); err != nil {
		return err
	}

	keyLog.Printf("KeyboardControll: keyCodeInput: keycode %s", c.printKeyCode(int(kc)))

	t := byte(xproto.KeyPress)
	if up {
		t = xproto.KeyRelease
	}
	keyLog.Printf("xtest.FakeInputChecked")
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
	return nil
}

func (c *KeyboardControll) unusedKeyCodes() []xproto.Keycode {
	min, max := c.Display().Setup().MinKeycode, c.Display().Setup().MaxKeycode
	keyMap, err := xproto.GetKeyboardMapping(c.Display().Conn, min, byte(max-min+1)).Reply()
	if err != nil {
		return nil
	}

	res := []xproto.Keycode{}
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
			res = append(res, xproto.Keycode(kc))
		}
	}
	return res
}

func (c *KeyboardControll) getUnusedKeyCode() (xproto.Keycode, error) {
	min, max := c.Display().Setup().MinKeycode, c.Display().Setup().MaxKeycode
	keyMap, err := xproto.GetKeyboardMapping(c.Display().Conn, min, byte(max-min+1)).Reply()
	if err != nil {
		keyLog.Fatal(err)
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

func (c *KeyboardControll) printKeyCode(kc int) string {
	min, max := c.Display().Setup().MinKeycode, c.Display().Setup().MaxKeycode
	keyMap, _ := xproto.GetKeyboardMapping(c.Display().Conn, min, byte(max-min+1)).Reply()
	start := (kc - int(min)) * int(keyMap.KeysymsPerKeycode)
	kckss := []string{}
	for c := 0; c < int(keyMap.KeysymsPerKeycode); c++ {
		i := start + c
		kckss = append(kckss, ksString[keyMap.Keysyms[i]])
	}
	return fmt.Sprintf("%d: %v", kc, kckss)
}

func (c *KeyboardControll) printKeyCodes() {
	min, max := c.Display().Setup().MinKeycode, c.Display().Setup().MaxKeycode
	keyMap, _ := xproto.GetKeyboardMapping(c.Display().Conn, min, byte(max-min+1)).Reply()
	keyLog.Printf("min: %d, max: %d, ksPkc: %d", min, max, keyMap.KeysymsPerKeycode)
	for kc := int(min); kc <= int(max); kc++ {
		c.printKeyCode(kc)
	}
}

func (c *KeyboardControll) remapKeyCode(kc xproto.Keycode, kss []xproto.Keysym) error {
	keyLog.Printf("KeyboardControll: remapKeyCode: start %v, %v", kc, kss)
	defer keyLog.Printf("KeyboardControll: remapKeyCode: end")

	stop := make(chan struct{})
	defer close(stop)
	mappingNotifyChannel := c.Display().Events().listenMappingNotify(stop)

	ckmCookie := xproto.ChangeKeyboardMappingChecked(c.Display().Conn, 1, kc, 1, kss)
	if err := ckmCookie.Check(); err != nil {
		return fmt.Errorf("x change keyboard mapping: %s", err)
	}

	//wait for mapping notify with ckmCookie Sequence or 100ms
	d := 100 * time.Millisecond
	alarm := time.After(d)
	for {
		select {
		case <-alarm:
			keyLog.Printf("KeyboardControll: remapKeyCode: no mapping notify event for %v", d)
			return nil
		case mne, ok := <-mappingNotifyChannel:
			if !ok {
				keyLog.Printf("KeyboardControll: remapKeyCode: mapping notify channel closed")
				mappingNotifyChannel = nil
			}
			if mne.Sequence == ckmCookie.Sequence {
				keyLog.Printf("KeyboardControll: remapKeyCode: got my mapping notify event")
				return nil
			}
			keyLog.Printf("KeyboardControll: remapKeyCode: got some mapping notify event")
		}
	}
	return fmt.Errorf("KeyboardControll: remapKeyCode: impossible error")
}

func (c *KeyboardControll) keySymInput(ks xproto.Keysym, up bool) (err error) {
	//keyLog.Printf("KeyboardControll: keySymInput: start %x, %v", ks, up)
	//defer keyLog.Printf("KeyboardControll: keySymInput: end")

	kcs := c.keyCodes(ks)
	//keyLog.Printf("KeyboardControll: keySymInput: keyCodes: %v", kcs)

	var (
		kc  xproto.Keycode
		kss []xproto.Keysym
	)
	// get first keycode with our keysym as first or second keysym on key
	for k, v := range kcs {
		if v[0] == ks || v[1] == ks {
			kc, kss = k, v
			break
		}
	}

	// no keycode, need to do temporary mapping
	if kss == nil {
		return fmt.Errorf("keysym \"%s\" not mapped to first two keysyms of keycodes", ks)
	}

	if kss[0] != ks {
		// our key is used with shift
		err = c.keySymInput(keysym["Shift_L"], up)
		if err != nil {
			return
		}
	}
	return c.keyCodeInput(kc, up)
}

func (c *KeyboardControll) press(ks xproto.Keysym) error {
	return c.keySymInput(ks, false)
}

func (c *KeyboardControll) relase(ks xproto.Keysym) error {
	return c.keySymInput(ks, true)
}

// if ks is not mapped, maps ks to first unused keycode, or last keycode if no unused
// returns function that restores mapping to original, or nil if no restore required
func (c *KeyboardControll) mapKeySym(ks xproto.Keysym) (restore func()) {
	// check if mapped
	for _, kss := range c.keyCodes(ks) {
		for i := range kss {
			if i > 1 {
				//TODO use modifiers
				// look only for first two, other modifiers not implemented
				break
			}
			if kss[i] == ks {
				return nil
			}
		}
	}

	original := []xproto.Keysym{0}
	tmpKeyCode, err := c.getUnusedKeyCode()
	if err != nil {
		tmpKeyCode = c.Display().Setup().MaxKeycode
		original = c.keyCode(tmpKeyCode)
	}

	//https://bbs.archlinux.org/viewtopic.php?pid=1715630#p1715630
	tmpKeySyms := []xproto.Keysym{ks, ks}

	if err := c.remapKeyCode(tmpKeyCode, tmpKeySyms); err != nil {
		//TODO what do with error?
		return nil
	}

	return func() {
		//TODO what do with error?
		c.remapKeyCode(tmpKeyCode, original)
	}
}

type keyAction interface {
	action(keyboarder) error
}

type errNoKeysymForUtf struct {
	r       rune
	replace keyAction
}

func (e errNoKeysymForUtf) Replace() keyAction {
	return e.replace
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

			if a == '-' {
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
		return nil, errNoKeysymForUtf{r, keySymStroke{keysym["emptyset"]}}
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
			if r, ok := err.(interface {
				Replace() keyAction
			}); ok {
				res = append(res, r.Replace())
				continue
			}
			return nil, err
		}
		res = append(res, a)
	}

	return res, nil
}

func (c *KeyboardControll) unmappedKeySymsInActions(actions []keyAction) []xproto.Keysym {
	res := []xproto.Keysym{}
	used := map[xproto.Keysym]bool{}
	for _, a := range actions {
		if ka, ok := a.(interface {
			KeySyms() []xproto.Keysym
		}); ok {
			for _, ks := range ka.KeySyms() {
				if !used[ks] {
					res = append(res, ks)
					used[ks] = true
				}
			}
		}
	}
	return res
}

func (c *KeyboardControll) Write(s string) error {
	actions, err := keyActionsFromString(s)
	if err != nil {
		return err
	}
	//keyLog.Printf("actions from \"%s\": %v", s, actions)
	keyWriteMx.Lock()
	defer keyWriteMx.Unlock()

	mappedKeyCodes := []xproto.Keycode{}
	defer func() {
		if len(mappedKeyCodes) > 0 {
			for _, kc := range mappedKeyCodes {
				c.remapKeyCode(kc, []xproto.Keysym{0})
			}
		}
	}()

	unusedKeyCodes := c.unusedKeyCodes()
	if len(unusedKeyCodes) > 0 {
		//leave one empty for sure
		unusedKeyCodes = unusedKeyCodes[1:]
	}
	//keyLog.Printf("unusedKeyCodes: %v", unusedKeyCodes)
	unmappedKeySyms := c.unmappedKeySymsInActions(actions)
	//keyLog.Printf("unmappedKeySyms: %v", unmappedKeySyms)
	if len(unmappedKeySyms) > 0 && len(unusedKeyCodes) > 0 {
		for i := range unusedKeyCodes {
			if len(unmappedKeySyms) == 0 {
				break
			}
			if err := c.remapKeyCode(unusedKeyCodes[i], []xproto.Keysym{unmappedKeySyms[0], unmappedKeySyms[0]}); err != nil {
				return err
			}
			mappedKeyCodes = append(mappedKeyCodes, unusedKeyCodes[i])
			unmappedKeySyms = unmappedKeySyms[1:]
		}
	}
	//keyLog.Printf("mappedKeyCodes: %v", mappedKeyCodes)
	if len(mappedKeyCodes) > 0 {
		time.Sleep(500 * time.Millisecond)
	}

	for _, a := range actions {
		if err := a.action(c); err != nil {
			return err
		}
		time.Sleep(50 * time.Millisecond)
	}
	return nil
}
