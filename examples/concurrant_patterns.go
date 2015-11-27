package main
import (
	"fmt"
)

/*This file is understanding each concucrrancy pattern found on 
 * http://www.golangpatterns.info/concurrency/
 */

/* KEEP FORGETTING:
 * c := make(chan int)
 * c <- some_int  puts a value inside the channel
 * some_int <- c  retrieves a value once it exists from the channel. (blocks)
 */ 

/*Coroutines
 * differences between coroutines and goroutines:
 *	1) goroutine implies parallelism. With coroutines, you usually
 * 		end up with event loops which are not concurrant, not parallel.
 *	2) goroutines communicate via channels.
 * 		coroutines communicate via yield and resume operations
 * 
 */ 

// this function increases the output integer by 1 each time you call it
// thus it basically acts as a generator

func count() int{
	c := make(chan int) //make channel
	count := 0 //count
	
	go func() { // goroutine so infinite for loop doesnt block internalCount from returning channel
		for {
			c <- count  //blocks. puts count into channel.
			count += 1 //increase count
		}
	}()
	return <- c //return channel so count has access to it.
}



func main(){
	for i:=1; i<10; i++{
		fmt.Println(count())
	}

}

