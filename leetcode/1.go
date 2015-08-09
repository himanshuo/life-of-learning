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

func findLargestSigfigLength(list []int) int {
	if len(list) <1 {
		panic("no elements in list")
	}
	largest := list[0]
	for _,v := range list {
		val := v
		if val < 0 {
			val = val*-1
		}
		if val > largest {
			largest = val
		}
	}
	
	return len(strconv.Itoa(largest))
}

func valAtSigfig(num int, sigfig int) int{
	if len(strconv.Itoa(num)) > sigfig {
		valString := string(strconv.Itoa(num)[sigfig])
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
	for i:=0; i<1; i++ {
		buckets = append(buckets, make([]int, 0)) 
	}
	return buckets
}

func RadixSort( list []int) []int {
	if len(list) < 2{
		return list
	}
	out := make([]int, len(list))
	buckets := createBuckets()
	largestSigfig := findLargestSigfigLength(list)
	for sigfig :=0; sigfig < largestSigfig; sigfig++ {
		for _,v := range list {
			val := valAtSigfig(v,sigfig)
			buckets[val] = append(buckets[val], v)
		}
		out = joinBucketsInOrder(buckets)
	}
	return out
}

func main(){

}
