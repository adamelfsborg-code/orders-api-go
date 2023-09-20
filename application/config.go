package application

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort   uint16
	RedisAddress string
}

func LoadConfig() Config {
	cfg := Config{
		RedisAddress: "localhost:6379",
		ServerPort:   3000,
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
