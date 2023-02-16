package commands

import (
	"github.com/go-niom/niom/pkg/config"
	"github.com/go-niom/niom/pkg/logger"
	"github.com/go-niom/niom/pkg/migrate"
	"github.com/go-niom/niom/pkg/utils"
)

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

func up(args []string) {
	migrate.Up(args)
}

func down(args []string) {
	if len(args) > 0 {
		migrate.Down(args[0], args)
	} else {
		migrate.Down("", args)
	}

}
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

func dbInit() {
	cfg := utils.GetNiomCliConfig()
	if cfg != nil {
		if cfg.ConfigFile == "" {
			println("Please specify valid config file in niom-cli.json")
			return
		}
		config.LoadAllConfigs(cfg.ConfigFile)
	}

}

func migrations(args []string) {
	if len(args) == 0 {
		logger.Warn("Invalid Commands")
		return
	}
	dbInit()

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
