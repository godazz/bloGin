package db

import (
	"fmt"
	"log"

	"github.com/godazz/bloGin/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	dbConfiguration := config.Get()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfiguration.DB.Username, dbConfiguration.DB.Password, dbConfiguration.DB.Host, dbConfiguration.DB.Port, dbConfiguration.DB.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("The database connection failed")
		return
	}
	DB = db
}
