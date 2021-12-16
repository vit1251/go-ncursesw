package main

import (
	"github.com/vit1251/goncurses"
	"log"
	"time"
	"fmt"
)

func main() {

	stdscr, err1 := goncurses.Init()
	if err1 != nil {
		log.Fatal("init:", err1)
	}
	defer goncurses.End()

	err2 := goncurses.StartColor()
	if err2 != nil {
		log.Fatal("StartColor", err2)
	}

	/* Box Drawing unicode example */
	stdscr.Printf("FMT_HORIZ = '─'\n")
	stdscr.Printf("FMT_VERT = '│'\n")
	stdscr.Refresh()

	time.Sleep(500 * time.Millisecond)

	fmt.Printf("FMT_HORIZ = '─'\n")
	fmt.Printf("FMT_VERT = '│'\n")

}
