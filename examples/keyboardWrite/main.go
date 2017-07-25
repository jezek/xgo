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

	awkc := d.ActiveWindow().Screen().Window().Keyboard().Control()

	fmt.Println("Writing to active window's root window.")

	awkc.Stroke(0x0061)
	awkc.Stroke(0x0062)
	awkc.Stroke(0x0063)

	time.Sleep(time.Second)
	fmt.Println("Done.")

	fmt.Println("Sending Ctrl+C to active window's root window.")

	awkc.Down(0xffe3)
	awkc.Stroke(0x0063)
	awkc.Up(0xffe3)

	time.Sleep(time.Second)
	fmt.Println("Done.")
}
