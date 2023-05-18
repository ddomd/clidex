package repl

import "os"

func quit() error {
	os.Exit(0)
	return nil
}