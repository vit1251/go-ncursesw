# Overview

Goncurses is an unicode ncurses library for the Go programming language. It
requires both pkg-config and ncurses C development files be installed.

# Installation

The go tool is the recommended method of installing goncurses. Issue the
following command on the command line:

    $ go get github.com/vit1251/goncurses

## Linux compile errors

Since native nature of `ncurses` no problem is register.

## OSX compile errors

Most OSX version provide outdate curses C library version 5.3 and `goncurses`
unable operate with that outdate version.

You may use `brew` system to update your ncurses (recomended version >= 6.3).

    $ brew install ncurses

And setup you `pkg-config` directory as shown after:

    $ export PKG_CONFIG_PATH="/usr/local/opt/ncurses/lib/pkgconfig"

## Microsoft Windows compile errors

Since Microsoft Windows actually does not support POSIX subsystem and
terminal subsystem, use a C library such as `ncurses` does not make sence.

You can use more suitable WinAPI terminal library instead.

Initially, the author tried to add support for `pdncurses`, but this issue
needs to be resolved by higher-level libraries outside of support for the
`ncurses` library name.

Support for the Microsft Windows system has been removed and in the future
the library will issue an explicit build error on this platform.

## FreeBSD compile errors

No checks.
