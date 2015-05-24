package main
import (
	"fmt"
	"time"
)


func slow(c chan int){
	time.Sleep(time.Second*5)
	fmt.Println("slow is done")
	c <- 5
}
func fast(c chan int){
	time.Sleep(time.Second)
	fmt.Println("fast is done")	
	c <- 1
}	


func main() {
	c := make(chan int,2)
	go slow(c)
	go fast(c)
	_,_ = <-c, <-c	
	
}
