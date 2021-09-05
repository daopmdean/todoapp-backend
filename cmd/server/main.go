package main

import (
	"todoapp/internal/adapter/http/rest"
	"todoapp/internal/adapter/memdb"
	"todoapp/internal/common/config"
)

func main() {
	userRepo := memdb.NewUserRepo()

	s := rest.NewServer(userRepo)
	p := config.GetPort()
	s.Run(p)
}
