package mongodb

import "todoapp/internal/usecase/repo"

func CreateRepos() (repo.UserRepo, repo.TodoRepo) {
	database := Connect()

	userCollection := database.Collection("users")
	todoCollection := database.Collection("todos")

	userRepo := NewUserRepo(userCollection)
	todoRepo := NewTodoRepo(todoCollection)

	return userRepo, todoRepo
}
