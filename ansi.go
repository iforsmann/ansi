package ansi

import "fmt"

const (
	RED Color = "255;0;0"
)

type Color string

func PrintC(s string, c Color) {
	fmt.Printf("\x1b[38;2;" + string(c) + "m" + s + "\x1b[0m")
}

func PrintLnC(s string, c Color) {
	fmt.Printf("\x1b[38;2;" + string(c) + "m" + s + "\x1b[0m\n")
}

func PrintBC(s string, c Color, b Color) {
	fmt.Printf("\x1b[38;2;" + string(c) + "m" + "\x1b[48;2;" + string(b) + "m" + s + "\x1b[0m")
}
