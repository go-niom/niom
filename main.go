package main

import (
	"os"

	"github.com/go-niom/niom/src/commands"
	"github.com/go-niom/niom/src/handler"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		handler.Welcome()
		return
	}
	commands.Commands(args)
}
