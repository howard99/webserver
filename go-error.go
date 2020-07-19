package main

import (
	"fmt"
	"time"
)

type MyError struct {
	when string
	what string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v,%s", e.when, e.what)
}

func run() error {
	return &MyError{
		time.Now().Format("2006-01-02 15:04:05"), "it didn't work",
	}
}

//struct   error
//
//

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
