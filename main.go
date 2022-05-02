package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var brick = "*" //"█"

func main() {

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
