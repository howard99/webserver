package main

import (
	"fmt"
)

//
//Buffered channels   ch :=make(chan int,100)
//
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
