package main

import (
	"fmt"
	"image"

	"github.com/jezek/xgo"
)

func main() {
	d, err := xgo.OpenDisplay("")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Display opened: %#v\n", d)

	w, err := d.NewWindow(
		xgo.WindowOperations{}.Size(image.Pt(100, 100)),
	)
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
	//Note:
	//if window is closed, so the connection is, and so the current program.
	//TODO examine the panic error produced: panic: close of closed channel
	//I have red, that this happens only on well behaved window managers
	//see createAndHandleWindow example for window close handling
	fmt.Println("<enter to end>")
	fmt.Scanln()

}
