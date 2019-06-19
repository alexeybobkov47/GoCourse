package main

import "fmt"

func main() {
	var number int
	fmt.Println("Введите число:")
	fmt.Scanln(&number)
	// Проверка четное\нечетное
	if number%2 == 0 {
		fmt.Print("Число ", number, " четное")
	} else {
		fmt.Print("Число ", number, " нечетное")
	}
	// Проверка деления на 3
	if number%3 == 0 {
		fmt.Print(" и делится без остатка на 3")
	} else {
		fmt.Print(" и не делится без остатся на 3")
	}

}
