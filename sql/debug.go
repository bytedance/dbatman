package sql

import "fmt"

var debug bool = false

func DEBUG(i interface{}) {
	if debug {
		fmt.Printf("%v", i)
	}
}

func setDebug(dbg bool) {
	debug = dbg
}
