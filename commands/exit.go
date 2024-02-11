package commands

import (
	"fmt"
	"os"
)

func Exit() error {
	fmt.Println("Bye!")
	os.Exit(1)
	return nil
}
