package config

const RedisConfig = `package config

// DB holds the DB configuration
type RedisDB struct {
	Host     string
	Port     string
	User     string
	Password string
}

var rDB = &RedisDB{}

// DBCfg returns the default DB configuration
func RedisDBCfg() *RedisDB {
	return rDB
}

// LoadDBCfg loads DB configuration
func LoadRedisCfg() {
	rDB.Host = getEnv("REDIS_HOST", "localhost")
	rDB.Port = getEnv("REDIS_PORT", "6379")
	rDB.User = getEnv("REDIS_USER", "")
	rDB.Password = getEnv("REDIS_PASSWORD", "1234")
}`
