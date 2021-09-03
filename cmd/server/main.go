package main

import (
	"todoapp/internal/adapter/http/rest"
	"todoapp/internal/adapter/memdb"
	"todoapp/internal/common/config"
)

func main() {
	userRepo := memdb.NewUserRepo()

	server := rest.NewServer(userRepo)
	port := config.GetPort()
	server.Run(port)
}
