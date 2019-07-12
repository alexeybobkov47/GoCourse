package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	time.Sleep(10 * time.Second)
	fmt.Printf("\nПрошло 10 секунд\n")

	/*	for x := 0; x <= 10; x++ {
			t := time.Now().Format("15:04:05\n\r")
			time.Sleep(1 * time.Second)
			fmt.Print("\n", t)
		}
	*/
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)

			time.Sleep(delay)

		}
	}
}
