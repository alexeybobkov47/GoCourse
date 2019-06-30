package main

import (
	"fmt"
	"sort"
)

type byFuel []car

func (a byFuel) Len() int           { return len(a) }
func (a byFuel) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFuel) Less(i, j int) bool { return a[i].fuel < a[j].fuel }

func sortcars() {

	cars := []car{
		{"mw1", 2019, 434},
		{"mw2", 1987, 134},
		{"mw3", 2005, 531},
		{"mw1", 2015, 231},
	}
	fmt.Println(cars)
	//[{mw1 2019 434} {mw2 1987 134} {mw3 2005 531} {mw1 2015 231}]
	sort.Sort(byFuel(cars))
	fmt.Println(cars)
	//[{mw2 1987 134} {mw1 2015 231} {mw1 2019 434} {mw3 2005 531}]

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].release > cars[j].release
	})
	fmt.Println(cars)
	//[{mw1 2019 434} {mw1 2015 231} {mw3 2005 531} {mw2 1987 134}]
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].fuel > cars[j].fuel
	})
	fmt.Println(cars)
	//[{mw3 2005 531} {mw1 2019 434} {mw1 2015 231} {mw2 1987 134}]
}
