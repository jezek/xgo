package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jezek/xgo"
)

var keySymString string
var formated bool

func main() {
	flag.BoolVar(&formated, "formated", false, "Argumet is a formated string to write")
	flag.Parse()

	keySymString = flag.Arg(0)

	if flag.NArg() != 1 {
		keySymString = "date%\"Return\""
		//TODO not working, only printing b (and Return)
		//keySymString = "あbč%\"Return\""
		formated = true
	}

	d, err := xgo.OpenDisplay("")
	if err != nil {
		log.Fatal(err)
		return
	}

	//awkc := d.ActiveWindow().Screen().Window().Keyboard().Control()
	awkc := d.ActiveWindow().Keyboard().Control()

	t := keySymString
	if !formated {
		t = "%\"" + t + "\""
	}
	fmt.Printf("Writing \"%s\" to active window: ", t)
	if err := awkc.Write(t); err != nil {
		fmt.Println("Write error:", err)
	}

	fmt.Println("\nDone.")
}
