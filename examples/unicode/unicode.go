package main

import (
	"github.com/vit1251/goncurses"
	"log"
	"time"
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
	stdscr.Print("┌─ Title ──────────────────────┐")
	stdscr.Print("│                              │")
	stdscr.Print("└──────────────────────────────┘")
	stdscr.Refresh()

	time.Sleep(1 * time.Minute)

}
