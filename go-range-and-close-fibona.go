package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//兔子系数，费波纳希舒x，y =y,x+y
func main() {
	c := make(chan int, 16)

	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
