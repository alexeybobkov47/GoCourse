package p3

import "fmt"

// P3 - для 3 практического задания
func P3() {
	var deposit float64
	fmt.Println("Введите сумму вклада в банк")
	fmt.Scanln(&deposit)
	fmt.Println("Введите годовой процент")
	var procent float64
	fmt.Scanln(&procent)
	const period float64 = 5 * 365
	var income = ((deposit * procent * period) / (365 * 100))
	fmt.Println("Доход по процентам за 5 лет составит:", income)
	fmt.Println("Сумма на счете с учетом дохода через 5 лет:", income+deposit)
}
