package goncurses

// #cgo !windows pkg-config: ncursesw
// #include <ncurses.h>
// #include "goncurses.h"
import "C"

func init() {
	C.ncurses_init()
}
