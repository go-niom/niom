package migrate

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/go-niom/niom/pkg/logger"
	"github.com/go-niom/niom/pkg/utils"
	"github.com/go-niom/niom/res/db/blog"
)

// MigrationScheme holds migration rows details
type MigrationScheme struct {
	ID        int       `db:"id"`
	FileName  string    `db:"file_name"`
	Batch     string    `db:"batch"`
	CreatedAt time.Time `db:"created_at"`
}

// getMigrationPath checks form the migrations path(-p=) in the args
// if not found then return the default path
func getMigrationPath(filePath string, args []string) string {
	argsPath := utils.ReadArgs("-p=", args)
	if argsPath != "" {
		filePath = argsPath
		if !strings.HasSuffix(filePath, "/") {
			filePath = filePath + "/"
		}
	}
	return filePath
}

// this is called by CreateSample
// Generates seeded and unseeded migration examples
func pgSqlBlog(isSeed bool, args []string) {

	if !isSeed {
		Create(args, "users", blog.PGCreateUsersTable, blog.PGDropUsersTable)
		Create(args, "categories", blog.PGCreateCategoriesTable, blog.PGDropCategoriesTable)
		Create(args, "posts", blog.PGCreatePostTable, blog.PGDropPostTable)
		Create(args, "comments", blog.PGCreateCommentTable, blog.PGDropCommentTable)
		time.Sleep(1 * time.Second)
		Create(args, "table_alter", blog.PGBlogAlter, blog.PGBlogOppositeAlter)
	} else {
		Create(args, "users", blog.PGCreateUsersTable+"\n"+blog.PGBlogUserInsert, blog.PGDropUsersTable)
		Create(args, "categories", blog.PGCreateCategoriesTable+"\n"+blog.PGBlogCategoriesInsert, blog.PGDropCategoriesTable)
		Create(args, "posts", blog.PGCreatePostTable+"\n"+blog.PGBlogPostInsert, blog.PGDropPostTable)
		Create(args, "comments", blog.PGCreateCommentTable+"\n"+blog.PGBlogCommentInsert, blog.PGDropCommentTable)
		time.Sleep(1 * time.Second)
		Create(args, "table_alter", blog.PGBlogAlter, blog.PGBlogOppositeAlter)
	}
}

// readFileAndMigrate function is used to read sql file
// IT db execute to sql content of the file.
func readFileAndMigrate(db *DB, filePath string) error {
	// read the SQL file contents
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	// execute the SQL commands to create the table
	return db.Execute(string(content))
}

// createMigrationTable creates base migration table to holds migration related data
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

// readFileAndMigrate returns files in ascending order from the migration directory
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
