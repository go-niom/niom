package config

const JWTConfig = `package config

import "strconv"

type JWT struct {
	// JWT Conf
	JWTSecretKey                string
	JWTSecretExpireMinutesCount int
}

var jwt = &JWT{}

// JWT returns the default JWT configuration
func JWTCfg() *JWT {
	return jwt
}

func LoadJWTCfg() {
	jwt.JWTSecretKey = getEnv("JWT_SECRET_KEY", "JWT_SECRET_KEY")
	jwt.JWTSecretExpireMinutesCount, _ = strconv.Atoi(getEnv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT", "15"))
}
`
