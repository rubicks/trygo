package main

import (
	"fmt"
	"log"

	"github.com/chzyer/readline"
)

func main() {
	fmt.Println("Gispy v0.0.0")
	fmt.Println("Press Ctrl+c to exit")
	rl, err := readline.New("gispy> ")
	if err != nil {
		log.Fatal(err)
	}
	defer rl.Close()
	for {
		text, err := rl.Readline()
		if nil != err {
			log.Fatal(err)
		}
		println(text)
	}
}
