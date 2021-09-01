package config

import (
	"os"
	"strconv"
)

func GetPort() int {
	envPort := os.Getenv("PORT")
	if envPort == "" {
		return 5000
	}

	port, err := strconv.Atoi(envPort)
	if err != nil {
		panic("PORT must be empty (default 5000) or a number")
	}

	return port
}
