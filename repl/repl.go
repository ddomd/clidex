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

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("-- WELCOME TO THE CLIDEX --")

	for {
		fmt.Print("Pokédex > ")
		scanner.Scan()

		err := scanner.Err()

		if err != nil {
			log.Fatal()
		}

		fmt.Printf("Your input is: %s\n", scanner.Text())
	}
}
