package leetcode
import (
	"strconv"
	//"fmt"
)
/*
 * 1) Reverse an integer
 * 2-3)  
 * 	12-> 21
 * 	-12 -> -21
 * 	99910 -> 1999
 *	10000-> 1
 *  MAX_INT (0111...) -> (1111..0)
 *  MIN_INT (1111...) -> (NOT POSSIBLE)
 *  no decimals
 *  23e1 = 23x10^1 = 230 -> 32
 *  For overflows, return 0
 * 
 * 
 * 
 * 4) ideas:
 * 		1: turn into string, while iterating through string in  reverse
 * 			sum = 0
 * 			sum = 10^str_index * to_int(str[str_index]) + sum
 * 		...
 * 		if overflow: return 0   
 */

const MaxUint = ^uint(0) //1111...
const MaxInt = int(MaxUint >> 1) //0111...
const MinInt = -MaxInt - 1 //1000...

//func abs(n int) int{
	//if n < 0{
		//return n*-1
	//} else {
		//return n
	//}
//}

func to_int(n rune) int {
	conversion := map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	if res, ok := conversion[n]; ok {
		return res
	} else {
		panic("incorrect input to to_int")
	}
	
}

func overflow(original int, toAdd int, neg bool) bool{
	original = abs(original)
	if neg {
		return MinInt + original >  -1*toAdd 
		//return toAdd > MinInt - original 
	} else {
		return toAdd > MaxInt - original
		//return MaxInt - toAdd > original
	}
}

func pow(base int, power int) int {
	total := 1
	for i:=0; i < power; i++{
		total *= base
	}
	return total
}

func ReverseInteger(n int) int {
	str := strconv.Itoa(abs(n))
	sum := 0
	for i := len(str)-1; i>-1; i--{
		toAdd := pow(10,i) * to_int(rune(str[i]))
		if overflow(sum, toAdd, n < 0){
			return 0
		} else {
			sum += toAdd 
		}
	}
	
	if n < 0 {
		sum = -1*sum
	}
	return sum
}
