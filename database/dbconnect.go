package database

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = ""
const DB_NAME = "test"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var db *gorm.DB

func GetDbConnection() *gorm.DB {
	db = connectDB()
	return db
}

func connectDB() (*gorm.DB) {
	var err error
	dsn := DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.SetFormatter(&log.JSONFormatter{})
    	log.Error(err)
		return nil
	}
	//fmt.Println("Connected to DB")
 	return db
}