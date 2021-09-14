package main

import (
	"os"
	"todoapp/internal/adapter/http/rest"
	"todoapp/internal/adapter/memdb"
	"todoapp/internal/adapter/mssqldb"
	"todoapp/internal/common/config"
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"

	"gorm.io/gorm"
)

func main() {
	var userRepo repo.UserRepo
	var todoRepo repo.TodoRepo

	switch os.Getenv("DB_MODE") {
	case "mssql":
		db, err := mssqldb.Connect()
		if err != nil {
			panic("can't connect to database")
		}
		migrateDB(db)

		userRepo = mssqldb.NewUserRepo(db)
		todoRepo = mssqldb.NewTodoRepo(db)
	default:
		userRepo = memdb.NewUserRepo()
		todoRepo = memdb.NewTodoRepo()
	}

	s := rest.NewServer(userRepo, todoRepo)
	p := config.GetPort()
	s.Run(p)
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&entity.User{}, &entity.Todo{})
}
