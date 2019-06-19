package p2

import (
	"fmt"
	"math"
)

// P2 - Функция к практике 2
func P2() {
	fmt.Println("Расчет гипотенузы, периметра и площади треугольника при известных катетах.")
	const a float64 = 35
	const b float64 = 44
	var c = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	fmt.Println("Катеты известы:", a, b)
	fmt.Println("Гипотенуза равна:", c)
	var p = a + b + c
	fmt.Println("Периметр равен: ", p)
	var pp = p / 2
	fmt.Println("Полупериметр равен: ", pp)
	var formulagerona = math.Sqrt(pp * (pp - a) * (pp - b) * (pp - c))
	fmt.Println("Площадь треугольника равна: ", formulagerona)

}
