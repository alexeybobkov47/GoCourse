package main

import (
	"encoding/csv"
	"fmt"

	"io"
	"os"
)

type cars struct {
	year     string `csv:"year"`
	carbrand string `csv:"carbrand"`
	carmodel string `csv:"carmodel"`
	comment  string `csv:"comment"`
	price    string `csv:"price"`
}

func main() {
	csvFile, err := os.Open("test.csv")
	if err != nil {
		return
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	for {
		var car []cars

		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return
		}
		fmt.Println(record)

		car = append(car, cars{
			year:     record[0],
			carbrand: record[1],
			carmodel: record[2],
			comment:  record[3],
			price:    record[4],
		})
		fmt.Println(car)

	}
}
