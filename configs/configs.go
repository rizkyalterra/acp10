package configs

import (
	"acp10/models/news"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Database string
}

func InitConfigDB() ConfigDB {
	var configDB = ConfigDB{
		DB_Username: "root",
		DB_Password: "123ABC4d.",
		DB_Host:     "127.0.0.1",
		DB_Port:     "3306",
		DB_Database: "acp10",
	}
	return configDB
}

func InitDB() {
	configDB := InitConfigDB()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		configDB.DB_Username,
		configDB.DB_Password,
		configDB.DB_Host,
		configDB.DB_Port,
		configDB.DB_Database)

	var error error
	DB, error = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		panic("Database failed connection : " + error.Error())
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&news.News{})
}
