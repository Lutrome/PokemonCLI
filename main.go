package main

import (
	"bufio"
	"fmt"
	"github.com/Lutrome/PokemonCLI/utils"
	"os"
)

func main() {
	cmdParse := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("PokemonCLI v0.0.1 > ")
		cmdParse.Scan()
		utils.RunCommand(cmdParse.Text())
	}
}
