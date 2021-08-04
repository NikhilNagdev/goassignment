package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = ""
const DB_NAME = "test"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var db *gorm.DB

func GetDbConnection(log *logrus.Logger) *gorm.DB {
	db = connectDB(log)
	return db
}

func connectDB(log *logrus.Logger) *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Connecting to Database ", err)
		return nil
	}
	//fmt.Println("Connected to DB")
	return db
}
