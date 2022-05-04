package main

import (
	"fmt"
	"strings"
)

// Main game board
type board struct {
	lines []*line
	size  int
}

// Initializes new board
func newBoard(size int) (nb board) {
	brick = "█"
	lineSep = strings.Repeat(brick, 29)
	linePad = strings.Repeat(brick+"      ", 4) + brick

	nb.size = size
	nb.lines = make([]*line, size)
	return
}

// Returns score of the board
func (b *board) score() (s int) {
	for _, l := range b.lines {
		s += l.score
	}
	return
}

// Returns transposed board without changing the original
func (b *board) transposed() board {
	nb := newBoard(b.size)

	for i := range b.lines {
		nl := line{}
		nl.items = make([]int, b.size)
		for j := range b.lines {
			nl.items[j] = b.lines[j].items[i]
		}
		nb.lines[i] = &nl
	}
	return nb
}

// Makes left move
func (b *board) toLeft() {
	for _, l := range b.lines {
		l.move()
	}
}

// Makes right move
func (b *board) toRight() {
	for _, l := range b.lines {
		l.reverseMove()
	}
}

// Makes up move
func (b *board) toUp() {
	nb := b.transposed()
	for _, l := range nb.lines {
		l.move()
	}
	b.lines = nb.transposed().lines
}

// Makes down move
func (b *board) toDown() {
	nb := b.transposed()
	for _, l := range nb.lines {
		l.reverseMove()
	}
	b.lines = nb.transposed().lines
}

func (b *board) show() {
	clear()
	fmt.Println(lineSep)
	for _, l := range b.lines {
		l.show()
	}
	fmt.Printf("\nSCORE | \033[1m%d\033[0m\n\n", b.score())
}
