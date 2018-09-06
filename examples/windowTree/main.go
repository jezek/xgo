// Prints visible named window tree for all screens.
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

	fmt.Println("Visible window with name tree:")
	fmt.Print(d.Screens())
	for i, s := range d.Screens() {
		fmt.Println("Screen", i)
		printWithCildren(s.Window(), 0)
	}
}

func printWithCildren(w *xgo.Window, depth int) {
	if true { //w.IsVisible() && w.Name() != "" {
		for i := 0; i < depth-1; i++ {
			fmt.Print("| ")
		}
		if depth > 0 {
			fmt.Print("|- ")
		}
		a, _ := w.AttributesInfo()
		fmt.Printf("%s - \"%s\" dnpem: %b\n", w, w.Name(), a.GetWindowAttributesReply.DoNotPropagateMask)
		//fmt.Printf("%[1]*s %#v\n", depth*2, "", a.GetWindowAttributesReply.DoNotPropagateMask)
	}
	chs, err := w.Children()
	if err != nil {
		log.Print(err)
	}
	for _, ch := range chs {
		printWithCildren(ch, depth+1)
	}
}
