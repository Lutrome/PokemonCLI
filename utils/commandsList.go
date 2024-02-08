package utils

import (
	"fmt"
	"github.com/Lutrome/PokemonCLI/utils/commands"
)

type cliCommand struct {
	cmdName     string
	description string
	callback    func() error
}

func RunCommand(cmd string) {
	validCommands := map[string]cliCommand{
		"help": {
			"help",
			"Shows this help menu.",
			commands.Help,
		},
		"exit": {
			"exit",
			"Exits the Pokedex.",
			commands.Exit,
		},
		"map": {
			"map",
			"",
			commands.Map,
		},
		"mapb": {
			"map",
			"",
			commands.Mapb,
		},
	}

	_, exists := validCommands[cmd]
	if exists {
		err := validCommands[cmd].callback()
		if err != nil {
			fmt.Println("[ERROR]", err)
			return
		}
	} else {
		fmt.Println("Wrong input, or command not implemented yet!")
		return
	}
}
