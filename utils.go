package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func addNew(board [][]int) {
	empties := [][]int{}
	for i, l := range board {
		for j, v := range l {
			if v == 0 {
				empties = append(empties, []int{i, j})
			}
		}
	}
	if len(empties) == 0 {
		over()
	}
	r := rand.Intn(len(empties))
	i := empties[r][0]
	j := empties[r][1]
	board[i][j] = 2
}

func initBoard() [][]int {
	rand.Seed(time.Now().UnixMicro())
	board := [][]int{
		make([]int, 4),
		make([]int, 4),
		make([]int, 4),
		make([]int, 4),
	}
	addNew(board)
	return board
}

func score(board [][]int) int {
	score := 0
	for _, l := range board {
		for _, v := range l {
			score += v
		}
	}
	return score
}

func over() {
	fmt.Println("GAME OVER!")
	os.Exit(0)
}

func reverseLine(line []int) []int {
	newLine := make([]int, 0, len(line))
	for i := range line {
		newLine = append(newLine, line[len(line)-i-1])
	}
	return newLine
}

func transpose(board [][]int) [][]int {
	size := len(board)
	newBoard := make([][]int, size)
	for i := range newBoard {
		newLine := make([]int, size)
		for j := range newBoard {
			newLine[j] = board[j][i]
		}
		newBoard[i] = newLine
	}
	return newBoard
}
