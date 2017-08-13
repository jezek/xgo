package xgo

import (
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

//var eventsLog *log.Logger = log.New(os.Stderr, "events: ", log.LstdFlags)
var eventsLog *log.Logger = log.New(ioutil.Discard, "events: ", log.LstdFlags)

type eventsControlFunc func(e *Events)

type Events struct {
	evt chan xgb.Event
	err <-chan xgb.Error

	mx      *sync.Mutex
	running chan struct{}
	control chan interface{}

	wls map[byte]map[xproto.Window]map[chan<- xgb.Event]xgb.Event
	ols map[byte]map[chan<- xgb.Event]xgb.Event
}

func (e *Events) run() {
	t := time.Now()
	eventsLog.Println(t, "run: start")
	defer eventsLog.Println(t, "run: end")
	e.mx.Lock()
	eventsLog.Println(t, "run: locked")
	defer eventsLog.Println(t, "run: unlocked")
	defer e.mx.Unlock()

	if e.running != nil {
		eventsLog.Println(t, "run: allready running")
		return
	}

	e.running = make(chan struct{})
	e.control = make(chan interface{})
	events := make(chan xgb.Event)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		//e.waitForEventLoop(events, e.running)
		func(events chan<- xgb.Event, stop <-chan struct{}) {
			eventsLog.Println(t, "waitForEventLoop: start")
			defer eventsLog.Println(t, "waitForEventLoop: end")
			for {
				select {
				case xe := <-e.evt:
					eventsLog.Println(t, "waitForEventLoop: routing event", xe)
					select {
					case events <- xe:
						eventsLog.Println(t, "waitForEventLoop: routed")
					case <-stop:
						eventsLog.Println(t, "waitForEventLoop: got stop while routing")
						go func() {
							eventsLog.Println(t, "waitForEventLoop: gofunc: recycling", xe)
							e.evt <- xe
							eventsLog.Println(t, "waitForEventLoop: gofunc: recycled")
						}()
						return
					}
				case <-e.err:
				case <-stop:
					eventsLog.Println(t, "waitForEventLoop: got stop")
					return
				}
			}
		}(events, e.running)
		wg.Done()
	}()

	go func(control chan interface{}, running chan struct{}) {
		eventsLog.Println(t, "go process: start")
		defer eventsLog.Println(t, "go process: end")
		e.process(events, control)
		close(running)
		eventsLog.Println(t, "go process: ending")
		e.mx.Lock()
		eventsLog.Println(t, "go process: locked")

		close(control)
		e.control = nil

		e.running = nil

		eventsLog.Println(t, "go process: waiting for waitForEventLoop end")
		wg.Wait()
		e.mx.Unlock()
		eventsLog.Println(t, "go process: unlocked")
	}(e.control, e.running)
}

func (e *Events) process(events <-chan xgb.Event, control <-chan interface{}) error {
	eventsLog.Println("process: start")
	defer eventsLog.Println("process: end")
	defer func() {
		eventsLog.Println("process: clean up")
		for _, wls := range e.wls {
			for _, ls := range wls {
				for lch, _ := range ls {
					close(lch)
				}
			}
		}
		for _, ls := range e.ols {
			for lch, _ := range ls {
				close(lch)
			}
		}
	}()

	for {
		eventsLog.Println("process: for: waiting process")
		select {
		case ev, ok := <-events:
			if !ok {
				eventsLog.Println("process: for: got events channel closed")
				return nil
			}
			eventsLog.Println("process: for: got event", ev)
			switch event := ev.(type) {
			case xproto.ButtonPressEvent:
				eventsLog.Println("process: for: got ButtonPressEvent:", event)
			case xproto.ButtonReleaseEvent:
				eventsLog.Println("process: for: got ButtonReleaseEvent:", event)
			case xproto.MotionNotifyEvent:
				eventsLog.Println("process: for: got MotionNotifyEvent:", event)
				for evch, _ := range e.wls[xproto.MotionNotify][event.Event] {
					evch <- event
					eventsLog.Printf("process: for: event %v sent to window %d listener %v\n", event, event.Event, evch)
				}
			case xproto.MappingNotifyEvent:
				eventsLog.Println("process: for: got MappingNotifyEvent:", event)
				for evch, _ := range e.ols[xproto.MappingNotify] {
					evch <- event
					eventsLog.Printf("process: for: event %v sent to listener %v\n", event, evch)
				}
			default:
				eventsLog.Println("process: for: got some other event:", event)
			}
		case i, ok := <-control:
			if !ok {
				eventsLog.Println("process: for: got control closed")
				return nil
			}
			eventsLog.Println("process: for: got control", i)
			switch ctrl := i.(type) {
			case eventsControlFunc:
				eventsLog.Println("process: for: got eventsControlFunc", ctrl)
				ctrl(e)
				eventsLog.Println("process: for: eventsControlFunc executed")
			default:
				eventsLog.Println("process: for: got unidentified control")
			}
			if len(e.wls) == 0 && len(e.ols) == 0 {
				eventsLog.Println("process: for: no listeners, quit")
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
			eventsLog.Printf("listenMotionNotify: rcf: can't get window %s attributes due to error: %s", w, err)
			close(le)
			return
		}
		m := a.GetWindowAttributesReply.YourEventMask | xproto.EventMaskPointerMotion
		eventsLog.Printf("listenMotionNotify: rcf: setting pointer motion event mask %b for window %s\n", m, w)
		if err := xproto.ChangeWindowAttributesChecked(
			w.Screen().Display().Conn,
			w.Window,
			xproto.CwEventMask,
			[]uint32{m},
		).Check(); err != nil {
			eventsLog.Printf("listenMotionNotify: rcf: can't set window %s event mask %b due to error: %s", w, m, err)
			close(le)
			return
		}
		e.registerEventWindowListener(xproto.MotionNotify, w.Window, le, xproto.MotionNotifyEvent{})
	})

	ucf := eventsControlFunc(func(e *Events) {
		e.unregisterEventWindowListener(xproto.MotionNotify, w.Window, le)
		if _, ok := e.wls[xproto.MotionNotify][w.Window]; !ok {
			eventsLog.Printf("listenMotionNotify: ucf: no MotionNotify event listeners for window %s\n", w)
			a, err := w.Attributes()
			if err != nil {
				eventsLog.Printf("listenMotionNotify: ucf: can't get window %s attributes due to error: %s", w, err)
			}
			m := a.GetWindowAttributesReply.YourEventMask
			nm := m & ^uint32(xproto.EventMaskPointerMotion)
			eventsLog.Printf("listenMotionNotify: ucf: changing event mask from %b to %b for window %s\n", m, nm, w)
			if err := xproto.ChangeWindowAttributesChecked(
				w.Screen().Display().Conn,
				w.Window,
				xproto.CwEventMask,
				[]uint32{nm},
			).Check(); err != nil {
				eventsLog.Printf("listenMotionNotify: ucf: can't set window %s event mask %b due to error: %s", w, nm, err)
				close(le)
			}
		}
	})

	e.run()
	eventsLog.Printf("listenMotionNotify: trying to register MotionNotify event listener %v\n", le)
	e.control <- rcf
	eventsLog.Printf("listenMotionNotify: registered MotionNotify event listener %v\n", le)
	ret := make(chan xproto.MotionNotifyEvent)
	go func() {
		eventsLog.Printf("listenMotionNotify: translator %v: start\n", le)
		defer eventsLog.Printf("listenMotionNotify: translator %v: end\n", le)
		defer close(ret)

		unregister := func(control chan<- interface{}) {
			eventsLog.Printf("listenMotionNotify: translator %v: unregister: start", le)
			defer eventsLog.Printf("listenMotionNotify: translator %v: unregister: end", le)
			for {
				select {
				case _, ok := <-le:
					if !ok {
						return
					}
				case control <- ucf:
					eventsLog.Printf("listenMotionNotify: translator %v: unregister: sent", le)
					control = nil
				}
			}
		}

		for {
			select {
			case ei, ok := <-le:
				if !ok {
					eventsLog.Printf("listenMotionNotify: translator %v: got event listener closed\n", le)
					return
				}
				switch event := ei.(type) {
				case xproto.MotionNotifyEvent:
					eventsLog.Printf("listenMotionNotify: translator %v: got MotionNotifyEvent %v", le, event)
					select {
					case ret <- event:
						eventsLog.Printf("listenMotionNotify: translator %v: forwarded to %v", le, ret)
					case _, ok := <-stop:
						if !ok {
							eventsLog.Printf("listenMotionNotify: translator %v: got stop while forwarding", le)
							defer unregister(e.control)
							return
						}
					}
				default:
					eventsLog.Fatal("listenMotionNotify: translator %v: BUG: expecting only MotionNotifyEvent", le)
				}
			case _, ok := <-stop:
				if !ok {
					eventsLog.Printf("listenMotionNotify: translator %v: got stop while listening", le)
					defer unregister(e.control)
					return
				}
			}
		}
	}()
	return ret
}

func (e *Events) registerEventListener(ev byte, l chan<- xgb.Event, init xgb.Event) {
	if _, ok := e.ols[ev]; !ok {
		e.ols[ev] = map[chan<- xgb.Event]xgb.Event{}
	}
	e.ols[ev][l] = init
}

func (e *Events) unregisterEventListener(ev byte, l chan<- xgb.Event) {
	if l != nil {
		close(l)
		delete(e.ols[ev], l)
	}

	if len(e.ols[ev]) == 0 {
		delete(e.ols, ev)
	}
}

func (e *Events) listenMappingNotify(stop <-chan struct{}) <-chan xproto.MappingNotifyEvent {
	le := make(chan xgb.Event)

	rcf := eventsControlFunc(func(e *Events) {
		e.registerEventListener(xproto.MappingNotify, le, xproto.MappingNotifyEvent{})
	})

	ucf := eventsControlFunc(func(e *Events) {
		e.unregisterEventListener(xproto.MappingNotify, le)
	})

	e.run()
	eventsLog.Printf("listenMappingNotify: trying to register MappingNotify event listener %v\n", le)
	e.control <- rcf
	eventsLog.Printf("listenMappingNotify: registered MappingNotify event listener %v\n", le)
	ret := make(chan xproto.MappingNotifyEvent)
	go func() {
		eventsLog.Printf("listenMappingNotify: translator %v: start\n", le)
		defer eventsLog.Printf("listenMappingNotify: translator %v: end\n", le)
		defer close(ret)

		unregister := func(control chan<- interface{}) {
			eventsLog.Printf("listenMappingNotify: translator %v: unregister: start", le)
			defer eventsLog.Printf("listenMappingNotify: translator %v: unregister: end", le)
			for {
				select {
				case _, ok := <-le:
					if !ok {
						return
					}
				case control <- ucf:
					eventsLog.Printf("listenMappingNotify: translator %v: unregister: sent", le)
					control = nil
				}
			}
		}

		for {
			select {
			case ei, ok := <-le:
				if !ok {
					eventsLog.Printf("listenMappingNotify: translator %v: got event listener closed", le)
					return
				}
				switch event := ei.(type) {
				case xproto.MappingNotifyEvent:
					eventsLog.Printf("listenMappingNotify: translator %v: got MappingNotifyEvent %v", le, event)
					select {
					case ret <- event:
						eventsLog.Printf("listenMappingNotify: translator %v: forwarded to %v", le, ret)
					case _, ok := <-stop:
						if !ok {
							eventsLog.Printf("listenMappingNotify: translator %v: got stop while forwarding", le)
							defer unregister(e.control)
							return
						}
					}
				default:
					eventsLog.Fatal("listenMappingNotify: translator %v: BUG: Expecting only MappingNotifyEvent", le)
				}
			case _, ok := <-stop:
				if !ok {
					eventsLog.Printf("listenMappingNotify: translator %v: got stop while listening", le)
					defer unregister(e.control)
					return
				}
			}
		}
	}()
	return ret
}
