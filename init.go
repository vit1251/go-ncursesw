
package goncurses

// #cgo !darwin,!openbsd,!windows pkg-config: ncursesw
// #cgo windows CFLAGS: -DNCURSES_MOUSE_VERSION
// #cgo windows LDFLAGS: -lpdcurses
// #cgo darwin openbsd LDFLAGS: -lncursesw
// #include <ncurses.h>
// #include "goncurses.h"
import "C"

func init() {
    C.ncurses_init()
}
