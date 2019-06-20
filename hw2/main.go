package main

import "fmt"

func main() {
	var number int
	fmt.Println("Введите число для проверки на четность и деление на 3:")
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
		fmt.Print(" и не делится без остатся на 3\n")
	}

	// 100 первых чисел Фибоначчи
	fmt.Println("Вывод 100 чисел Фибоначчи")
	for a, b, c := 0, 1, 0; c < 100; a, b = a+b, a {
		c = c + 1
		fmt.Println(c, a)

	}
}
