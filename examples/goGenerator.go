package main
import ( 
	"fmt"
	"math/rand"
)
// prime returns true if n is a prime number.
func prime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// primes returns a channel of ints on which it writes the first n prime
// numbers before closing it.
func primes(n int) chan int {
	c := make(chan int)
	go func() {
		numPrimes := 0
		for i:=0;numPrimes<n;i++ {
	  	if prime(i){
			c <- i	
			numPrimes += 1
			}
		}
		close(c)
	}()
	
	return c
		
}



func generator(start int, funcToNext func(int) (int), stopCondition func(int) (bool)) chan int{
	c := make(chan int)
	go func(){
		current := start
		for ;!stopCondition(current); {
				current = funcToNext(current)
				c <- current
			}
		close(c)
	}()
	
	return c
}

func nextRandomNum(in int) int {
	a := rand.Int()
	fmt.Println("inside nextRandomNum", a)
	return a
}
func isEven(n int) bool {
	return n%2 == 0
}

func isDivisibleByThree(n int) bool {
	return n%3 == 0
}

func nextPrimeNum(n int) int {
	for i:= n+1 ; ;i++{
		if prime(i){
			return i
		}
	}
}

var numPrimes = 0

func numPrimesIsTen(n int) bool{
	return numPrimes == 10
}

func main() {
	
	for p := range primes(10) {
		fmt.Println(p)
	}
	
	for p := range generator(-1, nextPrimeNum, numPrimesIsTen){
		fmt.Println(p)
		numPrimes += 1
	}
	
}
