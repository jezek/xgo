package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
	"xgo"
)

func main() {
	d, err := xgo.OpenDisplay("")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Display name:", d.Name())
	fmt.Println("Number of screens:", d.NumberOfScreens())
	fmt.Println("Number of desktops:", d.NumberOfDesktops())
	fmt.Println("Default screen:", d.DefaultScreen().Id())
	fmt.Println()

	fmt.Println("Mouse info:")
	s, err := d.Pointer().Status()
	if err != nil {
		fmt.Println("Can't obtain mouse status due to error:", err)
	}
	fmt.Printf("x: %d, y: %d, screen: %d, buttons: %v\n", s.RootX, s.RootY, s.Root.Screen().Id(), s.Button)
	a := d.ActiveWindow()
	ws, err := a.Pointer().Status()
	if err != nil {
		fmt.Println("Can't obtain mouse status for window %s due to error:", a, err)
	}
	fmt.Println(s)
	fmt.Println(ws)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	stop := make(chan struct{})
	mnch := a.Pointer().MotionNotify(stop)

	fmt.Println("Listening to pointer motion notify in active window (hit CTRL-C to exit):")
	fmt.Printf("x: %4d, y: %4d\n", s.RootX, s.RootY)
	for mnch != nil {
		select {
		case <-signalChannel:
			fmt.Println("ctrl-c")
			close(stop)
			stop = nil
			signalChannel = nil
		case mn, ok := <-mnch:
			if !ok {
				if stop != nil {
					close(stop)
					stop = nil
				}
				mnch = nil
				break
			}
			fmt.Printf("x: %4d, y: %4d\n", mn.RootX, mn.RootY)
			fmt.Println(mn)
		}
	}
	fmt.Println("Done")
	time.Sleep(time.Second)
}
