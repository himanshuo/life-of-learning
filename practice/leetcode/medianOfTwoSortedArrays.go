/*
 * 1) 2 sorted arrays of size m & n. Find the median of the two sorted arrays. 
 * TIME: O(log (m+n)).
 * 
 * 2-3) [1,2,3], [3,4,5] -> [1,2,3,3,4,5] -> 3
 * 	[1,2,3], [2,4] -> [1,2,2,3,4] -> 2
 * [],[1,2] -> [1,2] -> 1.5
 * [3,5],[] -> 4
 * [3,5],[1,2] -> [1,2,3,5] -> 2.5
 * [1,2,3],[0] -> [0,1,2,3] -> 1.5
 * [1,10],[5] -> [1,5,10] -> 5
 * 
 */
package leetcode
func MedianOfTwoSortedArrays(arrA []int, arrB []int) float64 {
	
	aMax := len(arrA) - 1
	bMax := len(arrB) - 1
	aMin := 0
	bMin := 0
	var newAMax, newAMin, newBMax, newBMin int
	
	if len(arrA) == 0 {
		return median(aMin, aMax, arrA) 
	} else if len(arrB) == 0 {
		return median(bMin, bMax, arrB)
	}
	
	for ;!smallRange(aMax, aMin, bMax, bMin); {
		
		if medianLess(aMin, aMax, arrA) < medianLess(bMin, bMax, arrB){
			newAMin = int(medianIndex(aMin, aMax, arrA)) // doesnt handle even/odd length of input
			newBMin = bMin
		} else {
			newAMin = aMin
			newBMin = int(medianIndex(bMin, bMax, arrB))// doesnt handle even/odd length of input
		}
		
		if medianMore(aMin, aMax, arrA) < medianMore(bMin, bMax, arrB){
			newAMax = int(medianIndex(aMin, aMax, arrA))  // doesnt handle even/odd length of input
			newBMax = bMax
		} else {
			newAMax = aMax
			newBMax = int(medianIndex(bMin, bMax, arrB)) // doesnt handle even/odd length of input
		}
		
		aMin = newAMin
		aMax = newAMax
		bMin = newBMin
		bMax = newBMax
		
	}
	
	var one int
	var two int
	if (aMax-aMin+1) == 2 {
		one = arrA[int(medianIndex(aMin, aMax, arrA))]
		two = arrA[int(medianIndex(aMin, aMax, arrA)) + 1]
		return (float64(two) - float64(one)) / float64(2)
	} else if (bMax-bMin+1) == 2 {
		one = arrB[int(medianIndex(bMin, bMax, arrB))]
		two = arrB[int(medianIndex(bMin, bMax, arrB)) + 1]
		return (float64(two) - float64(one)) / float64(2)
	} else if (aMax-aMin+1) + (bMax-bMin+1) == 2 {
		one = arrA[int(medianIndex(aMin, aMax, arrA))]
		two = arrB[int(medianIndex(bMin, bMax, arrB))]
		return (float64(two) - float64(one)) / float64(2)
	} else if (aMax-aMin+1) == 1{
		return median(aMin, aMax, arrA)
	} else {
		return median(bMin, bMax, arrB)
	}
}

//[] [0] [0 1] [0 1 2] [0 1 2 3 4] [0 1 2 3 4 5]
func medianIndex(min int, max int, arr[] int) float64{
	if len(arr) == 0 {
		panic("empty array")
	}
	return float64(min) + (float64(max-min)/2.0)
}

//[] [0] [0 1] [0 1 2] [0 1 2 3 4] [0 1 2 3 4 5]
func medianLess(min int, max int, arr []int) int {
	//note: [0] returns arr[0]. could be problem.
	if len(arr) == 0 {
		panic("empty array")
	} else if len(arr) < 3 {
		return arr[0]
	} else {
		index := int(medianIndex(min, max, arr))
		if (max - min +1 ) % 2 == 0 {
			return arr[index]
		} else {
			return arr[index-1]
		}
	}
}

//[] [0] [0 1] [0 1 2] [0 1 2 3 4] [0 1 2 3 4 5]
func medianMore(min int, max int, arr []int) int {
	if len(arr) == 0 {
		panic("empty array")
	} else if len(arr) == 1 {
		return arr[0]
	} else {
		index := int(medianIndex(min, max, arr))
		return arr[index+1]
	}
}

func median(min int, max int, arr []int) float64{
	index := int(medianIndex(min, max, arr))
	if (max-min+1) % 2 == 0 {
		return float64(arr[index])
	} else {
		return float64(arr[index] +  arr[index+1]) / 2.0
	}
}

func smallRange(aMax, aMin, bMax, bMin int) bool{
	return (aMax-aMin+1)+(bMax-bMin+1) < 3
}
