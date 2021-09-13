package mssqldb

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	username := "sa"
	password := "MyStrongPassword123"
	databaseHost := "localhost:1433"
	databaseName := "todoapp"
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", username, password, databaseHost, databaseName)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	return db, err
}
