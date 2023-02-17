/*
Niom is a Golang web framework for developing efficient and scalable server-side applications.
*/
package main

import (
	"os"

	"github.com/go-niom/niom/src/commands"
	"github.com/go-niom/niom/src/handler"
)

// main magic starts from here
func main() {
	args := os.Args
	if len(args) == 1 {
		handler.Welcome()
		return
	}
	commands.Commands(args)
}
