package io

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type FileHandler interface {
	ReadFile()
	WriteFile()
	CreateFile()
}

func WriteToJsonFile(jsonStr string)  {
	f, err :=os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString(jsonStr)
}

func ReadCSVFile(fileName string) *csv.Reader {
	userRecords, err := os.Open(fileName)
	csvReader := csv.NewReader(userRecords)
	if err != nil {
		fmt.Println(err)
	}
	return csvReader
}

func GetRecordsFromCSVReader(fileReader *csv.Reader) [][]string{
	var records [][]string
	for {
		record, err := fileReader.Read()
		records = append(records, record)
		if err == io.EOF {
			break
		}
		if err == nil{
			//fmt.Println(record)
		}
	}
	return records
}