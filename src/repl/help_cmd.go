package repl

import "fmt"

func help() error{

	fmt.Println("\nHere is a list of available commands: \n")

	for _, cmd := range getCmds() {
		fmt.Printf("\t%s: %s\n", cmd.name, cmd.description)
	}
		
	return nil
}