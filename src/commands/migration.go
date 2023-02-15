package commands

import (
	"github.com/go-niom/niom/pkg/logger"
	"github.com/go-niom/niom/pkg/migrate"
)

func create(args []string) {
	if len(args) == 0 {
		logger.Error("Please specify `file name` or -s (to generate sample)", "")
		return
	}
	switch args[0] {
	case "-s", "--sample":
		migrate.CreateSample(args[1:])
	case "up":
	case "down":
	case "refresh":
	default:
		migrate.Create(args[0], "-- Your Script", "-- Your Script")
	}

}

func up(args []string) {
	migrate.Up()
}

func down(args []string) {
	if len(args) > 0 {
		migrate.Down(args[0])
	} else {
		migrate.Down("")
	}

}

func migrations(args []string) {
	switch args[0] {
	case "cr", "create":
		create(args[1:])
	case "up":
		up(args[1:])
	case "down":
		down(args[1:])
	case "refresh":
	default:
		logger.Warn("Invalid Commands")
	}
}
