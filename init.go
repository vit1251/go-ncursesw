
package goncurses

// #cgo !darwin,!openbsd,!windows pkg-config: ncurses
// #cgo windows CFLAGS: -DNCURSES_MOUSE_VERSION
// #cgo windows LDFLAGS: -lpdcurses
// #cgo darwin openbsd LDFLAGS: -lncurses
// #include <curses.h>
// #include "goncurses.h"
import "C"

func init() {
    C.ncurses_init()
}
