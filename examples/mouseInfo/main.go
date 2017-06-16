package main

import (
	"fmt"
	"log"
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
	s, err := d.Mouse().Status()
	if err != nil {
		fmt.Println("Can't obtain mouse status due to error:", err)
	}
	fmt.Printf("x: %d, y: %d, screen: %d, buttons: %v\n", s.X, s.Y, s.Screen.Id(), s.Button)
}
