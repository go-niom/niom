package migrate

import (
	"database/sql"
	"fmt"

	"github.com/go-niom/niom/pkg/config"
	"github.com/go-niom/niom/pkg/logger"
	_ "github.com/lib/pq"
)

// DB holds the database
type DB struct{ *sql.DB }

// database instance
var defaultDB = &DB{}

// connect sets the db client of database using configuration
func (db *DB) connect(dbURI string) (err error) {
	cfg := config.DBCfg()

	if dbURI == "" {
		dbURI = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Name,
			cfg.SslMode,
		)
	}

	db.DB, err = sql.Open("postgres", dbURI)
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
func GetDB(dbUrl string) *DB {
	err := defaultDB.connect(dbUrl)
	if err != nil {
		logger.Error("Execution Failed", err.Error())
		return nil
	}
	return defaultDB
}

// Execute sql queries
func (db *DB) Execute(sql string) error {
	_, err := db.Exec(sql)
	if err != nil {
		logger.Error("Execution Failed", err.Error())
		return err
	}
	return nil
}

// Delete row with Given file dedicated to the migration task
func (db *DB) DeleteRow(fileName string) error {
	_, err := db.Exec(`DELETE FROM migration_scheme WHERE file_name = $1`, fileName)
	if err != nil {
		println(err)
		return err
	}

	return nil
}
