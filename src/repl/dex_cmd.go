package repl

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/ddomd/clidex/internal/pokeapi"
)

func longestString(list []string) int {
	max := 20
	

	for i, _ := range list {
		if utf8.RuneCountInString(list[i]) > max {
			max = utf8.RuneCountInString(list[i])
		}
	}

	return max
}


func normalizeFlavorText(s string) []string {
	replacer1 := strings.NewReplacer("\f", "\n")
	replacer2 := strings.NewReplacer("\u00ad\n", "")
	replacer3 := strings.NewReplacer("\u00ad", "")

	s = replacer1.Replace(s)
	s = replacer2.Replace(s)
	s = replacer3.Replace(s)

	formattedText := strings.Split(s, "\n")

	return formattedText
}

func display(p pokeapi.Pokedex) {
	name := strings.ToUpper(p.Name)
	height := fmt.Sprintf("%.2f", float64(p.Height) / 10)
	weight := fmt.Sprintf("%.2f", float64(p.Weight) / 10)
	num := fmt.Sprintf("%d", p.Number)
	flavor := normalizeFlavorText(p.Flavor)

	separator := fmt.Sprintf("        +%s+", strings.Repeat("-", longestString(flavor)+2))
	pad := utf8.RuneCountInString(separator)

	fmt.Println("")
	fmt.Println(separator)
	fmt.Printf("\t|%s|\n", strings.Repeat(" ", pad-10))
	fmt.Printf("\t| %s%s|\n", strings.ToUpper(name), strings.Repeat(" ", (pad - utf8.RuneCountInString(name)) - 11))
	fmt.Printf("\t|%s|\n", strings.Repeat(" ", pad-10))
	fmt.Printf("\t| HT: %s m%s |\n", height, strings.Repeat(" ", (pad - utf8.RuneCountInString(height))-18))
	fmt.Printf("\t| WT: %s kg%s |\n", weight, strings.Repeat(" ", (pad - utf8.RuneCountInString(weight))-19))
	fmt.Printf("\t| No. %s%s |\n", num, strings.Repeat(" ", (pad - utf8.RuneCountInString(num))-16))
	fmt.Printf("\t|%s|\n", strings.Repeat(" ", pad-10))
	fmt.Println(separator)
	fmt.Printf("\t|%s|\n", strings.Repeat(" ", pad-10))

	for i, _ := range flavor {
		fmt.Printf("        | %s%s|\n", flavor[i], strings.Repeat(" ", (pad - utf8.RuneCountInString(flavor[i]))-11))
	}

	fmt.Printf("\t|%s|\n", strings.Repeat(" ", pad-10))
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