package config

import (
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	currentConfig := GetConfig()
	connectstring := currentConfig.Database.DBUser + ":" + currentConfig.Database.DBPassword + "@tcp(" + currentConfig.Database.DBHost + ":" + strconv.Itoa(currentConfig.Database.DBPort) + ")/" + currentConfig.Database.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := connectstring
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = db
}

func GetDBConn() *gorm.DB {

	return DB
}
