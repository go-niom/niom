package migrate

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-niom/niom/pkg/logger"
	"github.com/go-niom/niom/pkg/utils"
	"github.com/gookit/color"
	_ "github.com/lib/pq"
)

// Up migrates the migrations files
// `niom migration up` may used to invoke this function
func Up(args []string) {

	//indicates migration started
	logger.Info("Migration started")
	pattern := "up.sql"

	//check for migration path in the args or set db/migrations/ as default migration path
	dir := getMigrationPath("db/migrations/", args)

	//fileList holds all files from migration path
	fileList := getFilesAscendingOrder(dir, pattern)
	if fileList == nil {
		logger.Warn("Migration files are not available.")
	}

	// read db url from args
	dbUrl := utils.ReadArgs("-d=", args)

	// initialize db
	db := GetDB(dbUrl)
	if db == nil {
		return
	}
	var err error

	//fetch latest migration from migration_scheme table
	res := db.QueryRow(`SELECT file_name,batch from migration_scheme order by id desc`)
	if res.Err() != nil {
		if strings.Contains(res.Err().Error(), "does not") {
			err := createMigrationTable(db)
			if err != nil {
				logger.Error("", res.Err().Error())
				return
			}
		}
	}
	var selectedMigration MigrationScheme
	res.Scan(&selectedMigration.FileName, &selectedMigration.Batch)

	//this function do migration and store migration information into the migration_scheme table
	migrator := func(fileName, filePath string) error {

		// readfile and run the execution
		err = readFileAndMigrate(db, filePath)
		if err != nil {
			logger.Error("Failed :", err.Error())
			return err
		}
		timestamp := time.Now().Format("20060102150405")

		// INSERT a new row into the migration_scheme table
		migration := MigrationScheme{
			FileName:  fileName,
			Batch:     timestamp,
			CreatedAt: time.Now(),
		}
		insertQuery := "INSERT INTO public.migration_scheme (file_name, batch, created_at) VALUES ($1, $2, $3) RETURNING id"
		err = db.QueryRow(insertQuery, migration.FileName, migration.Batch, migration.CreatedAt).Scan(&migration.ID)
		if err != nil {
			panic(err)
		}
		logger.Info("Migrated: " + fileName)
		return nil
	}

	hasMigration := false
	for _, file := range fileList {
		var last, current int64
		hasLastFile := selectedMigration.FileName != ""
		if hasLastFile {
			last, _ = strconv.ParseInt(strings.Split(selectedMigration.FileName, "_")[0], 10, 64)
			current, _ = strconv.ParseInt(strings.Split(file, "_")[0], 10, 64)
		}
		//compare latest executed file with new file in the migration directory
		//call migrator func. to migrate latest file from the migration directory
		if (current > last) || !hasLastFile {
			fileToBeMigrated := fmt.Sprintf("%s/%s", dir, file)
			err := migrator(file, fileToBeMigrated)
			if err != nil {
				return
			}
			hasMigration = true
		}
	}

	//check migration was there or not and log appropriate message
	if hasMigration {
		logger.Info("Migration completed")
	} else {
		logger.Warn("Nothing new to Migrate")
	}
}

// Down rollback the latest migration
// `niom migration down` may used to invoke this function
func Down(arg string, args []string) {

	//indicates rollback started
	logger.Info("Rollback...")

	//check for migration path in the args or set db/migrations/ as default migration path
	dir := getMigrationPath("db/migrations/", args)

	//check rollback all file or not
	isMulti := arg == "-a"

	// read db url from args
	dbUrl := utils.ReadArgs("-d=", args)

	//initialize db
	db := GetDB(dbUrl)
	if db == nil {
		return
	}
	var err error

	//fetch all migration from migration_scheme table
	rows, err := db.Query(`SELECT file_name,batch from migration_scheme order by id desc`)
	if err != nil {
		if strings.Contains(err.Error(), "does not") {
			err := createMigrationTable(db)
			if err != nil {
				logger.Error("", err.Error())
				return
			}
		}
	}

	//holds the list migrations fetch from the db
	var selectedMigration []MigrationScheme
	for rows.Next() {
		var s MigrationScheme
		err := rows.Scan(&s.FileName, &s.Batch)
		if err != nil {
			log.Fatal(err)
		}
		selectedMigration = append(selectedMigration, s)
		if !isMulti {
			break
		}
	}

	//this function do migration and store migration information into the migration_scheme table
	migrator := func(fileName, filePath string) error {

		err = readFileAndMigrate(db, filePath)
		if err != nil {
			logger.Error("Failed :", err.Error())
			return err
		}

		// delete the rollback row form the database
		err = db.DeleteRow(fileName)
		if err != nil {
			panic(err)
		}

		//log rollback file details
		logger.Info("Rollback: " + filePath)
		return nil
	}
	var fileDirs []map[string]string
	for _, ms := range selectedMigration {

		// Changing the name from .up. to .down.
		fileName := strings.ReplaceAll(ms.FileName, ".up.", ".down.")
		fileToBeMigrated := fmt.Sprintf("%s/%s", dir, fileName)
		color.Warnf("File: %s\n", fileName)

		// Create the list of file to be rollback
		fileDir := map[string]string{
			"fileName": ms.FileName,
			"filePath": fileToBeMigrated,
		}
		fileDirs = append(fileDirs, fileDir)
	}

	color.Warnp("\nDo you want to proceed with the Rollback (Y/N)? ")
	//wait for the user confirmation
	value := utils.UserPrompt("")

	//check for the user response
	if strings.ToUpper(strings.Trim(value, " ")) == "Y" {
		logger.Info("Rollback entered")
		for _, v := range fileDirs {
			migrator(v["fileName"], v["filePath"])
			if err != nil {
				return
			}
		}
		return
	}
	logger.Info("Rollback aborted")
}

// CreateSample create blogs migration sample
// `niom migration -s` invokes this function
// `niom migration -s seed` may be used generate insert statement
func CreateSample(args []string) {
	isSeed := false
	if len(args) > 0 && args[0] == "seed" {
		isSeed = true
	}
	pgSqlBlog(isSeed, args)
}

// Status fetch the migration details from the db
// `niom migration status` invokes this function
func Status(args []string) {
	// read db url from args
	dbUrl := utils.ReadArgs("-d=", args)

	//initialize db
	db := GetDB(dbUrl)
	if db == nil {
		return
	}
	var err error
	rows, err := db.Query(`SELECT * from migration_scheme`)
	if err != nil {
		if strings.Contains(err.Error(), "does not") {
			err := createMigrationTable(db)
			if err != nil {
				logger.Error("", err.Error())
				return
			}
		}
	}

	count := 0
	for rows.Next() {
		var s MigrationScheme
		err := rows.Scan(&s.ID, &s.FileName, &s.Batch, &s.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s	%d 	%s \n", s.CreatedAt.Format(time.RFC3339), s.ID, s.FileName)
		count += 1
	}
	logger.Info(fmt.Sprintf("Total No. of Migration: %d", count))
}

// CreateWithPath creates migration file with give fileName and filePath
// It also accept up and down scripts to add into file while creating it.
func CreateWithPath(fileName, filePath, up, dn string) {
	timestamp := time.Now().Format("20060102150405") // format timestamp as yyyymmddHHMMSS
	upFile := fmt.Sprintf("%s_%s.%s", timestamp, fileName, "up.sql")
	downFile := fmt.Sprintf("%s_%s.%s", timestamp, fileName, "down.sql")
	utils.CreateFileWithData(filePath+upFile, "BEGIN;\n"+up+"\nCOMMIT;")
	utils.CreateFileWithData(filePath+downFile, "BEGIN;\n"+dn+"\nCOMMIT;")
}

// Create is the parent func of CreateWithPath
// Checks migration path in the args if not then set db/migrations/ as default
func Create(args []string, fileName string, up, dn string) {
	filePath := getMigrationPath("db/migrations/", args)
	CreateWithPath(fileName, filePath, up, dn)
}
