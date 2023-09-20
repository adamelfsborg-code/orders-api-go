package application

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort       uint16
	RedisAddress     string
	PostgresAddr     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	RepoAdapter      string
}

func LoadConfig() Config {
	cfg := Config{
		ServerPort:       3000,
		RedisAddress:     "localhost:6379",
		PostgresAddr:     "localhost:5432",
		PostgresUser:     "admin",
		PostgresPassword: "123",
		PostgresDatabase: "sudokus",
		RepoAdapter:      "PSQL",
	}

	redisAddr, exists := os.LookupEnv("REDIS_ADDR")
	if exists {
		cfg.RedisAddress = redisAddr
	}

	serverPort, exists := os.LookupEnv("SERVER_PORT")
	if exists {
		const base = 10
		const bitSize = 16
		port, err := strconv.ParseUint(serverPort, base, bitSize)
		if err == nil {
			cfg.ServerPort = uint16(port)
		}
	}
	return cfg
}
