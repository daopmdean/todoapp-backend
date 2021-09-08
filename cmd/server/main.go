package main

import (
	"todoapp/internal/adapter/http/rest"
	"todoapp/internal/adapter/memdb"
	"todoapp/internal/common/config"
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"
)

func main() {
	userRepo := memdb.NewUserRepo()
	todoRepo := memdb.NewTodoRepo()

	initUserData(userRepo)
	initTodoData(todoRepo)

	s := rest.NewServer(userRepo, todoRepo)
	p := config.GetPort()
	s.Run(p)
}

func initUserData(userRepo repo.UserRepo) {
	dao := entity.User{
		Username:  "daopham",
		FirstName: "Dao",
		LastName:  "Pham",
	}
	dao.SetPassword("12345678")
	hung := entity.User{
		Username:  "hungpham",
		FirstName: "Hung",
		LastName:  "Pham",
	}
	hung.SetPassword("87654321")
	userRepo.SaveUser(&dao)
	userRepo.SaveUser(&hung)
}

func initTodoData(todoRepo repo.TodoRepo) {
	daoTodo := entity.Todo{
		Username: "daopham",
		Content:  "Dao walking",
		IsDone:   false,
	}
	daoTodo2 := entity.Todo{
		Username: "daopham",
		Content:  "Dao learning",
		IsDone:   false,
	}
	hungTodo := entity.Todo{
		Username: "hungpham",
		Content:  "Hung working",
		IsDone:   false,
	}

	todoRepo.CreateTodo(&daoTodo)
	todoRepo.CreateTodo(&daoTodo2)
	todoRepo.CreateTodo(&hungTodo)
}
