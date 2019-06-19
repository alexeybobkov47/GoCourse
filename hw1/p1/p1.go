package p1

import (
	"fmt"
	"strconv"
)

// Dollar - курс рубля к доллару
const dollar float64 = 64.33

var rub float64

// P1 - Функция к 1 ПЗ
func P1() {
	fmt.Println("Привет, с тобой конвектор рублей в доллары по курсу 64.33 за доллар,\nвведи кол-во рублей:")
	fmt.Scanln(&rub)
	var convert = rub / dollar
	// Конвертирование(форматирование) итога в 2 знака после запятой
	var convert2 = strconv.FormatFloat(float64(convert), 'f', 2, 64)
	fmt.Println("Получится в долларах: ", convert2)
}
