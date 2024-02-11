package main

import (
	"bufio"
	"fmt"
	"github.com/Lutrome/PokemonCLI/commands"
	"os"
)

func main() {
	cmdParse := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("PokemonCLI v0.0.1 > ")
		cmdParse.Scan()
		if cmdParse.Text() == "help" {
			err := commands.Help()
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		commands.RunCommand(cmdParse.Text())
	}
}
