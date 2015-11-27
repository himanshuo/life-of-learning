package main

import (
	"fmt"
	"time"
)

func say(s string) {
	fmt.Println(s)
	time.Sleep(100)
}

// func main() {
// 	// go say("world")
// 	// go say("hello")
// 	// fmt.Println("done")
// 	// time.Sleep(100)


	
// }

func sum(a []int, c chan int) {
	fmt.Println(time.Now(), "start")
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Println(time.Now(),"stop")
	c <- sum // send sum to c
	
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
