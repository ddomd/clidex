package repl

import (
	"fmt"
	"github.com/ddomd/clidex/internal/pokeapi"
)

func help(client *pokeapi.Client, args ...string) error{

	fmt.Println("\n\tHere is a list of available commands: \n")

	for _, cmd := range getCmds() {
		fmt.Printf("\t\t%s: %s\n", cmd.Name, cmd.Description)
	}
		
	return nil
}