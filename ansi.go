package ansi

import "fmt"

const (
	RED Color = "255;0;0"
)

type Color string

func PrintC(s string, c Color) {
	fmt.Printf("\x1b[%sm%s\x1b[0m", c, s)
}

func rintLnC(s string, c Color) {
	fmt.Printf("\x1b[%sm%s\x1b[0m\n", c, s)
}
