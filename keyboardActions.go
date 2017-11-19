package xgo

import (
	"fmt"
	"time"

	"github.com/BurntSushi/xgb/xproto"
)

type keyboarder interface {
	press(xproto.Keysym) error
	relase(xproto.Keysym) error
	mapKeySym(xproto.Keysym) func()
}

type keySymDown struct {
	ks xproto.Keysym
}

func (a keySymDown) action(k keyboarder) error {
	if restore := k.mapKeySym(a.ks); restore != nil {
		time.Sleep(time.Second)
		defer func() {
			//TODO don't know why, but I have to wait before mapping back
			//otherwise the key will mostly not be pressed
			//can I do this better?
			time.Sleep(100 * time.Millisecond)
			restore()
		}()
	}
	return k.press(a.ks)
}

func (a keySymDown) String() string {
	return fmt.Sprintf("d{%s}", keySymToString(a.ks))
}

type keySymUp struct {
	ks xproto.Keysym
}

func (a keySymUp) action(k keyboarder) error {
	if restore := k.mapKeySym(a.ks); restore != nil {
		defer func() {
			//TODO don't know why, but I have to wait before mapping back
			//otherwise the key will mostly not be pressed
			//can I do this better?
			time.Sleep(100 * time.Millisecond)
			restore()
		}()
	}
	return k.relase(a.ks)
}

func (a keySymUp) String() string {
	return fmt.Sprintf("u{%s}", keySymToString(a.ks))
}

type keySymStroke struct {
	ks xproto.Keysym
}

func (a keySymStroke) action(k keyboarder) error {
	if restore := k.mapKeySym(a.ks); restore != nil {
		time.Sleep(time.Second)
		defer func() {
			//TODO don't know why, but I have to wait before mapping back
			//otherwise the key will mostly not be pressed
			//can I do this better?
			time.Sleep(100 * time.Millisecond)
			restore()
		}()
	}
	if err := k.press(a.ks); err != nil {
		return err
	}
	return k.relase(a.ks)
}

func (a keySymStroke) String() string {
	return fmt.Sprintf("du{%s}", keySymToString(a.ks))
}

func keySymToString(ks xproto.Keysym) string {
	if kss := ksString[ks]; kss != "" {
		return kss
	}

	return fmt.Sprintf("0x%X(%v)", ks, ks)
}
