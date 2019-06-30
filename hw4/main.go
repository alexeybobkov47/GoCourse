package main

import "fmt"

type car struct {
	model   string // марка
	release int    // год выпуска
	fuel    int    // потрачено бензина за день
}

var fuelprice = 42.34 // цена за 1ед топлива

type airplane struct {
	model   string // марка
	release int    // год выпуска
	flights int    // кол-во рейсов за день
}

var flightprice = 5312 // цена за 1 рейс

// затраты за день на car
func (c car) cost() float64 {
	return fuelprice * float64(c.fuel)
}
func (c *car) setFuel(fu int) {
	c.fuel = fu
}

// затраты за день на airplane
func (a airplane) cost() float64 {
	return float64(flightprice) * float64(a.flights)
}
func (a *airplane) setFlights(fl int) {
	a.flights = fl
}

type spending interface {
	cost() float64
}

// Сумма затрат на car and airplane
func summcost(spendingX ...spending) float64 {
	res := 0.0
	for _, spending := range spendingX {
		res += spending.cost()
	}
	return res
}
func main() {

	var c1, c2, c3, c4 car
	var a1, a2, a3, a4 airplane
	c1 = car{"mw1", 2019, 434}
	c2 = car{"mw2", 2019, 134}
	c3 = car{"mw3", 2019, 531}
	c4 = car{"mw1", 2019, 231}
	a1 = airplane{"warcry767", 1985, 3}
	a2 = airplane{"warcry767", 1985, 1}
	a3 = airplane{"warcry767", 1985, 5}
	a4 = airplane{"warcry767", 1985, 4}
	fmt.Printf("c1: %.2f, c2: %.2f, c3: %.2f, c4: %.2f\n", c1.cost(), c2.cost(), c3.cost(), c4.cost())
	fmt.Println(a1.cost(), a2.cost(), a3.cost(), a4.cost())
	fmt.Printf("Сумма затрат за день на весь транспорт составила: %.2f\n", summcost(c1, c2, c3, c4, a1, a2, a3, a4))

	/*
	   c1: 18375.56, c2: 5673.56, c3: 22482.54, c4: 9780.54
	   15936 5312 26560 21248
	   Сумма затрат за день на весь транспорт составила: 125368.20
	*/

	sortcars()
}
