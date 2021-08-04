package record

import (
	"fmt"
	gouuid "github.com/satori/go.uuid"
	"goassignment/validator"
	"strconv"
)

type Record struct {
	Id string
	Name string
	Email string
	Phone int
	IsActive bool
}

var recordMap = map[string]Record{}

func (record *Record) ProcessRecordId(){
	_, v := recordMap[record.Id]
	fmt.Println(record)
	if !v{
		fmt.Println("Not same")
		recordMap[record.Id] = *record
		fmt.Println(recordMap)
	}else{
		myuuid := gouuid.NewV4()
		record.setId(myuuid.String())
		recordMap[myuuid.String()] = *record
	}
}

func (record *Record) setId(id string) {
	record.Id = id
}

func (record *Record) ValidatePhoneNo() bool {
	return validator.ValidatePhoneNo(strconv.Itoa(record.Phone))
}

func (record *Record) ValidateName() bool {
	return validator.ValidateName(record.Name)
}

func(record *Record) ValidateEmail() bool {
	return validator.ValidateEmail(record.Email)
}

func (record *Record) IsRecordCorrect() bool{
	return record.ValidatePhoneNo() && record.ValidateName() && record.ValidateEmail()
}

