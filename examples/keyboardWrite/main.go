package main

import (
	"flag"
	"fmt"
	"log"
	"xgo"
)

var keySymString string

func main() {
	flag.Parse()

	keySymString = flag.Arg(0)

	if flag.NArg() != 1 {
		//flag.PrintDefaults()
		//os.Exit(1)
		keySymString = "aacute"
		//awkc.Write("あbč%\"Enter\"")
	}

	d, err := xgo.OpenDisplay("")
	if err != nil {
		log.Fatal(err)
		return
	}

	//awkc := d.ActiveWindow().Screen().Window().Keyboard().Control()
	awkc := d.ActiveWindow().Keyboard().Control()

	//fmt.Println("Writing to active window's root window.")
	fmt.Printf("Writing \"%s\" to active window: ", keySymString)

	if err := awkc.Write("%\"" + keySymString + "\""); err != nil {
		fmt.Println("Write error:", err)
	}

	fmt.Println("\nDone.")
}
