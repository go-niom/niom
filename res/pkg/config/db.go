package config

const DBConfig = `package config

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
func LoadDBCfg() {
	db.Host = getEnv("DB_HOST", "localhost")
	db.Port, _ = strconv.Atoi(getEnv("DB_PORT", "5432"))
	db.User = getEnv("DB_USER", "root")
	db.Password = getEnv("DB_PASSWORD", "1234")
	db.Name = getEnv("DB_NAME", "niom")
	db.SslMode = getEnv("DB_SSL_MODE", "false")
	db.Debug, _ = strconv.ParseBool(getEnv("DB_DEBUG", "false"))
	db.MaxOpenConn, _ = strconv.Atoi(getEnv("DB_MAX_OPEN_CONNECTIONS", "3"))
	db.MaxIdleConn, _ = strconv.Atoi(getEnv("DB_MAX_IDLE_CONNECTIONS", "1"))
	lifeTime, _ := strconv.Atoi(getEnv("DB_MAX_LIFETIME_CONNECTIONS", "10"))
	db.MaxConnLifetime = time.Duration(lifeTime) * time.Second
}
`
