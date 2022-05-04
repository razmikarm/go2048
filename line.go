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
}

// Generates new line and
// returns a pointer to it
func NewLine(size int) *line {
	return &line{make([]int, size)}
}

/*---------------------------------------------------*/
/*-----------| HERE GOES THE ENGINE PART |-----------*/
/*---------------------------------------------------*/

// Moves elements to the Left
func (l *line) move() {
	cursor := 0
	nl := NewLine(len(l.items))
	for _, v := range l.items {
		if v != 0 {
			if nl.items[cursor] == 0 {
				nl.items[cursor] = v
			} else if v == nl.items[cursor] {
				nl.items[cursor] *= 2
				cursor += 1
			} else {
				cursor += 1
				nl.items[cursor] = v
			}
		}
	}
	*l = *nl
}

//Reverses line items without changing the original line
func (l *line) reversed() line {
	c := len(l.items)
	nl := NewLine(c)
	for i := 0; i < c/2; i++ {
		nl.items[i], nl.items[c-i-1] = l.items[c-i-1], l.items[i]
	}
	return *nl
}

// Moves elements to the Right
func (l *line) reverseMove() {
	nl := l.reversed()
	nl.move()
	*l = nl.reversed()
}

/*---------------------------------------------------*/
/*-----------| HERE ENDS THE ENGINE PART |-----------*/
/*---------------------------------------------------*/

// Returns line score
func (l *line) score() (s int) {
	for _, v := range l.items {
		s += v
	}
	return
}

// Draws line in CLI
func (l *line) draw() {
	fmt.Println(linePad)
	fmt.Print(brick)
	for _, v := range l.items {
		fmt.Print(center(v, 6), brick)
	}
	fmt.Println()
	fmt.Println(linePad)
	fmt.Println(lineSep)
}
