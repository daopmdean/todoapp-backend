package main

import (
	"os"
	"todoapp/internal/adapter/http/rest"
	"todoapp/internal/adapter/memdb"
	"todoapp/internal/adapter/mongodb"
	"todoapp/internal/adapter/mssqldb"
	"todoapp/internal/common/config"
	"todoapp/internal/usecase/repo"
)

func main() {
	var userRepo repo.UserRepo
	var todoRepo repo.TodoRepo

	switch os.Getenv("DB_MODE") {
	case "mssql":
		userRepo, todoRepo = mssqldb.CreateRepos()
	case "mongodb":
		userRepo, todoRepo = mongodb.CreateRepos()
	default:
		userRepo, todoRepo = memdb.CreateRepos()
	}

	s := rest.NewServer(userRepo, todoRepo)
	p := config.GetPort()
	s.Run(p)
}
