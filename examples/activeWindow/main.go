// Prints active window's parents and children for default screeen.
package main

import (
	"fmt"
	"log"
	"xgo"
)

func main() {
	d, err := xgo.OpenDisplay("")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Display name:", d.Name())
	fmt.Println("Number of screens:", d.NumberOfScreens())
	fmt.Println("Number of desktops:", d.NumberOfDesktops())
	fmt.Println()

	fmt.Println("Active window status:")
	a := d.ActiveWindow()
	aps, err := a.Parents()
	if err != nil {
		log.Println("Can't get active window parents,", err)
	}
	achs, err := a.Children()
	if err != nil {
		log.Println("Can't get active window children,", err)
	}
	for i := len(aps) - 1; i >= 0; i-- {
		fmt.Printf("%s - %s\n", aps[i], aps[i].Name())
	}
	fmt.Printf("* active %s - %s\n", a, a.Name())
	for i := len(achs) - 1; i >= 0; i-- {
		fmt.Printf("|- %s - %s\n", achs[i], achs[i].Name())
	}
	fmt.Println()

}
