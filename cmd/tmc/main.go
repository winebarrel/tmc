package main

import (
	"fmt"

	"github.com/chzyer/readline"
	"github.com/winebarrel/tmc"
)

func main() {
	rl, err := readline.NewEx(&readline.Config{Prompt: "> "})

	if err != nil {
		panic(err)
	}

	defer rl.Close()

	for {
		line, err := rl.Readline()

		if err != nil {
			break
		}

		dur, err := tmc.Eval(line)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(tmc.DurToStr(dur))
	}
}
