package main
import (
	"fmt"
	"time"
)

/*scenario: 
start fast and slow run at same time. 
As soon as slow finishes, it sends a signal to fast that lets slow go faster.


non sharing version: 
	fast -> 1
	slow -> 1 2 3 4 5

sharing version:
	fast -> 1(transfers as soon as 1st second is over)
	slow -> 1 2 (once slow recieves transfer, it goes for only 1 sec instead of 4 more.)  


*/
func slow(c chan int, done chan int){
	//keep sleeping (5 sec max)
	for a :=0; a<5;a++{
		fmt.Println(a)
		
		
		//recieve from the channel.
		select {
			case <-c:
				fmt.Println(a,">-----forward------>", 999)
				a = 999
				break
			default:
				time.Sleep(time.Second)
		} 
		
	}
	
	fmt.Println("slow is done")
	done <- 1
	
}
func fast(c chan int){
	//sleep
	time.Sleep(time.Second*2)

	fmt.Println("fast is done")	
	//send a value into the channel
	c <- 9999
}	


func main() {
	//make the channel. It will be an int channel.
	c , done:= make(chan int),make(chan int)
	//slow and fast are run at the same time.
	go slow(c,done)
	go fast(c)

	_=<-done
	
	
}
