package main

import (
	"fmt"
	"os"
	"time"
)

//MyError some error
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("В %v произошло событие %s", e.When, e.What)
}

func runWithError() *MyError {
	return &MyError{
		time.Now(), "it didn't work",
	}
}

func runWithOk() *MyError {
	return &MyError{
		time.Now(), "it works again",
	}
}

//include in your main() myerrordemo()
func myerrordemo() {
	if err := runWithError(); err != nil {
		fmt.Println(err)
	}

	err1 := &MyError{time.Now(), "it works again"}
	fmt.Println(err1)

	file, err3 := os.Open("nevermind")
	if err3 != nil {
		fmt.Println(err3)
	} else {
		file.Close()
	}
}
