package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/eiannone/keyboard"
)

var brick = "█"
var lineSep string   //= strings.Repeat(brick, 29)
var lineSpace string //= strings.Repeat(brick+"      ", 4) + brick

func main() {

	if len(os.Args) > 1 {
		if utf8.RuneCountInString(os.Args[1]) > 1 {
			fmt.Println("Enter one character to fill borders")
			return
		} else {
			brick = os.Args[1]
		}
	}

	lineSep = strings.Repeat(brick, 29)
	lineSpace = strings.Repeat(brick+"      ", 4) + brick

	board := initBoard()
	draw(board)

	_, key, err := keyboard.GetSingleKey()
	for ; err == nil; _, key, err = keyboard.GetSingleKey() {
		switch key {
		case 0xffea:
			toRight(board)
		case 0xffeb:
			toLeft(board)
		case 0xffec:
			toDown(board)
		case 0xffed:
			toUp(board)
		case 0x3:
			clear()
			return
		default:
			continue
		}
		addNew(board)
		draw(board)
	}
	clear()
	fmt.Println(err)
}
