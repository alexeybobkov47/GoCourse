package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"sort"

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

var car []cars

func main() {
	csvFile, err := os.Open("test.csv")
	if err != nil {
		return
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	for {

		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return
		}
		car = append(car, cars{
			year: record[0], carbrand: record[1], carmodel: record[2], comment: record[3], price: record[4],
		})
	}

	sort.Slice(car, func(i, j int) bool {
		return car[i].carmodel > car[j].carmodel
	})
	fmt.Println(car)
	sort.Slice(car, func(i, j int) bool {
		return car[i].price < car[j].price

	})
	fmt.Println(car)
	sort.Slice(car, func(i, j int) bool {
		return car[i].year < car[j].year

	})
	fmt.Println(car)

	/*
	   [{1999 Chevy Venture «Extended Edition»  4900.00} {1996 Jeep Grand Cherokee MUST SELL! air, moon roof, loaded 4799.00} {1997 Ford E350 ac, abs, moon 3000.00}]
	   [{1997 Ford E350 ac, abs, moon 3000.00} {1996 Jeep Grand Cherokee MUST SELL! air, moon roof, loaded 4799.00} {1999 Chevy Venture «Extended Edition»  4900.00}]
	   [{1996 Jeep Grand Cherokee MUST SELL! air, moon roof, loaded 4799.00} {1997 Ford E350 ac, abs, moon 3000.00} {1999 Chevy Venture «Extended Edition»  4900.00}]

	*/

	openfile()

}
func openfile() {

	fi, _ := os.Stat("main.go")

	if fi.Size() > 1000 {
		fmt.Println("Слишком большой файл. Размер файла:", fi.Size())
	} else {
		bs, err := ioutil.ReadFile("main.go")
		if err != nil {
			return
		}
		fmt.Println(string(bs))
	}

}
