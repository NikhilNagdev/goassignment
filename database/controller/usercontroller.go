package controller

import (
	"fmt"
	"goassignment/database"
	"goassignment/database/model"
	log "github.com/sirupsen/logrus"
)

func CreateUser(user *model.User)  {
	err := database.GetDbConnection().Create(user).Error
	if err != nil{
		log.SetFormatter(&log.JSONFormatter{})
    	log.Error("An error occured!")
	}
}

func GetUsers(users *[]model.User)  {
	err := database.GetDbConnection().Find(&users).Error
	if err != nil {
		fmt.Println(err)
	}
}
