package validator

import (
	"fmt"
	"log"
	"os"
)

func ValidatePhoneNo(phoneNo string) bool{
	noOfDigits := len(phoneNo)
	if noOfDigits != 10 {
		file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		defer file.Close()
		log.SetOutput(file)
		log.Println("Phone number should be 10 digits")
		return false
	}
	return true
}
func ValidateName(name string) bool{
	if name == "" {
		file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		defer file.Close()
		log.SetOutput(file)
		log.Println("Name cannot be empty")
		return false
	}
	return true

}

func ValidateEmail(email string) bool{
	if email == ""{
		file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		defer file.Close()
		log.SetOutput(file)
		log.Println("Email cannot be empty")
		return false
	}
	return true

}


