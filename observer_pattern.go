package main
import (
	"fmt"
	"time"
)
//observer pattern in go


/*SCENARIO 1:
 * 	master has two children.
 * 		child A screams "A" every second
 * 		child B screams "B" every 2 seconds
 * 
 * master gets annoyed when both children have screamed 10 times and then he yells "STOP DUDES"
 * 0 1 2 3 4 5 6
 * A   A   A   A
 * B     B     B
 * DESIGN: master function waits on channels sent by children in case statement.
 * 
 * 
 */
 
 func child(name string, c chan string){
	//keep sleeping (5 sec max)
	a :=0
	for {
		c <- name
		a = a + 1
		switch name {
			case "A":
				time.Sleep(time.Second)
				break
			case "B":
				time.Sleep(time.Second*2)
				break
		}
	}
	

	
	
}

func master(c chan string){
	//numA := 0
	//numB := 0
	for{
		name := <- c
	    switch name {
			case "A":
				fmt.Println("A")
				break
			case "B":
				fmt.Println("B")
				break
			default:
				fmt.Println("error.")
		}
	}
	
}

func main(){
		c := make(chan string)
		go child("A", c)
		go master(c)
		go child("B",c)
		
		time.Sleep(time.Second*6) //let the thing above playout for 5 seconds
		
}
 
 
 /*SCENARIO 2:
 * 	master has two children.
 * 		child A screams "A" every x seconds
 * 		child B screams "B" every x seconds
 * 
 * master gets annoyed when both children have screamed 10 times and then he yells "STOP DUDES"
 */
 
  /*SCENARIO 3:
 * 	master has two children.
 * 		child A screams "A" every x seconds
 * 		child B screams "B" every y seconds
 * 
 * master gets annoyed when both children have screamed 10 times and then he yells "STOP DUDES"
 */
