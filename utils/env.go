package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// StringMap define a custom type for the map
type StringMap map[string]string

var AppEnv StringMap = make(StringMap)

// GetOrDefault method on StringMap type, acting like an extension method in C#
func (m StringMap) GetOrDefault(key string, defaultValue string) string {
	// Check if the key exists in the map
	if value, exists := m[key]; exists {
		return value
	}
	// Return the default value if the key doesn't exist
	return defaultValue
}

// GetEnv retrieves the value of the environment variable named by the key.
func GetEnv(key, defaultValue string) string {
	//value, exists := os.LookupEnv(key)
	value, exists := AppEnv[key]
	if !exists {
		log.Printf("Environment variable %s not found, using default: %s", key, defaultValue)
		return defaultValue
	}
	return value
}

// InitEnv load environment variables
func InitEnv() {

	var filenames []string

	if FileExists(".env") {
		filenames = append(filenames, ".env")
	}

	if FileExists(".env.dev") {
		filenames = append(filenames, ".env.dev")
	}

	env, err := godotenv.Read(filenames...)
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	for k, v := range env {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
		AppEnv[k] = v
	}
	fmt.Println()

}
