package xgo

import (
	"fmt"
	"log"
	"sync"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

type eventsControlFunc func(e *Events)

type Events struct {
	c *xgb.Conn

	mx      *sync.Mutex
	running chan struct{}
	control chan interface{}

	wls map[byte]map[xproto.Window]map[chan<- xgb.Event]xgb.Event
}

func (e *Events) run() {
	e.mx.Lock()
	fmt.Println("Running events")
	defer fmt.Println("Stopped events")
	defer e.mx.Unlock()

	if e.running != nil {
		fmt.Println("allready running")
		return
	}

	e.running = make(chan struct{})
	e.control = make(chan interface{})
	events := make(chan xgb.Event)

	go func(running chan struct{}) {
		fmt.Println("Running WaitForEvent router")
		defer fmt.Println("Stopped WaitForEvent router")
		for {
			evnt, err := e.c.WaitForEvent()
			if err != nil {
				log.Printf("waitForEvent error: %s\n", err)
				continue
			}
			fmt.Println("waitForEvent got event")
			select {
			case events <- evnt:
				fmt.Println("waitForEvent routed event")
			case <-running:
				fmt.Println("waitForEvent got stop")
				close(events)
				return
			}
		}
	}(e.running)

	go func(control chan interface{}) {
		e.process(events, control)
		close(e.running)
		defer close(control)

		e.mx.Lock()
		defer e.mx.Unlock()
		e.running = nil
		e.control = nil
	}(e.control)
}

func (e *Events) process(events <-chan xgb.Event, control <-chan interface{}) error {
	fmt.Println("Running process")
	defer func() {
		for _, wls := range e.wls {
			for _, ls := range wls {
				for lch, _ := range ls {
					close(lch)
				}
			}
		}
		fmt.Println("Ending process")
	}()

	for {
		fmt.Println("waiting process")
		select {
		case ev, ok := <-events:
			if !ok {
				fmt.Println("got Events channel closed")
				return nil
			}
			fmt.Println("got event", ev)
			switch event := ev.(type) {
			case xproto.ButtonPressEvent:
				fmt.Println("Got ButtonPressEvent:", event)
			case xproto.ButtonReleaseEvent:
				fmt.Println("Got ButtonReleaseEvent:", event)
			case xproto.MotionNotifyEvent:
				fmt.Println("Got MotionNotifyEvent:", event)
				for evch, _ := range e.wls[xproto.MotionNotify][event.Event] {
					evch <- event
					fmt.Printf("Event %v sent to window %d listener %v\n", event, event.Event, evch)
				}
			default:
				fmt.Println("Got some other event:", event)
			}
		case i, ok := <-control:
			if !ok {
				fmt.Println("got control closed")
				return nil
			}
			fmt.Println("got control", i)
			switch ctrl := i.(type) {
			case eventsControlFunc:
				fmt.Println("got eventsControlFunc", ctrl)
				ctrl(e)
				fmt.Println("eventsControlFunc executed")
			default:
				fmt.Println("got unidentified control")
			}
			if len(e.wls) == 0 {
				fmt.Println("no listeners")
				return nil
			}
		}
	}
	return nil
}

func (e *Events) registerEventWindowListener(ev byte, w xproto.Window, l chan<- xgb.Event, init xgb.Event) {
	if _, ok := e.wls[ev]; !ok {
		e.wls[ev] = map[xproto.Window]map[chan<- xgb.Event]xgb.Event{}
	}
	if _, ok := e.wls[ev][w]; !ok {
		e.wls[ev][w] = map[chan<- xgb.Event]xgb.Event{}
	}
	e.wls[ev][w][l] = init
}

func (e *Events) unregisterEventWindowListener(ev byte, w xproto.Window, l chan<- xgb.Event) {
	if l != nil {
		close(l)
		delete(e.wls[ev][w], l)
	}

	if len(e.wls[ev][w]) == 0 {
		delete(e.wls[ev], w)
	}

	if len(e.wls[ev]) == 0 {
		delete(e.wls, ev)
	}
}

func (e *Events) listenMotionNotify(w *Window, stop <-chan struct{}) <-chan xproto.MotionNotifyEvent {
	le := make(chan xgb.Event)

	rcf := eventsControlFunc(func(e *Events) {
		a, err := w.Attributes()
		if err != nil {
			log.Printf("Can't get window %s attributes due to error: %s", w, err)
			close(le)
			return
		}
		m := a.GetWindowAttributesReply.YourEventMask | xproto.EventMaskPointerMotion
		fmt.Printf("Setting pointer motion event mask %b for window %s\n", m, w)
		if err := xproto.ChangeWindowAttributesChecked(
			w.Screen().Display().Conn,
			w.Window,
			xproto.CwEventMask,
			[]uint32{m},
		).Check(); err != nil {
			log.Printf("Can't set window %s event mask %b due to error: %s", w, m, err)
			close(le)
			return
		}
		e.registerEventWindowListener(xproto.MotionNotify, w.Window, le, xproto.MotionNotifyEvent{})
	})

	ucf := eventsControlFunc(func(e *Events) {
		e.unregisterEventWindowListener(xproto.MotionNotify, w.Window, le)
		if _, ok := e.wls[xproto.MotionNotify][w.Window]; !ok {
			fmt.Printf("No MotionNotify event listeners for window %s\n", w)
			a, err := w.Attributes()
			if err != nil {
				log.Printf("Can't get window %s attributes due to error: %s", w, err)
			}
			m := a.GetWindowAttributesReply.YourEventMask
			nm := m & ^uint32(xproto.EventMaskPointerMotion)
			fmt.Printf("Changing event mask from %b to %b for window %s\n", m, nm, w)
			if err := xproto.ChangeWindowAttributesChecked(
				w.Screen().Display().Conn,
				w.Window,
				xproto.CwEventMask,
				[]uint32{nm},
			).Check(); err != nil {
				log.Printf("Can't set window %s event mask %b due to error: %s", w, nm, err)
				close(le)
			}
		}
	})

	e.run()
	fmt.Printf("Trying to register MotionNotify event listener %v\n", le)
	e.control <- rcf
	fmt.Printf("Registered MotionNotify event listener %v\n", le)
	ret := make(chan xproto.MotionNotifyEvent)
	go func() {
		fmt.Printf("Running listenMotionNotify translator %v\n", le)
		defer fmt.Printf("Closing listenMotionNotify translator %v\n", le)
		defer close(ret)

		unregister := func(control chan<- interface{}) {
			fmt.Printf("unregistering motion notify listener %v\n", le)
			defer fmt.Printf("unregistered motion notify listener%v\n", le)
			for {
				select {
				case _, ok := <-le:
					if !ok {
						return
					}
				case control <- ucf:
					fmt.Printf("sent request to unregister %v\n", le)
					control = nil
				}
			}
		}

		for {
			select {
			case ei, ok := <-le:
				if !ok {
					fmt.Printf("got MotionNotify event listener %v closed\n", le)
					return
				}
				switch event := ei.(type) {
				case xproto.MotionNotifyEvent:
					fmt.Printf("got MotionNotify event %v from listener %v\n", event, le)
					select {
					case ret <- event:
						fmt.Printf("forwarded from %v, to %v\n", le, ret)
					case _, ok := <-stop:
						if !ok {
							fmt.Printf("got stop while trying to forward %v\n", le)
							defer unregister(e.control)
							return
						}
					}
				default:
					log.Fatal("BUG: Expecting only motion notify events")
				}
			case _, ok := <-stop:
				if !ok {
					fmt.Printf("got stop while listening to %v\n", le)
					defer unregister(e.control)
					return
				}
			}
		}
	}()
	return ret
}
