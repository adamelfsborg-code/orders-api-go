package application

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort       uint16
	RedisAddr        string
	PostgresAddr     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	RepoAdapter      string
}

func LoadConfig() Config {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}

	cfg := Config{}

	serverPort, exists := os.LookupEnv("SERVER_PORT")
	if exists {
		const base = 10
		const bitSize = 16
		port, err := strconv.ParseUint(serverPort, base, bitSize)
		if err == nil {
			cfg.ServerPort = uint16(port)
		}
	}

	redisAddr, exists := os.LookupEnv("REDIS_ADDR")
	if exists {
		cfg.RedisAddr = redisAddr
	}

	postgresAddr, exists := os.LookupEnv("POSTGRES_ADDR")
	if exists {
		cfg.PostgresAddr = postgresAddr
	}

	postgresUser, exists := os.LookupEnv("POSTGRES_USER")
	if exists {
		cfg.PostgresUser = postgresUser
	}

	postgresDatabase, exists := os.LookupEnv("POSTGRES_DATABASE")
	if exists {
		cfg.PostgresDatabase = postgresDatabase
	}

	postgresPassword, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if exists {
		cfg.PostgresPassword = postgresPassword
	}

	repoAdapter, exists := os.LookupEnv("REPO_ADAPTER")
	if exists {
		cfg.RepoAdapter = repoAdapter
	}

	return cfg
}
