package main

import (
	"fmt"
	"math/rand"
	"os"
)

// Main game board
type board struct {
	lines []*line
	size  int
}

// Generates new board
func NewBoard(size int) (nb board) {
	nb.size = size
	nb.lines = make([]*line, 0, size)
	for i := 0; i < size; i++ {
		nb.lines = append(nb.lines, NewLine(size))
	}
	return
}

// Adds new point in the random free cell
func (b *board) add() {
	empties := make([][]int, 0, b.size)
	for i, l := range b.lines {
		for j, v := range l.items {
			if v == 0 {
				empties = append(empties, []int{i, j})
			}
		}
	}
	if len(empties) == 0 {
		b.over()
	}
	r := rand.Intn(len(empties))
	i, j := empties[r][0], empties[r][1]
	b.lines[i].items[j] = 2
}

/*---------------------------------------------------*/
/*-----------| HERE GOES THE ENGINE PART |-----------*/
/*---------------------------------------------------*/

// Returns transposed board without changing the original
func (b *board) transposed() board {
	nb := NewBoard(b.size)

	for i := range b.lines {
		nl := NewLine(b.size)
		for j := range b.lines {
			nl.items[j] = b.lines[j].items[i]
		}
		nb.lines[i] = nl
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

/*---------------------------------------------------*/
/*-----------| HERE ENDS THE ENGINE PART |-----------*/
/*---------------------------------------------------*/

// Returns score of the board
func (b *board) score() (s int) {
	for _, l := range b.lines {
		s += l.score()
	}
	return
}

// Draws board in CLI
func (b *board) draw() {
	clearCLI()
	fmt.Println()
	fmt.Println(lineSep)
	for _, l := range b.lines {
		l.draw()
	}
	fmt.Printf("\nSCORE | \033[1m%d\033[0m\n\n", b.score())
}

func (b *board) over() {
	clearCLI()
	b.draw()
	fmt.Println("GAME OVER! SCORE | ", b.score())
	os.Exit(0)
}
