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
