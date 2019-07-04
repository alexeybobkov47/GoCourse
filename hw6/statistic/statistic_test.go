package statistic

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.50},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverageSet(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

// ТЕСТ СУММ 1
type testpair2 struct {
	values1 int
	values2 int
	summ    int
}

var tests2 = []testpair2{
	{1, 2, 3},
	{33, 123, 156},
	{55, 7, 62},
}

func TestSumm(t *testing.T) {
	for _, ss := range tests2 {
		s := Summ(ss.values1, ss.values2)
		if s != ss.summ {
			t.Error(
				"For", ss.values1, ss.values2,
				"expected", ss.summ,
				"got", s,
			)
		}
	}
}

// ТЕСТ СУММ 2

type testpair3 struct {
	values []int
	summ   int
}

var tests3 = []testpair3{
	{[]int{1, 2, 3, 4, 5}, 15},
	{[]int{1, 1, 1, 1, 2, 1}, 7},
	{[]int{-1, 1, 3, -3}, 0},
}

func TestSumm2(t *testing.T) {
	for _, ss := range tests3 {
		s := Summ2(ss.values)
		if s != ss.summ {
			t.Error(
				"For", ss.values,
				"expected", ss.summ,
				"got", s,
			)
		}
	}
}
