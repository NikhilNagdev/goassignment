package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"goassignment/database/controller"
	"goassignment/database/model"
	"goassignment/io"
	"net/http"
	"strconv"
)

var users []model.User
var log = logrus.New()

func CovertCSVToModel(fileName string) []model.User {
	userRecords := io.GetRecordsFromCSVReader(
		io.ReadCSVFile(fileName))
	fmt.Println(userRecords)
	for i := 0; i < len(userRecords)-1; i++ {
		phone, _ := strconv.Atoi(userRecords[i][3])
		isActive, _ := strconv.Atoi(userRecords[i][4])
		user := model.User{
			Id:       userRecords[i][0],
			Name:     userRecords[i][1],
			Email:    userRecords[i][2],
			Phone:    phone,
			IsActive: isActive,
		}
		users = append(users, user)
	}
	return users
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	controller.GetUsers(&users, log)
	jsonStr, _ := json.Marshal(users)
	_, err := fmt.Fprintln(w, string(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	users := CovertCSVToModel(r.URL.Query().Get("filename"))
	for i := 0; i < len(users); i++ {
		controller.CreateUser(&users[i], log)
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", GetUsers)
	myRouter.HandleFunc("/create/users", CreateUsers)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
