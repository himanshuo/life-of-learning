//learning concurrancy in golang. 
//apparently trivial. apparently fun.
//will evaluate at end.
package main 
import (
	"fmt"
	"os"
	"sort"
)


//dumb function.
func doSomethingForAWhile() {
	var a int
			for i:=0; i<1;i++{
				a++
			}

}


//Share by communicating

//Problem with concurrent programming in other languages: subtleties required to implement correct access to shared variables. 
//Why Go is awesome; go allows for shared values to be passed around on channels!!
// by passing around in channels, you never actively share a value by separate threads of execution. 
//	THUS, only one goroutine has access to the value at any given time. 
//	THEREFORE, Data races cannot occur, by design. 
//	SLOGAN: Do not communicate by sharing memory; instead, share memory by communicating.

// there are some cases where you dont like this:
//	Reference counts may be best done by putting a mutex around an integer variable. 
//	But as a high-level approach, using channels to control access makes it easier to write clear, correct programs.

/*
Consider a typical single-threaded program running on one CPU:
	No need for synchronization primitives. 
Now run another such instance: 
	still no need for synchronization. 
Now let those two communicate:
	if the communication is the synchronizer, there's still no need for other synchronization.
 */


//Goroutines

/*
A goroutine has a simple model: 
	it is a function executing concurrently(at the same time as) with other goroutines in the same address space. 
	lightweight - costing little more than the allocation of stack space. 
	stacks start small, so they are cheap, and grow by allocating (and freeing) heap storage as required.

Goroutines are "multiplexed" onto multiple OS threads so if one should block, such as while waiting for I/O, others continue to run. 
	This design hides many of the complexities of thread creation and management.
	Multiplexing - sending multiple signals or streams of info at the same time in the form of a single, complex signal. Then, recover the separate signals at the receiving end.

how to use a goroutine:
	Put the "go" keyword before a function call to call a new goroutine. 
	When the call completes, the goroutine exits, silently. 
		Similar to Unix shell's & notation for running a command in the background.

example:
	go list.Sort()  // run list.Sort concurrently; don't wait for it.
	
	A function literal (closure) can be handy in a goroutine invocation. Also, this is a anonymous function.
		func Announce(message string, delay time.Duration) {
	    	go func() {
	        	time.Sleep(delay)
	        	fmt.Println(message)
	    	}()  // Note the parentheses - must call the function.
		}
	In Go, function literals are closures: 
		the implementation makes sure the variables referred to by the function survive as long as the variables are being used in the function.

These examples aren't too practical because the functions have no way of signaling completion. For that, we need channels.
*/

func useGoroutine(){
	go func (){
			doSomethingForAWhile()
			fmt.Println("background task done.")
		}()
	fmt.Println("foreground task done.")
	doSomethingForAWhile()
	
}



//Channels

/*

Like maps, channels are allocated with make, and the resulting value acts as a reference to an underlying data structure. 
	If an optional integer parameter is provided, it sets the buffer size for the channel. 
		The default is zero, for an unbuffered or synchronous channel.
*/
func makeChannels(){
	//chan is type.
	//have a channel of some type. In these cases, channel of integers, and channel  of pointers to files.
	//use make keyword with optional int param for buffer size

	ci := make(chan int)            // unbuffered channel of integers
	cj := make(chan int, 0)         // unbuffered channel of integers
	cs := make(chan *os.File, 100)  // buffered channel of pointers to Files


	//this here is not part of given code. this is just to use ci,cj,cs.
	go func(){
		ci <- 1
		cj <- 1
		cs <- nil
	}()

}

/*
Unbuffered channels combine communication and synchronization.
	communication - the exchange of a value
	synchronization - guaranteeing that two calculations (goroutines) are in a known state.
	I THINK this means that unbuffered channel (thus when you have make(chan <type>, buffer=0) and buffer size of 0 ), then you end up with
		the ability to exchange values, and can be sure that two functions that use the go keyword both exist.
		this makes sense. You need to be able to share the same value, and need to be confident that the functions that are sharing the
			value both exist.
*/


//IDIOMS using channels!!!!!

//1) channel allows us to wait for sort goroutine to complete.
func waitForSort(){
	c := make(chan int)  // Allocate a channel.
	
	// Start the sort in a goroutine; when it completes, signal on the channel.
	go func() {
		s := []int{5, 2, 6, 3, 1, 4} // unsorted
		sort.Sort(sort.IntSlice(s))
	    c <- 1  // Send a signal; value does not matter.
	}()
	
	doSomethingForAWhile()
	<-c   // Wait for sort to finish; discard sent value.   (THIS IS BLOCKING UNTIL the c channel gets a signal aka until the goroutine finishes.)
	//Receivers always block until there is data to receive. 
		//If the channel is unbuffered, the sender blocks until the receiver has received the value. 
		//If the channel has a buffer, the sender blocks only until the value has been copied to the buffer; 
			//if the buffer is full, this means waiting until some receiver has retrieved a value.
			//A buffered channel can be used like a semaphore, for instance to limit throughput. 
				//in this example, incoming requests are passed to handle, which sends a value into the channel, processes the request, 
				//and then receives a value from the channel to ready the “semaphore” for the next consumer. T
				//the capacity of the channel buffer limits the number of simultaneous calls to process.
}
//side test to see what the returned value of the channel is. is it the signal that we provide it?
func whatIsSignalVal(){
	c := make(chan string)
	go func(){
			doSomethingForAWhile()
			c <- "hi, print this please."
		}()
	doSomethingForAWhile()

	fmt.Println(<-c)
}

//2) use channel as a semaphore
/*

SIMPLE WAY TO UNDERSTAND WHATS GOING ON:
	in the below code, what is happening is that 
		1) we make a semaphore which is just channel with some int representing the maximum amount of connections we want to be handled at the same time.
			can think of this literally as tornado async connections. the buffer represents the max number of users we want connected to our server.
		2) the Serve function is the MainHandler in tornado. This takes in requests. It sends them to the handle function.
			the handle function works in the background. THUS, it allows the Serve function to handle other requests. 
		3) the handle function then actually handles the requests.
			however, the handle function needs to be blocking. How do we make it blocking? Semaphore.
			the way channel buffer works is that it just will make all requests wait until something in the buffer gets opened up.

If the channel has a buffer, the sender blocks only until the value has been copied to the buffer; 
	if the buffer is full, this means waiting until some receiver has retrieved a value.
	A buffered channel can be used like a semaphore, for instance to limit throughput. 
		in this example, incoming requests are passed to handle, which sends a value into the channel, processes the request, 
		and then receives a value from the channel to ready the “semaphore” for the next consumer.
		the capacity of the channel buffer limits the number of simultaneous calls to process.
}
*/

// func channelAsSemaphore(){
// 	var sem = make(chan int, MaxOutstanding)    //create the semaphore with some MaxOutstanding (max number of connections)

// 	func handle(r *Request) {
// 	    sem <- 1    // Wait for active queue to drain.
// 	    process(r)  // May take a long time.
// 	    <-sem       // Done; enable next request to run.
		//OHHHHH, the point is that the semaphore basically doesnt allow other things to run while "process(r)" is running.
			//remember, <-sem blocks. It doesnt allow you to go past <-sem point until the previous stuff is done. THUS, process(r) must complete.
// 	}

// 	func Serve(queue chan *Request) {
// 	    for {
// 	        req := <-queue	//just keep adding all user requests onto a queue.
// 	        go handle(req)  // Don't wait for handle to finish.
// 	    }
// }
/*
upto MaxOutstanding handlers (say, 5 handlers) are executing process.
	any more will block trying to send into the filled channel buffer.
	until one of the existing handlers finishes and receives from the buffer.

//This design has a problem, though: 
	Serve creates a new goroutine for every incoming request, even though only MaxOutstanding of them can run at any moment. 
	As a result, the program can consume unlimited resources if the requests come in too fast. 
	We can address that deficiency by changing Serve to gate the creation of the goroutines. 
	Here's an obvious solution, but beware it has a bug we'll fix subsequently:


func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func() {
            process(req) // Buggy; see explanation below. 
            <-sem
        }()
    }
}
The bug is that in a Go for loop, the loop variable is reused for each iteration, so the req variable is shared across all goroutines. 
That's not what we want. We need to make sure that req is unique for each goroutine. 
	PROBLEM is that the closure will use req as whatever the function has defined req to be. 
		WHY? because req is not a param to the closure func()
			instead, req is a variable in the function that has func() as a closure, thus func() will use whatever value of req
			is the current value of req. req is constantly changing though. It is inside a for loop in which there is a function that runs
			in the background. THUS, what will happen is that as soon as Serve runs, the for loop will very quickly go through the 
			entire queue. At each iteration, func() will be started in the background. HOWEVER, when the process(req) starts, by that time, 
			the for loop will be done. Thus, req will now just be the last element in the queue. THUS, the func() closure will end up just
			using that last req value for each time it is run.
			THEREFORE, result will be running func() on the last value of queue len(queue) times.

Here's one way to fix this, passing the value of req as an argument to the closure in the goroutine:

func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func(req *Request) {
            process(req)
            <-sem
        }(req)
    }
}
Compare this version with the previous to see the difference in how the closure is declared and run.
What happens here is that the closure will in the background but each time have the current req value in the for loop passed into it.

Another solution is just to create a new variable with the same name, as in this example:
func Serve(queue chan *Request) {
    for req := range queue {
        req := req // Create new instance of req for the goroutine.
        sem <- 1
        go func() {
            process(req)
            <-sem
        }()
    }
}

What happens here is that a new variable req is created. Now,  when the closure refers to req, it will use the new req created by req := req
HOWEVER, I AM UNSURE OF WHAT IS GOING ON???
SHOULDN'T the new req just go out of scope? When the closure is called, it will use whatever value of req exists for it to use. The value 
that will exist will be the last req out there.

}
*/



// func channelAsSemaphore(){
// 	var sem = make(chan int, MaxOutstanding)    //create the semaphore with some MaxOutstanding (max number of connections)

// 	func handle(r *Request) {
// 	    sem <- 1    // Wait for active queue to drain.
// 	    process(r)  // May take a long time.
// 	    <-sem       // Done; enable next request to run.
		//OHHHHH, the point is that the semaphore basically doesnt allow other things to run while "process(r)" is running.
			//remember, <-sem blocks. It doesnt allow you to go past <-sem point until the previous stuff is done. THUS, process(r) must complete.
// 	}

// 	func Serve(queue chan *Request) {
// 	    for {
// 	        req := <-queue	//just keep adding all user requests onto a queue.
// 	        go handle(req)  // Don't wait for handle to finish.
// 	    }
// }




func main(){
	//INTERESTING. most of the time, this only prints foreground task done.
	//WHY? because the useGoroutine returns and thus allows main to return and stop program execution before the background task is done.
	useGoroutine()
	whatIsSignalVal()//YUP! "signal" is the thing you send to the channel.

}



