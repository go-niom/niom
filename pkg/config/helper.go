package config

import "os"

// This function tries to read env value from key
// If the key is not available then returns default value
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
