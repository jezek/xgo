// Golang wrapper fo xdo
// sudo apt install libxdo-dev to compile
package xgo

import (
	"log"
	"os"
)

var (
	DEBUG bool = os.Getenv("XGO_DEBUG") != ""
)

func debug(a ...interface{}) {
	if DEBUG {
		log.Print(a...)
	}
}
func debugf(f string, a ...interface{}) {
	if DEBUG {
		log.Printf(f, a...)
	}
}
