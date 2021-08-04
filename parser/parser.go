package parser

import (
	"encoding/json"
	"fmt"
	"goassignment/record"
	"strconv"
)

func GetJsonFromRecordObjects(records []record.Record) string {
	jsonString, err := json.Marshal(records)
	if err != nil {
		fmt.Println(err)
	}
	return string(jsonString)
}

func convertRecordToObject(csvRecords []string) record.Record {
	var recordObj = record.Record{}
	phone, _ := strconv.Atoi(csvRecords[3])
	isActive, _ := strconv.ParseBool(csvRecords[4])
	recordObj = record.Record{
		Id:       csvRecords[0],
		Name:     csvRecords[1],
		Email:    csvRecords[2],
		Phone:    phone,
		IsActive: isActive,
	}
	recordObj.ProcessRecordId()
	if !recordObj.IsRecordCorrect() {
		recordObj = record.Record{}
	}
	return recordObj
}

func GetRecordObjectsFromCSVReader(csvRecords [][]string) []record.Record {
	var records []record.Record
	for i := 0; i < len(csvRecords)-1; i++ {
		recordObj := convertRecordToObject(csvRecords[i])
		if(record.Record{} != recordObj){
			records = append(records, recordObj)
		}
	}
	return records
}
