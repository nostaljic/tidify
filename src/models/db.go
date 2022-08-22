package models

import (
	"fmt"
	"tidify/devlog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection(user string, pass string, host string, port string, dbname string) *gorm.DB {

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	devlog.Debug(url)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		devlog.Panic(err)
	}
	return db

}
