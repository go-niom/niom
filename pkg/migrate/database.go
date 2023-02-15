package migrate

import (
	"database/sql"
	"fmt"

	"github.com/go-niom/niom/pkg/logger"
	_ "github.com/lib/pq"
)

// DB holds the database
type DB struct{ *sql.DB }

// database instance
var defaultDB = &DB{}

// connect sets the db client of database using configuration
func (db *DB) connect() (err error) {

	db.DB, err = sql.Open("postgres", "postgresql://dev:passwd@127.0.0.1:5432/niom-local?sslmode=disable")
	if err != nil {
		logger.Error("DB Connection Error", err.Error())
	}
	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return fmt.Errorf("can't sent ping to database, %w", err)
	} else {
		logger.Info("Database Connected")
	}

	return nil
}

// GetDB returns db instance
func GetDB() *DB {
	err := defaultDB.connect()
	if err != nil {
		return nil
	}
	return defaultDB
}

func (db *DB) Execute(sql string) error {
	_, err := db.Exec(sql)
	if err != nil {
		logger.Error("Execution Failed", err.Error())
		return err
	}

	return nil
}

func (db *DB) DeleteRow(fileName string) error {
	_, err := db.Exec(`DELETE FROM migration_scheme WHERE file_name = $1`, fileName)
	if err != nil {
		println(err)
		return err
	}

	return nil
}
