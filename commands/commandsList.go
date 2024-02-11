package commands

import (
	"fmt"
)

type cliCommand struct {
	cmdName     string
	description string
	callback    func() error
}

var ValidCommands = map[string]cliCommand{
	"exit": {
		"exit",
		"Exits the Pokedex.",
		Exit,
	},
	"map": {
		"map",
		"Shows the next 20 locations of the map.",
		Map,
	},
	"mapb": {
		"mapb",
		"Shows the previous 20 locations of the map.",
		Mapb,
	},
}

func RunCommand(cmd string) {
	_, exists := ValidCommands[cmd]
	if exists {
		err := ValidCommands[cmd].callback()
		if err != nil {
			fmt.Println("[ERROR]", err)
			return
		}
	} else {
		fmt.Println("Wrong input, or command not implemented yet!")
		return
	}
}
