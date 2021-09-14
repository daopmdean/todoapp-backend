package main

import (
	"todoapp/internal/adapter/http/rest"
	"todoapp/internal/adapter/mssqldb"
	"todoapp/internal/common/config"
	"todoapp/internal/entity"

	"gorm.io/gorm"
)

func main() {
	db, err := mssqldb.Connect()
	if err != nil {
		panic("can't connect to database")
	}
	migrateDB(db)

	userRepo := mssqldb.NewUserRepo(db)
	todoRepo := mssqldb.NewTodoRepo(db)

	s := rest.NewServer(userRepo, todoRepo)
	p := config.GetPort()
	s.Run(p)
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&entity.User{}, &entity.Todo{})
}
