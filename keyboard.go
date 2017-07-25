package xgo

import (
	"fmt"
	"log"

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

func (c *KeyboardControll) keySymInput(ks xproto.Keysym, up bool) error {
	for _, kc := range c.keyCodesFromKeySym(ks) {
		if err := c.keyCodeInput(kc, up); err != nil {
			return err
		}
	}
	return nil
}

func (c *KeyboardControll) Down(ks xproto.Keysym) error {
	return c.keySymInput(ks, false)
}

func (c *KeyboardControll) Up(ks xproto.Keysym) error {
	return c.keySymInput(ks, true)
}

func (c *KeyboardControll) Stroke(ks xproto.Keysym) error {
	if err := c.Down(ks); err != nil {
		return err
	}
	if err := c.Up(ks); err != nil {
		return err
	}
	return nil
}
