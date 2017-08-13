package main

import (
	"fmt"
	"log"
	"time"
	"xgo"
)

func main() {
	d, err := xgo.OpenDisplay("")
	if err != nil {
		log.Fatal(err)
		return
	}

	//awkc := d.ActiveWindow().Screen().Window().Keyboard().Control()
	awkc := d.ActiveWindow().Keyboard().Control()

	//fmt.Println("Writing to active window's root window.")
	fmt.Println("Writing to active window")

	if err := awkc.Write("%\"aacute\""); err != nil {
		fmt.Println("Write error:", err)
	}
	//awkc.Write("あbč%\"Enter\"")

	time.Sleep(time.Second)
	fmt.Println("Done.")
	return

	fmt.Println("Sending Ctrl+C to active window's root window.")

	//awkc.Down(0xffe3)
	//awkc.Stroke(0x0063)
	//awkc.Up(0xffe3)
	awkc.Write("\uffe3c%-\"Control_L\"")

	time.Sleep(time.Second)
	fmt.Println("Done.")
}
