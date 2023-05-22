package repl

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ddomd/clidex/internal/pokeapi"
)

//CLI commands contain a name a description and a callback to execute
type CliCommand struct {
	Name        string
	Description string
	Command     func(*pokeapi.Client, ...string) error
}

//Transforms all input to lower case and the splits it into multiple fields to pass as command options
func normalizeInput(s string) []string {
	lower := strings.ToLower(s)
	normalized := strings.Fields(lower)
	return normalized
}

//Creates a map containing all available commands
func getCmds() map[string] CliCommand{
	return map[string] CliCommand {
		"dex": {
			Name: "dex",
			Description: "Displays pokedex information of a pokemon, takes a name or number as an argument 'dex [name/number]'",
			Command: dex,
		},
		"help": {
			Name: "help",
			Description: "Displays a description of all available commands",
			Command: help,
		},
		"quit": {
			Name: "quit",
			Description: "Exits the program",
			Command: quit,
		},
	}
}

func Repl(client *pokeapi.Client) {
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

		input := normalizeInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		args := []string{}

		if len(input) > 1 {
			args = input[1:]
		}

		cmd, ok := cmds[input[0]]
		if ok {
			err := cmd.Command(client, args...)
			if err != nil {
				fmt.Println(err.Error())
			}
			continue
		} else {
			fmt.Printf("\n'%s' command does not exist, for a list of available commands use 'help'\n\n", scanner.Text())
			continue
		}
	}
}
