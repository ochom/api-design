package utils

import "os"

// GetEnv returns the value of the environment variable named by the key.
// If the environment variable is not set it returns the defaultValue.
func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

// MustGetEnv returns the value of the environment variable named by the key.
// It panics if the environment variable is not set.
func MustGetEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	panic("environment variable " + key + " is not set")
}
