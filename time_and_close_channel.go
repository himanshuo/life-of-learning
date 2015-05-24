package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	fmt.Println("start:",time.Now())
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		fmt.Println(x)
		x, y = y, x+y
	}
	//close(c)
	fmt.Println("stop:",time.Now())
}





func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i,"-----------")
	}
}
