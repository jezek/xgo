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

	w, err := d.CreateWindow()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Window created: %#v\n", w)

	if err := w.Map(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Window mapped")

	stop := make(chan struct{})
	done := w.CloseNotify(stop)
	fmt.Println("Listening to window close event")

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	fmt.Println("<ctrl-c to end>")

	for done != nil {
		select {
		case <-signalChannel:
			fmt.Println("select got signal")
			close(stop)
			stop = nil
			signalChannel = nil
		case _, ok := <-done:
			fmt.Println("select go done")
			if !ok {
				fmt.Println("done was closed")
				if stop != nil {
					close(stop)
					stop = nil
				}
				done = nil
				break
			}
			fmt.Println("wtf? this supposed not to happen")
		}
	}
	fmt.Println("end")

}
