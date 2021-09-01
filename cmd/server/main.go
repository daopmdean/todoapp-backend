package main

import (
	"todoapp/internal/adapter/http/rest"
	"todoapp/internal/common/config"
)

func main() {
	server := rest.NewServer()
	port := config.GetPort()
	server.Run(port)
}
