package config

import (
	"strconv"
	"time"
)

// DB holds the DB configuration
type DB struct {
	Host            string
	Port            int
	SslMode         string
	Name            string
	User            string
	Password        string
	Debug           bool
	MaxOpenConn     int
	MaxIdleConn     int
	MaxConnLifetime time.Duration
}

var db = &DB{}

// DBCfg returns the default DB configuration
func DBCfg() *DB {
	return db
}

// LoadDBCfg loads DB configuration
// This load db config from env file
// If not it set the default value
func LoadDBCfg() {
	db.Host = getEnv("DB_HOST", "127.0.0.1")
	db.Port, _ = strconv.Atoi(getEnv("DB_PORT", "5432"))
	db.User = getEnv("DB_USER", "dev")
	db.Password = getEnv("DB_PASSWORD", "passwd")
	db.Name = getEnv("DB_NAME", "niom-local")
	db.SslMode = getEnv("DB_SSL_MODE", "false")
	db.Debug, _ = strconv.ParseBool(getEnv("DB_DEBUG", "false"))
	db.MaxOpenConn, _ = strconv.Atoi(getEnv("DB_MAX_OPEN_CONNECTIONS", "3"))
	db.MaxIdleConn, _ = strconv.Atoi(getEnv("DB_MAX_IDLE_CONNECTIONS", "1"))
	lifeTime, _ := strconv.Atoi(getEnv("DB_MAX_LIFETIME_CONNECTIONS", "10"))
	db.MaxConnLifetime = time.Duration(lifeTime) * time.Second
}
