package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; x <= 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// возведение в квадрат
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // канал закрыт и пуст
			}
			squares <- x * x
		}
		close(squares)
	}()

	// печать
	for {
		x, ok := <-squares
		if !ok {
			break // канал закрыт и пуст
		}
		fmt.Println(x)
	}
}
