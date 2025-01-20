package ansi

import "fmt"

const (
	RED    Color = "255;0;0"
	GREEN  Color = "0;255;0"
	BLUE   Color = "0;0;255"
	ORANGE Color = "255;128;0"
	YELLOW Color = "255;255;0"
	CYAN   Color = "0;255;255"
	INDIGO Color = "127;0;255"
	PINK   Color = "255;0;255"
	BLACK  Color = "0;0;0"
	WHITE  Color = "255;255;255"
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

func SPrintC(s string, c Color) string {
	return "\x1b[38;2;" + string(c) + "m" + s + "\x1b[0m"
}

func SPrintLnC(s string, c Color) string {
	return "\x1b[38;2;" + string(c) + "m" + s + "\x1b[0m\n"
}

func SPrintBC(s string, c Color, b Color) string {
	return "\x1b[38;2;" + string(c) + "m" + "\x1b[48;2;" + string(b) + "m" + s + "\x1b[0m"
}
