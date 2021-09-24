package mssqldb

import (
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"
)

func CreateRepos() (repo.UserRepo, repo.TodoRepo) {
	db, err := Connect()
	if err != nil {
		panic("can't connect to database")
	}
	db.AutoMigrate(&entity.User{}, &entity.Todo{})

	userRepo := NewUserRepo(db)
	todoRepo := NewTodoRepo(db)
	return userRepo, todoRepo
}
