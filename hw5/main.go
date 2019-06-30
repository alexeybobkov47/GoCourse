package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 10008934500; i++ {
	}
	duration := time.Since(start)
	fmt.Printf("Operation took: %s", duration)
}
