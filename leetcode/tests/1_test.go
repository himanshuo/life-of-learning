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
	{[]int{-1,-3,5,4}, []int{-1,-3, 4, 5}},
	{[]int{-91,2,3}, []int{-91,2,3}},
	{[]int{-91,0,0}, []int{-91,0,0}},
	{[]int{-91,0,0}, []int{-91,0,0}},
	{[]int{-91,-100,-1}, []int{-100,-91,-1}},	
	
	
}


func TestRadixSort(t *testing.T) {
	for i, test := range testRadixSort {
		ret := leetcode.RadixSort(test.in)
		for k,_ :=range test.out{
			if test.out[k] != ret[k]{
				t.Errorf("#%d: Input: %v Output: %v Expected Output: %v", i, test.in, ret, test.out)
			}
		}
		
	}
}

