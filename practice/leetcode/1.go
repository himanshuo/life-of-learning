package leetcode
import (
	"strconv"
)
/*
1) find two nums from int list that add to given target number.
* function is called twoSum.
* function returns indices of two desired nums
* index1 < index2
* indexes are NOT zero-based. they start at 1.
* ASSUME: input only has 1 solution. NEVER 0 or multiple solutions.
2) [2,7,11,15],9 -> 1,2
*  [1,2,3],3 -> 1,2
*  [3,5,6,2],9 -> 1,3
*  [-1,-3,0,1],1 -> 4,1 
3) [0,0,0,0],1 -> error
*  [1,2,3,4],55 -> error
*  [0],0 -> 1,1
*  [0,1],0 -> 1,1
*  []
4) ideas:
* 	1) for each num, check all things to its right -> O(n^2) time, O(1)space
 *  2) radixsort list. move inwards from each end until sum matches or is too low. -> O(lgn) O(1)
 *
 * 
5) 
*   func twoSum(list, target)
* 		list = sort(list)
*   	l := 0
* 		r := len(list)
* 		while list[l]+list[r] !=  target: 			
* 			if r > l:
* 				error
* 			else if list[l+1]-list[l] < list[r]-list[r-1]:
* 				l = l+1
* 			else:
* 				r = r-1
*		return l,r
* 	//0,20,15,100,1 initial
* 	//0,20,100,1,15 1st round
* 	//0,100,1,15,20	2nd round
* 	//0,1,15,20,100 3rd round
* func radixSort(list):
* 		buckets = make([][]int, 9)
* 		for sigfig in 0:len(largest_num_length)
* 			for num in list:
* 				val = valAtSigfig(num,sigfig)
* 				buckets[val].append(num)
* 			list = join_buckets_in_order()
* //handle neg numbers with extra buckets
* 			 		
*/ 

//in this case: sigfig is determined by
// 4326 has 4 at sigfig 1. 2 at sigfig 3


func abs(n int) int{
	if n< 0{
		return n*-1
	} else{
		return n
	}
}

func findLargestSigfigLength(list []int) int {
	if len(list) <1 {
		panic("no elements in list")
	}
	largest := abs(list[0])
	for _,v := range list {
		if abs(v) > largest {
			largest = abs(v)
		}
	}
	
	return len(strconv.Itoa(largest))
}
//valAtSigfig(1433, 3) = 1
//valAtSigfig(0,0) = 0
//valAtSigfig(12,2)=0
func valAtSigfig(num int, sigfig int) int{
	num = abs(num)
	numAsString := strconv.Itoa(num)
	if len(numAsString) > sigfig {
		index := len(numAsString) - sigfig - 1
		valString := string(numAsString[index])
		val,_ := strconv.Atoi(valString)
		return val
	} else {
		return 0
	}
}
func extend(list []int, toAdd []int) []int{
	for _,v := range toAdd{
		list = append(list, v)
	}
	return list

}
func joinBucketsInOrder(buckets [][]int) []int{
	out := make([]int, 0)
	for _,v := range buckets{
		out = extend(out, v)
	}
	return out
}
func createBuckets()[][]int{
	buckets := make([][]int, 0)
	for i:=0; i<10; i++ {
		buckets = append(buckets, make([]int, 0)) 
	}
	return buckets
}

func appendToFront(list []int, num int) []int {
	out := make([]int, 0)
	out = append(out, num)
	for _,v := range list{
		out = append(out, v)
	}
	return out
}

//9, 2, -2, 29, 19  
//2, -2, 9, 29, 19
//2, -2, 9, 19, 29
//-2, 2, 9, 19, 29
func RadixSort( list []int) []int {
	if len(list) < 2{
		return list
	}
	out := make([]int, len(list))
	copy(out, list)
	largestSigfig := findLargestSigfigLength(list)
	for sigfig :=0; sigfig < largestSigfig; sigfig++ {
		buckets := createBuckets()
		for _,v := range out {
			val := valAtSigfig(v,sigfig)
			buckets[val] = append(buckets[val], v)
			
		}
		out = joinBucketsInOrder(buckets)
	}
	temp := make([]int, len(out))
	copy(temp, out)
	out = make([]int, 0)
	// 3,5,-6, 45, -345 -> -345, -6, 3,5,45
	for _,v := range temp {
		if v < 0 {
			out = appendToFront(out, v)
		} else{
			out = append(out, v)
		}
	}
	return out
}

func indexOfFirstInstance(list []int, num int) int {
	for i,v := range list{
		if v == num{
			return i
		}
	}
	panic("num not in list")
}
func indexOfSecondInstance(list []int, num int) int {
	i := indexOfFirstInstance(list, num) + 1
	for ; i < len(list); i++{
		if num==list[i]{
			return i
		}
	}
	panic("num does not exist in list twice.")
}

func TwoSum(list []int, target int) (int, int){
	if len(list) < 1 {
		panic("list too small")
	}
	sortedList := RadixSort(list)
   	l := 0
 	r := len(sortedList) - 1
 	for ;sortedList[l] + sortedList[r] !=  target; { 			
 			if l > r {
 				panic("shouldn't occur")
 			} else if sortedList[l] + sortedList[r] < target{
 				l = l+1
 			} else {
 				r = r-1
 			}
 	}
 	index1 := indexOfFirstInstance(list,sortedList[l])
 	var index2 int
 	if sortedList[l] == sortedList[r] {
		index2 = indexOfSecondInstance(list,sortedList[r])
	} else {
		index2 = indexOfFirstInstance(list,sortedList[r])
	}
	
	if index1 < index2 {
		return index1, index2
	} else {
		return index2, index1
	}
 	
}
