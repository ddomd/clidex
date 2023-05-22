package repl

import (
	"fmt"
	"strings"
	"errors"
	"github.com/ddomd/clidex/internal/pokeapi"
)

func display(p pokeapi.Pokedex) {
	separator := "\t----------------------------------------"

	fmt.Println("")
	fmt.Println(separator)
	fmt.Println("\t| ")
	fmt.Println("\t| \t\t" + strings.ToUpper(p.Name))
	fmt.Println("\t| ")
	fmt.Printf("\t| \t\tHT: %.2f m\n", float64(p.Height) / 10)
	fmt.Printf("\t| \t\tWT: %.2f kg\n", float64(p.Weight) / 10)
	fmt.Println("\t| ")
	fmt.Println(separator)
	fmt.Println("\t| ")
	fmt.Println("\t| ", p.Flavor)
	fmt.Println("\t| ")
	fmt.Println(separator)
}

func dex(client *pokeapi.Client, args ...string) error{

	if len(args) < 1 {
		return errors.New("You must provide a pokemon name or number")
	}

	if len(args) > 2 {
		return errors.New("Too many arguments command 'dex' only accepts one argument")
	}

	query := args[0]
	
	dexEntry, err := client.GetDex(query)

	if err != nil {
		return err
	}

	display(dexEntry)

	return nil
}