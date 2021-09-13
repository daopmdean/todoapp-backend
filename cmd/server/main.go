package main

import (
	"todoapp/internal/adapter/http/rest"
	"todoapp/internal/adapter/mssqldb"
	"todoapp/internal/common/config"
	"todoapp/internal/entity"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	// userRepo := memdb.NewUserRepo()
	// todoRepo := memdb.NewTodoRepo()
	dsn := "sqlserver://sa:MyStrongPassword123@localhost:1433?database=todoapp"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}
	setup(db)

	userRepo := mssqldb.NewUserRepo(db)
	todoRepo := mssqldb.NewTodoRepo(db)

	s := rest.NewServer(userRepo, todoRepo)
	p := config.GetPort()
	s.Run(p)
}

func setup(db *gorm.DB) {
	db.AutoMigrate(&entity.User{}, &entity.Todo{})
}
