package statistic

//Average - подсчет среднего.
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

//Summ - подсчет суммы.
func Summ(a int, b int) int {
	c := a + b
	return c
}

//Summ2 - подсчет сумм
func Summ2(xs []int) int {
	total := int(0)
	for _, x := range xs {
		total += x
	}
	return total
}
