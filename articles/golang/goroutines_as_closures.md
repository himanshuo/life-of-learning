Goroutines with closures
=========================

You can use goroutines with closures, allowing you to run a program in the main thread and have your variables all set up and defined there, BUT then even if the function has returned, the variables will exist in the goroutine.

Example:


	func myfunc(message String){
		//anonymous function
		//this anonymous function will run in the background.
		//thus as soon as myfunc is started, it will start this in the backgorund and return.
		//there is no time to sleep and print the message. Normally, youd think that message is nil in the anonymous function because myfunc already returned. But alas, thats the beaty of closures. message will remain in memory while the anonymous function is running.  
		go func(){ 
			time.Sleep(30)
			fmt.Println(message)
		}()
	}
	
	func main(){
		myfunc("hi") //func A will run and immidiently end.
		for {} //infinite loop so that main doesnt end
	}

