package main

import (
	"fmt"
	"strconv"
)

func center(v int, w int) string {
	s := strconv.Itoa(v)
	return fmt.Sprintf("\033[1m%*s\033[0m", -w, fmt.Sprintf("%*s", (w+len(s))/2, s))
}

func clear() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

func draw(board [][]int) {
	clear()
	fmt.Println(lineSep)
	for _, l := range board {
		fmt.Println(lineSpace)
		fmt.Print(brick)
		for _, v := range l {
			fmt.Print(center(v, 6), brick)
		}
		fmt.Println()
		fmt.Println(lineSpace)
		fmt.Println(lineSep)
	}
	fmt.Printf("\nSCORE | \033[1m%d\033[0m\n\n", score(board))
}
