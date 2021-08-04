package controller

import (
	"github.com/sirupsen/logrus"
	"goassignment/database"
	"goassignment/database/model"
)

func CreateUser(user *model.User, log *logrus.Logger) {
	err := database.GetDbConnection(log).Create(user).Error
	if err != nil {
		log.Error(err)
	}
}

func GetUsers(users *[]model.User, log *logrus.Logger) {
	err := database.GetDbConnection(log).Find(&users).Error
	if err != nil {
		log.Error(err)
	}
}
