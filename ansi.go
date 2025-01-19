package ansi

import "fmt"

const (
	RED Color = "255;0;0"
)

type Color string

func printC(s string, c Color) {
	fmt.Printf("\x1b[%sm%s\x1b[0m", c, s)
}

func printlnC(s string, c Color) {
	fmt.Printf("\x1b[%sm%s\x1b[0m\n", c, s)
}
