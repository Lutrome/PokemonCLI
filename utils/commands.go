package utils

import (
	"fmt"
	"os"
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
			commandHelp,
		},
		"exit": {
			"exit",
			"Exits the Pokedex.",
			commandExit,
		},
	}

	_, exists := validCommands[cmd]
	if exists {
		err := validCommands[cmd].callback()
		if err != nil {
			fmt.Println("[ERROR!]", err)
			return
		}
	} else {
		fmt.Println("Wrong input, or command not implemented yet!")
		return
	}
}

func commandHelp() error {
	fmt.Println("Help menu is under construction.")
	return nil
}

func commandExit() error {
	fmt.Println("Bye!")
	os.Exit(1)
	return nil
}
