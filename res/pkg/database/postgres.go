package database

const Postgres = `package database

import (
	"fmt"

	"{{ .ModuleName}}/pkg/config"
	"{{ .ModuleName}}/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type db struct {
	*gorm.DB
}

var (
	DB *db
)

func GetDB() *db {
	cfg := config.DBCfg()
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Name,
	)
	logger.Info("Connecting to database")
	logger.Info(dsn)

	// open database connection
	dbo, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	// check error
	if err != nil {
		logger.Error("failed to connect to the database: ", err)
		return nil
	}

	sqlDB, err := dbo.DB()
	if err != nil {
		logger.Error("failed to get underlying sql.DB: ", err)
		return nil
	}

	dbo.Debug()

	// Ping the database
	if err := sqlDB.Ping(); err != nil {
		// Close the database connection
		if err := sqlDB.Close(); err != nil {
			logger.Error("Error while closing the database connection", err)
		}
		logger.Error("failed to ping the database:", err)
		return nil
	}

	logger.Info("Database Connected")
	DB = &db{dbo}
	return DB
}
`
