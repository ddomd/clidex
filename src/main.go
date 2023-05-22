package main

import (
	"time"

	"github.com/ddomd/clidex/repl"
	"github.com/ddomd/clidex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(time.Second*5, time.Minute*5)
	repl.Repl(&client)
}
