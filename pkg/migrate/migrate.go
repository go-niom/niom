package migrate

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-niom/niom/pkg/logger"
	"github.com/go-niom/niom/pkg/utils"
	"github.com/go-niom/niom/res/db/blog"
	"github.com/gookit/color"
	_ "github.com/lib/pq"
)

type MigrationScheme struct {
	ID        int       `db:"id"`
	FileName  string    `db:"file_name"`
	Batch     string    `db:"batch"`
	CreatedAt time.Time `db:"created_at"`
}

func Up() {
	logger.Info("Migration started")
	pattern := "up.sql"
	dir := "db/migrations"
	fileList := getFilesAscendingOrder(dir, pattern)
	if fileList == nil {
		logger.Warn("Migration files are not available.")
	}

	db := GetDB()
	if db == nil {
		return
	}
	var err error
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

	migrator := func(fileName, filePath string) error {

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
		if (current > last) || !hasLastFile {
			fileToBeMigrated := fmt.Sprintf("%s/%s", dir, file)
			err := migrator(file, fileToBeMigrated)
			if err != nil {
				return
			}
			hasMigration = true
		}
	}
	if hasMigration {
		logger.Warn("Migration completed")
	} else {
		logger.Warn("Nothing new to Migrate")
	}

}

func Down(arg string) {
	logger.Info("Rollback...")
	dir := "db/migrations"
	isMulti := arg == "-a"
	db := GetDB()
	if db == nil {
		return
	}
	var err error
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

	// res.Scan(&selectedMigration.FileName, &selectedMigration.Batch)

	migrator := func(fileName, filePath string) error {

		err = readFileAndMigrate(db, filePath)
		if err != nil {
			logger.Error("Failed :", err.Error())
			return err
		}

		err = db.DeleteRow(fileName)
		if err != nil {
			panic(err)
		}
		logger.Info("Rollback: " + filePath)
		return nil
	}
	var fileDirs []map[string]string
	for _, ms := range selectedMigration {

		fileName := strings.ReplaceAll(ms.FileName, ".up.", ".down.")
		fileToBeMigrated := fmt.Sprintf("%s/%s", dir, fileName)
		color.Warnf("File: %s\n", fileName)

		fileDir := map[string]string{
			"fileName": ms.FileName,
			"filePath": fileToBeMigrated,
		}
		fileDirs = append(fileDirs, fileDir)
	}

	color.Warnp("\nDo you want to proceed with the Rollback (Y/N)? ")
	value := utils.UserPrompt("")

	if string([]byte(strings.ToUpper(strings.Trim(value, " ")))) == "Y" {
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

func Create(fileName, up, dn string) {
	timestamp := time.Now().Format("20060102150405") // format timestamp as yyyymmddHHMMSS
	upFile := fmt.Sprintf("%s_%s.%s", timestamp, fileName, "up.sql")
	downFile := fmt.Sprintf("%s_%s.%s", timestamp, fileName, "down.sql")
	utils.CreateFileWithData("db/migrations/"+upFile, "BEGIN;\n"+up+"\nCOMMIT;")
	utils.CreateFileWithData("db/migrations/"+downFile, "BEGIN;\n"+dn+"\nCOMMIT;")
}

func CreateSample(args []string) {
	isSeed := false
	if len(args) > 0 && args[0] == "seed" {
		isSeed = true
	}
	pgSqlBlog(isSeed)
}

func pgSqlBlog(isSeed bool) {

	if !isSeed {
		Create("users", blog.PGCreateUsersTable, blog.PGDropUsersTable)
		Create("categories", blog.PGCreateCategoriesTable, blog.PGDropCategoriesTable)
		Create("posts", blog.PGCreatePostTable, blog.PGDropPostTable)
		Create("comments", blog.PGCreateCommentTable, blog.PGDropCommentTable)
		time.Sleep(1 * time.Second)
		Create("table_alter", blog.PGBlogAlter, blog.PGBlogOppositeAlter)
	} else {
		Create("users", blog.PGCreateUsersTable+"\n"+blog.PGBlogUserInsert, blog.PGDropUsersTable)
		Create("categories", blog.PGCreateCategoriesTable+"\n"+blog.PGBlogCategoriesInsert, blog.PGDropCategoriesTable)
		Create("posts", blog.PGCreatePostTable+"\n"+blog.PGBlogPostInsert, blog.PGDropPostTable)
		Create("comments", blog.PGCreateCommentTable+"\n"+blog.PGBlogCommentInsert, blog.PGDropCommentTable)
		time.Sleep(1 * time.Second)
		Create("table_alter", blog.PGBlogAlter, blog.PGBlogOppositeAlter)
	}
}

func readFileAndMigrate(db *DB, filePath string) error {
	// read the SQL file contents
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// execute the SQL commands to create the table
	return db.Execute(string(content))

}

func createMigrationTable(db *DB) error {
	miTable := `
	CREATE TABLE migration_scheme (
		id SERIAL PRIMARY KEY,
		file_name VARCHAR(255) NOT NULL,
		batch VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	  );`
	_, err := db.Exec(miTable)
	if err != nil {
		logger.Error("Execution Failed", err.Error())
		return err
	}
	return nil
}

func getFilesAscendingOrder(dir string, pattern string) []string {

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	// Filter files by name ending with pattern

	var sortedFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), pattern) {
			sortedFiles = append(sortedFiles, file.Name())
		}
	}

	// Sort files by name in ascending order
	sort.Slice(sortedFiles, func(i, j int) bool {
		return sortedFiles[i] < sortedFiles[j]
	})

	// Return the sorted files with the pattern
	return sortedFiles
}
