package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/eiannone/keyboard"
)

func init() {

	brick = "█"
	rand.Seed(time.Now().UnixMicro())
	if len(os.Args) > 1 {
		if utf8.RuneCountInString(os.Args[1]) > 1 {
			fmt.Println("Enter one character to fill borders")
			return
		} else {
			brick = os.Args[1]
		}
	}
	lineSep = strings.Repeat(brick, 29)
	linePad = strings.Repeat(brick+"      ", 4) + brick

}

func main() {

	b := NewBoard(4)
	b.add()
	b.draw()

	_, key, err := keyboard.GetSingleKey()
	for ; err == nil; _, key, err = keyboard.GetSingleKey() {
		switch key {
		case 0xffea:
			b.toRight()
		case 0xffeb:
			b.toLeft()
		case 0xffec:
			b.toDown()
		case 0xffed:
			b.toUp()
		case 0x3:
			clearCLI()
			return
		default:
			continue
		}
		b.add()
		b.draw()
	}
	clearCLI()
	fmt.Println(err)
}
