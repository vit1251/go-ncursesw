// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows
// +build !windows

package goncurses

// #include <ncurses.h>
import "C"

import (
	"errors"
	"syscall"
)

var errList = map[C.int]string{}

func ncursesError(e error) error {
	errno, ok := e.(syscall.Errno)
	if int(errno) == C.OK {
		e = nil
	}
	if ok {
		errstr, ok := errList[C.int(errno)]
		if ok {
			return errors.New(errstr)
		}
	}
	return e
}
