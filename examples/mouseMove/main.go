package main

import (
	"log"

	"github.com/jezek/xgo"
)

func main() {
	d, err := xgo.OpenDisplay("")
	if err != nil {
		log.Fatal(err)
		return
	}

	d.DefaultScreen().Window().Pointer().Control().MoveRelative(10, 10)
}
