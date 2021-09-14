package mssqldb

import (
	"fmt"
	"todoapp/internal/common/config"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		config.MSSQL().Username,
		config.MSSQL().Password,
		config.MSSQL().Host,
		config.MSSQL().Port,
		config.MSSQL().DatabaseName,
	)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	return db, err
}
