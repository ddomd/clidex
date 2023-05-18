package repl

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type cliCommand struct {
	name        string
	description string
	command     func() error
}

func getCmds() map[string] cliCommand{
	return map[string] cliCommand {
		"help": {
			name: "help",
			description: "Displays a description of all available commands",
			command: help,
		},
		"quit": {
			name: "quit",
			description: "Exits the program",
			command: quit,
		},
	}
}

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)
	
	print_splash()

	cmds := getCmds()

	for {
		fmt.Print("\n\u25d2 > ")
		scanner.Scan()

		err := scanner.Err()

		if err != nil {
			log.Fatal()
		}

		cmd, ok := cmds[scanner.Text()]
		if ok {
			cmd.command()
		} else {
			fmt.Printf("\n'%s' command does not exist, for a list of available commands use 'help'\n\n", scanner.Text())
		}
	}
}
