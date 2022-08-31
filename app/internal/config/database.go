package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectPostgreSQL() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/latihan_1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
