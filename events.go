package xgo

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

var eventsLog *log.Logger = func() *log.Logger {
	writer := io.Writer(ioutil.Discard)
	if DEBUG {
		writer = os.Stderr
	}
	return log.New(writer, "events: ", log.LstdFlags)
}()

type eventsControlFunc func(e *events)

type events struct {
	evt chan xgb.Event
	err <-chan xgb.Error

	mx      *sync.Mutex
	running chan struct{}
	control chan interface{}

	wls map[byte]map[xproto.Window]map[chan<- xgb.Event]xgb.Event
	ols map[byte]map[chan<- xgb.Event]xgb.Event
}

func (e *events) run() {
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

func (e *events) process(events <-chan xgb.Event, control <-chan interface{}) error {
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
			case xproto.ClientMessageEvent:
				eventsLog.Println("process: for: got ClientMessageEvent:", event)
				for evch, _ := range e.ols[xproto.ClientMessage] {
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

func (e *events) registerEventWindowListener(ev byte, w xproto.Window, l chan<- xgb.Event, init xgb.Event) {
	if _, ok := e.wls[ev]; !ok {
		e.wls[ev] = map[xproto.Window]map[chan<- xgb.Event]xgb.Event{}
	}
	if _, ok := e.wls[ev][w]; !ok {
		e.wls[ev][w] = map[chan<- xgb.Event]xgb.Event{}
	}
	e.wls[ev][w][l] = init
}

func (e *events) unregisterEventWindowListener(ev byte, w xproto.Window, l chan<- xgb.Event) {
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

func (e *events) listenMotionNotify(w *Window, stop <-chan struct{}) <-chan xproto.MotionNotifyEvent {
	le := make(chan xgb.Event)

	rcf := eventsControlFunc(func(e *events) {
		a, err := w.AttributesInfo()
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

	ucf := eventsControlFunc(func(e *events) {
		e.unregisterEventWindowListener(xproto.MotionNotify, w.Window, le)
		if _, ok := e.wls[xproto.MotionNotify][w.Window]; !ok {
			eventsLog.Printf("listenMotionNotify: ucf: no MotionNotify event listeners for window %s\n", w)
			a, err := w.AttributesInfo()
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

func (e *events) registerEventListener(ev byte, l chan<- xgb.Event, init xgb.Event) {
	if _, ok := e.ols[ev]; !ok {
		e.ols[ev] = map[chan<- xgb.Event]xgb.Event{}
	}
	e.ols[ev][l] = init
}

func (e *events) unregisterEventListener(ev byte, l chan<- xgb.Event) {
	if l != nil {
		close(l)
		delete(e.ols[ev], l)
	}

	if len(e.ols[ev]) == 0 {
		delete(e.ols, ev)
	}
}

func (e *events) listenMappingNotify(stop <-chan struct{}) <-chan xproto.MappingNotifyEvent {
	le := make(chan xgb.Event)

	rcf := eventsControlFunc(func(e *events) {
		e.registerEventListener(xproto.MappingNotify, le, xproto.MappingNotifyEvent{})
	})

	ucf := eventsControlFunc(func(e *events) {
		e.unregisterEventListener(xproto.MappingNotify, le)
	})

	e.run()
	eventsLog.Printf("listenMappingNotify: trying to register MappingNotify event listener %v\n", le)
	e.control <- rcf
	eventsLog.Printf("listenMappingNotify: registered MappingNotify event listener %v\n", le)

	ret := make(chan xproto.MappingNotifyEvent)
	go func() {
		defer close(ret)

		eventsLog.Printf("listenMappingNotify: translator %v: start\n", le)
		defer eventsLog.Printf("listenMappingNotify: translator %v: end\n", le)

		unregister := func(control chan<- interface{}) {
			eventsLog.Printf("listenMappingNotify: translator %v: unregister: start", le)
			defer eventsLog.Printf("listenMappingNotify: translator %v: unregister: end", le)
			//TODO rewrite according to CloseNotify
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

func (e *events) listenWmDeleteWindow(w *Window, stop <-chan struct{}) <-chan struct{} {
	eventChannel := make(chan xgb.Event)

	registerControlFunction := eventsControlFunc(func(e *events) {
		e.registerEventListener(xproto.ClientMessage, eventChannel, xproto.ClientMessageEvent{})
	})

	unregisterControlFunction := eventsControlFunc(func(e *events) {
		e.unregisterEventListener(xproto.ClientMessage, eventChannel)
	})

	e.run()

	eventsLog.Printf("listenWmDeleteWindow: trying to register ClientMessage event listener %v\n", eventChannel)
	e.control <- registerControlFunction
	eventsLog.Printf("listenWmDeleteWindow: registered ClientMessage event listener %v\n", eventChannel)

	ret := make(chan struct{})
	//TODO who will wait for this function close on exit?
	go func() {
		defer close(ret)

		eventsLog.Printf("\tlistenWmDeleteWindow: go func(): start")
		defer eventsLog.Printf("\tlistenWmDeleteWindow: go func(): end")

		unregisterFunction := func(control chan<- interface{}) {
			eventsLog.Printf("\tlistenWmDeleteWindow: go func: unregisterFunction(): start")
			defer eventsLog.Printf("\tlistenWmDeleteWindow: go func: unregisterFunction(): end")

			for eventChannel != nil {
				eventsLog.Printf("\tlistenWmDeleteWindow: go func: unregisterFunction: for loop")
				select {
				case _, ok := <-eventChannel:
					eventsLog.Printf("\tlistenWmDeleteWindow: go func: unregisterFunction: select got event")
					if !ok {
						eventsLog.Printf("\tlistenWmDeleteWindow: go func: unregisterFunction: event channel was closed")
						eventChannel = nil
						break
					}
					eventsLog.Printf("\tlistenWmDeleteWindow: go func: unregisterFunction: wtf am i doing here?")

				case control <- unregisterControlFunction:
					eventsLog.Printf("\tlistenWmDeleteWindow: go func: unregisterFunction: select unregisterControlFunction sent to control")
					control = nil
				}
			}
		}

		for stop != nil {
			eventsLog.Printf("\tlistenWmDeleteWindow: go func: for loop")
			select {
			case eventInterface, ok := <-eventChannel:
				eventsLog.Printf("\tlistenWmDeleteWindow: go func: select got event: %#v, %v", eventInterface, ok)
				if !ok {
					eventsLog.Printf("\tlistenWmDeleteWindow: go func: event channel was closed")
					eventChannel = nil
					break
				}

				switch event := eventInterface.(type) {
				case xproto.ClientMessageEvent:
					eventsLog.Printf("\tlistenWmDeleteWindow: go func: got ClientMessageEvent event")
					// check if event is an WM_DELETE_WINDOW message
					if event.Format != 32 {
						// unknown format, do nothing
						eventsLog.Printf("\tlistenWmDeleteWindow: go func: ClientMessageEvent.Format=%d, want 32", event.Format)
						break
					}

					if typeName, err := w.Screen().Display().atoms().GetById(event.Type); err != nil {
						eventsLog.Printf("\tlistenWmDeleteWindow: go func: error getting atom name fo id %d: %v", event.Type, err)
					} else if typeName != "WM_PROTOCOLS" {
						// event type is not WM_PROTOCOLS
						eventsLog.Printf("\tlistenWmDeleteWindow: go func: event.Type name is not WM_PROTOCOLS")
						break
					}

					// event type is WM_PROTOCOLS, check first protocol name
					protocolId := xproto.Atom(event.Data.Data32[0])
					if protocolName, err := w.Screen().Display().atoms().GetById(protocolId); err != nil {
						eventsLog.Printf("\tlistenWmDeleteWindow: go func: error getting atom name fo id %d: %v", protocolId, err)
					} else if protocolName != "WM_DELETE_WINDOW" {
						eventsLog.Printf("\tlistenWmDeleteWindow: go func: first protocol name is not WM_DELETE_WINDOW")
						// first event protocol is not WM_DELETE_WINDOW
						break
					}

					// event protocol is WM_DELETE_WINDOW, window is closing
					//TODO maybe unregister all events for window for sure
					defer unregisterFunction(e.control)
					stop = nil
				default:
					eventsLog.Printf("\tlistenWmDeleteWindow: go func: unknown event type: %T", event)
				}

			case _, ok := <-stop:
				if !ok {
					eventsLog.Printf("\tlistenWmDeleteWindow: go func: stop channel closed")
					defer unregisterFunction(e.control)
					stop = nil
					break
				}
				eventsLog.Printf("\tlistenWmDeleteWindow: go func: wtf am i doing here? don't send anything to stop channel")
			}
		}

	}()

	return ret
}
