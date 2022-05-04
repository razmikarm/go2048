package main

import (
	"fmt"
	"strconv"
)

// Returns centered bold text of v value with w width filled with spaces
func center(v int, w int) string {
	s := " "
	if v != 0 {
		s = strconv.Itoa(v)
	}
	return fmt.Sprintf("\033[1m%*s\033[0m", -w, fmt.Sprintf("%*s", (w+len(s))/2, s))
}

// Clears CLI
func clearCLI() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}
