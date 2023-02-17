/*
This file dedicated to intercept command related to the migration
*/
package commands

import (
	"github.com/go-niom/niom/pkg/config"
	"github.com/go-niom/niom/pkg/logger"
	"github.com/go-niom/niom/pkg/migrate"
)

// create functions creates the migrations file
func create(args []string) {
	if len(args) == 0 {
		logger.Error("Please specify `file name` or -s (to generate sample)", "")
		return
	}
	switch args[0] {
	case "-s", "--sample":
		migrate.CreateSample(args[1:])
	default:
		migrate.Create(args[1:], args[0], "-- Your Script", "-- Your Script")
	}

}

// up runs the migrations
func up(args []string) {
	migrate.Up(args)
}

// down rollback the latest migration
func down(args []string) {
	if len(args) > 0 {
		migrate.Down(args[0], args)
	} else {
		migrate.Down("", args)
	}

}

// help shows available commands and arguments of the migrations command
func help() {
	println(`
#To create posts migration files 
$ niom migration cr posts

#This will create posts at the given path -p=
$ niom migration cr posts -p="database/test" 

#This will run migration
$ niom migration up 

#This will run migration from the given path -p=
$ niom migration up -p="database/test" 

#This will show the migration status
$ niom migration status 

#This will rollback migration
$ niom migration down

#This will rollback migration from the given path -p= 
$ niom migration down -p="database/test"

#This will rollback migration all migrations
$ niom migration down -a #-a => all

#This will rollback migration all migrations from the given path -p= 
$ niom migration down -a -p="database/test"
`)
}

// dbInit initialize user database read credentials from the specified env file in the niom-cli.json
func dbInit() string {
	cfg := config.GetNiomCliConfig()
	if cfg != nil {
		if cfg.ConfigFile == "" {
			println("Please specify valid config file in niom-cli.json")
			return ""
		}
		config.LoadAllConfigs(cfg.ConfigFile)
		return ""
	}
	return "Error"

}

// migrations() validate the user input
// and direct to do specified cmd to execute
func migrations(args []string) {
	if len(args) == 0 {
		logger.Warn("Invalid Commands")
		return
	}
	if s := dbInit(); s != "" {
		return
	}

	switch args[0] {
	case "-h", "--help":
		help()
	case "cr", "create":
		create(args[1:])
	case "up":
		up(args[1:])
	case "down":
		down(args[1:])
	case "status":
		migrate.Status()
	default:
		logger.Warn("Invalid Commands please try -h or --help")
	}
}
