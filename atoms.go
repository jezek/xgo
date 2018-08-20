package xgo

import (
	"fmt"
	"log"
	"sync"

	"github.com/jezek/xgb/xproto"
)

type atoms struct {
	mx           *sync.RWMutex
	cachedByName map[string]xproto.Atom // cached created atoms by name
	cachedById   map[xproto.Atom]string // cached created atoms by atomId

	d *Display
}

func (a *atoms) Display() *Display {
	if a.d == nil {
		log.Fatalf("Atoms have no display")
	}
	return a.d
}

func (a *atoms) cache(name string, aid xproto.Atom) {
	a.mx.Lock()
	a.cachedByName[name] = aid
	a.cachedById[aid] = name
	a.mx.Unlock()
}

func (a *atoms) GetByName(name string) (xproto.Atom, error) {

	// first look into cache
	a.mx.RLock()
	if aid, ok := a.cachedByName[name]; ok {
		// found. unlock & return
		a.mx.RUnlock()
		return aid, nil
	}
	a.mx.RUnlock()

	// atom name not in cache, create
	reply, err := xproto.InternAtom(a.Display().Conn, false, uint16(len(name)), name).Reply()
	if err != nil {
		return 0, fmt.Errorf("error interning atom \"%s\": %v", name, err)
	}

	aid := reply.Atom
	if aid == 0 {
		return 0, fmt.Errorf("interning atom \"%s\" returned 0", name)
	}

	// new atom created, cache it
	a.cache(name, aid)

	return aid, nil
}

func (a *atoms) GetById(atomId xproto.Atom) (string, error) {

	// first look into cache
	a.mx.RLock()
	if name, ok := a.cachedById[atomId]; ok {
		// found. unlock & return
		a.mx.RUnlock()
		return name, nil
	}
	a.mx.RUnlock()

	// atomId not in cache, get name from server
	reply, err := xproto.GetAtomName(a.Display().Conn, atomId).Reply()
	if err != nil {
		return "", errWrap{"Atoms.GetById", fmt.Errorf("error getting name for Atom %v: %v", atomId, err)}
	}

	// cache pair
	a.cache(reply.Name, atomId)

	return reply.Name, nil
}
