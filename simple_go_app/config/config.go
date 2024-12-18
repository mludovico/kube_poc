package config

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	configs map[string]string
)

const (
	APIHost     = "API_HOST"
	APIPort     = "API_PORT"
	DBUser      = "DB_USER"
	DBPass      = "DB_PASSWORD"
	DBHost      = "DB_HOST"
	DBPort      = "DB_PORT"
	DBName      = "DB_NAME"
	AdminDBUser = "ADMIN_DB_USER"
	AdminDBPass = "ADMIN_DB_PASSWORD"
	AdminDBName = "ADMIN_DB_NAME"
	SeedDB      = "SEED_DB"
)

var neededVars = []string{
	APIHost, APIPort, DBUser, DBPass, DBHost, DBPort, DBName, AdminDBUser, AdminDBPass, AdminDBName, SeedDB,
}

func LoadConfigs() {
	fmt.Println("Initializing configs...")
	configs = make(map[string]string)
	for _, key := range neededVars {
		if os.Getenv(key) == "" {
			fmt.Println("Could not find", key, "in environment variables")
			loadConfigsFromDotEnvFile()
			return
		}
	}
	loadConfigsFromEnv()
}

func loadConfigsFromEnv() {
	fmt.Println("Loading configs from environment variables")
	for _, key := range neededVars {
		configs[key] = os.Getenv(key)
	}
}

func PrintConfigs() {
	fmt.Println("Configs:")
	for key, value := range configs {
		fmt.Printf("%s = %s\n", key, value)
	}
}

func loadConfigsFromDotEnvFile() {
	fmt.Println("Loading configs from .env file")
	content, err := os.ReadFile(".env")
	if err != nil {
		log.Fatalf("could not read .env file: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if line == "" || strings.HasPrefix(line, "#") ||
			strings.HasPrefix(line, ";") ||
			!strings.Contains(line, "=") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		configs[parts[0]] = parts[1]
	}
}

func GetVar(key string) string {
	return configs[key]
}
