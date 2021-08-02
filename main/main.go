package main

import (
	"goassignment/io"
	"goassignment/parser"
)

func main() {
	records := io.GetRecordsFromCSVReader(io.ReadCSVFile("userrecords/user_records.csv"))
	io.WriteToJsonFile(
		parser.GetJsonFromRecordObjects(
			parser.GetRecordObjectsFromCSVReader(records)))
}
