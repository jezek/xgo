package xgo

import (
	"log"
	"os"
)

const (
	SizeReqestMax                = (1 << 16) * 4
	SizeRequestPutImageFixedPart = 24
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

type errWrap struct {
	msg string
	err error
}

func (e errWrap) Error() string {
	return e.msg + ": " + e.err.Error()
}

//TODO make window tests for diisplay, window, keyboard,... testing using user interacion and make them availablu only through a flag
//TODO test all non connection functions
