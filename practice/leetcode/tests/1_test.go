package tests

import (
	"github.com/himanshuo/golang-learn/leetcode"
	"testing"
)

var testRadixSort = []struct {
	in  []int
	out []int
}{  //good
	{[]int{1,2,3}, []int{1,2,3}},
	{[]int{}, []int{}},
	{[]int{1,1,1}, []int{1,1,1}},
	{[]int{1}, []int{1}},
	{[]int{2,4,1}, []int{1,2,4}},
	{[]int{1,1,1}, []int{1,1,1}},
	{[]int{2,1,1}, []int{1,1,2}},
	{[]int{1,2,1}, []int{1,1,2}},
	{[]int{1,1,2}, []int{1,1,2}},
	{[]int{-1,-3,5,4}, []int{-3,-1, 4, 5}},
	{[]int{-91,2,3}, []int{-91,2,3}},
	{[]int{-91,0,0}, []int{-91,0,0}},
	{[]int{-91,0,0}, []int{-91,0,0}},
	{[]int{91,0,0}, []int{0,0, 91}},
	{[]int{91,1,1}, []int{1,1, 91}},
	{[]int{81, 91, 1}, []int{1,81, 91}},
	{[]int{-91,-100,-1}, []int{-100,-91,-1}},
	{[]int{-91,-100,-1,1,1,1,1,1}, []int{-100,-91,-1,1,1,1,1,1}},	
	
}


func TestRadixSort(t *testing.T) {
	for i, test := range testRadixSort {
		ret := leetcode.RadixSort(test.in)
		for k,_ := range test.out{
			if test.out[k] != ret[k]{
				t.Errorf("#%d: Input: %v Output: %v Expected Output: %v", i, test.in, ret, test.out)
				break
			}
		}
		//t.Errorf("correct: #%d: Input: %v Output: %v Expected Output: %v", i, test.in, ret, test.out)
		
	}
}

var testTwoSum = []struct{
	inList []int
	inTarget int
	out1 int
	out2 int
}{  
	{ []int{1,2,3}, 3, 0, 1 },
	{ []int{0,0}, 0, 0, 1 },
	{ []int{5,5,2}, 10, 0, 1 },        //2 5 5
	{ []int{1,4,0,4,2}, 8, 1,3},
	{ []int{0,0,-2,-4,-3}, 0, 0, 1 },
	{ []int{1,0,2,4,5}, 4, 1, 3 },
	{ []int{1,1,2}, 2,0,1 },
	{ []int{0,-3,-4, 1}, -4,0,2 },
	{ []int{-2,-4,4,3}, -1, 1,3 },
	{ []int{1,3,4}, 7,1,2 },
	{ []int{1, 2, -1}, 1, 1,2},
	{ []int{1, 2}, 3, 0,1},
	
	
	
	
}


func TestTwoSum(t *testing.T) {
	for i, test := range testTwoSum {
		ret1, ret2 := leetcode.TwoSum(test.inList, test.inTarget)
		
		if ret1 != test.out1 || ret2 != test.out2 {
			t.Errorf("#%d: Input: (%v,%v)  Output: (%v, %v) Expected Output: (%v,%v)", 
				i, 
				test.inList,test.inTarget,
				ret1, ret2,
				test.out1, test.out2,
			)
			break
		}
		
		//t.Errorf("correct: #%d: Input: %v Output: %v Expected Output: %v", i, test.in, ret, test.out)
		
	}
}
