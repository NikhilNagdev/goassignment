package model

type User struct {
	Id string
	Name string
	Email string
	Phone int
	IsActive int
}

//func CreateUser(user *User)  {
//	err := database.GetDbConnection().Create(user).Error
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//
//func GetAllUsers(user *[]User)  {
//	err := database.GetDbConnection().Find(user).Error
//	if err != nil{
//		fmt.Println(err)
//	}
//	fmt.Println(database.GetDbConnection().Find(user))
//}
