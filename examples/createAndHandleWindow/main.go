package main

import (
	"fmt"
	"os"
	"os/signal"
	"xgo"
)

func main() {
	fmt.Println("hello")
	d, err := xgo.OpenDisplay("")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Display opened: %#v\n", d)
	//TODO how to detect closed display?

	w, err := d.NewWindow()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Window created: %#v\n", w)
	defer func() {
		w.Destroy()
		fmt.Printf("Window destroyed: %#v\n", w)
	}()

	//if reply, err := xproto.ListProperties(w.Screen().Display().Conn, w.Window).Reply(); err != nil {
	//	fmt.Printf("window %v list properties error: %v\n", w, err)
	//} else {
	//	fmt.Printf("window %v properties: %v\n", w, reply)
	//}

	if err := w.Map(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Window mapped")

	closedCountdown := 4
	stop := make(chan struct{})
	done, err := w.CloseNotify(stop)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening to window close event")

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	fmt.Println("close window 5 times, or <ctrl-c> for instant end")

	for done != nil {
		select {
		case <-signalChannel:
			fmt.Println("\nloop select: got kill signal")
			// close stop channel
			close(stop)
			stop = nil
			signalChannel = nil
		case _, ok := <-done:
			fmt.Println("loop select: got something from widow done channel")
			if !ok {
				fmt.Println("done was closed, window manager requests to close the window")
				if signalChannel != nil {
					fmt.Println("closedCountdown:", closedCountdown)
					if closedCountdown > 0 {
						stop = make(chan struct{})
						done, err = w.CloseNotify(stop)
						if err != nil {
							fmt.Println(err)
							return
						}
						fmt.Println("refusing to close, listening to window close event again")
						closedCountdown--
					} else {
						fmt.Println("you got me, I'm quitting...")
						done = nil
						stop = nil
					}
				} else {
					fmt.Println("kill signal was trigerred, exit loop")
					done = nil
					stop = nil
				}
			} else {
				fmt.Println("wtf? this supposed not to happen")
			}
		}
	}
	fmt.Println("end")
}
