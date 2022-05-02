package main

func lineToLeft(l []int) []int {
	newLine := make([]int, len(l))
	cursor := 0
	for _, v := range l {
		if v != 0 {
			if newLine[cursor] == 0 {
				newLine[cursor] = v
			} else if v == newLine[cursor] {
				newLine[cursor] *= 2
				cursor += 1
			} else {
				cursor += 1
				newLine[cursor] = v
			}
		}
	}
	return newLine
}

func toLeft(board [][]int) {
	for i, l := range board {
		newLine := lineToLeft(l)
		board[i] = newLine
	}
}

func toRight(board [][]int) {
	for i, l := range board {
		newLine := lineToLeft(reverseLine(l))
		board[i] = reverseLine(newLine)
	}
}

func toUp(board [][]int) {
	newBoard := transpose(board)
	for i, l := range newBoard {
		newLine := lineToLeft(l)
		newBoard[i] = newLine
	}
	newBoard = transpose(newBoard)
	copy(board, newBoard)
}

func toDown(board [][]int) {
	newBoard := transpose(board)
	for i, l := range newBoard {
		newLine := lineToLeft(reverseLine(l))
		newBoard[i] = reverseLine(newLine)
	}
	newBoard = transpose(newBoard)
	copy(board, newBoard)

}
