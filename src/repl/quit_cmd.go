package repl

import (
	"os"
	"github.com/ddomd/clidex/internal/pokeapi"
)

func quit(client *pokeapi.Client, args ...string) error {
	os.Exit(0)
	return nil
}