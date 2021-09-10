package main

import (
	"todoapp/internal/adapter/http/rest"
	"todoapp/internal/adapter/memdb"
	"todoapp/internal/common/config"
)

func main() {
	userRepo := memdb.NewUserRepo()
	todoRepo := memdb.NewTodoRepo()

	s := rest.NewServer(userRepo, todoRepo)
	p := config.GetPort()
	s.Run(p)
}
