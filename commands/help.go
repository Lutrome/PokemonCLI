package commands

import (
	"fmt"
)

func Help() error {
	fmt.Println("PokeCLI Help:\nAvailable Commands:")
	fmt.Println("help - Shows this help menu.")
	for _, command := range ValidCommands {
		fmt.Printf("%s - %s\n", command.cmdName, command.description)
	}
	return nil
}
