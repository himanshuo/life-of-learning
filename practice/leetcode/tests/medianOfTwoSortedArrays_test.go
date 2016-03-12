package tests

import (
	"github.com/himanshuo/golang-learn/leetcode"
	"testing"
)

var testMedianOfTwoSortedArrays = []struct {
	arrA []int
	arrB []int
	expected float64
}{  
	{[]int{0,1,2},[]int{0,1,2}, 1},
	{[]int{0,1,2},[]int{0,1,2}, 1},
	{[]int{0,1,2},[]int{0,1,2}, 1},
	{[]int{0,1,2},[]int{0,1,2}, 1},
	{[]int{0,1,2},[]int{0,1,2}, 1},
	{[]int{0,1,2},[]int{0,1,2}, 1},
	{[]int{0,1,2},[]int{0,1,2}, 1},
	
}

func TestMedianOfTwoSortedArrays(t *testing.T) {
	for i, test := range testMedianOfTwoSortedArrays {
		ret := leetcode.MedianOfTwoSortedArrays(test.arrA, test.arrB)
		if ret != test.expected {
			t.Errorf("#%d: (%v, %v) : Incorrect result. Expected:%v Got:%v", i, test.arrA,test.arrB, test.expected, ret)
		}
	}
}

