package db

import "gorm.io/gorm"

func Connecion() *gorm.DB {
	return DB
}
