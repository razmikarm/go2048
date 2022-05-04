package main

import (
	"fmt"
)

var brick string
var lineSep string
var linePad string

// Basic line that board consists of
type line struct {
	items []int
	score int
}

// Moves elements to the Left
func (l *line) move() {
	cursor := 0
	for _, v := range l.items {
		if v != 0 {
			if l.items[cursor] == 0 {
				l.items[cursor] = v
			} else if v == l.items[cursor] {
				l.items[cursor] *= 2
				cursor += 1
			} else {
				cursor += 1
				l.items[cursor] = v
			}
		}
	}
	l.score = 0
	for _, v := range l.items {
		l.score += v
	}
}

//Reverses line items without changing the original line
func (l *line) reversed() line {
	c := len(l.items)
	nl := line{make([]int, c), 0}
	for i := 0; i < c/2; i++ {
		nl.items[i], nl.items[c-i-1] = l.items[c-i-1], l.items[i]
	}
	return nl
}

// Moves elements to the Right
func (l *line) reverseMove() {
	nl := l.reversed()
	nl.move()
	l.items = l.reversed().items
}

func (l *line) show() {
	fmt.Println(linePad)
	fmt.Print(brick)
	for _, v := range l.items {
		fmt.Print(center(v, 6), brick)
	}
	fmt.Println()
	fmt.Println(linePad)
	fmt.Println(lineSep)
}
