// Package rpi is a library for interfacing with Raspberry Pi IO
package rpi

// This is the wrong url here ???
import (
	_ "github.com/bakhtik/goPi/MCP23S17"
	_ "github.com/bakhtik/goPi/ioctl"
	_ "github.com/bakhtik/goPi/piface"
	_ "github.com/bakhtik/goPi/spi"
)
