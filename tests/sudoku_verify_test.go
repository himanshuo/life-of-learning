package tests

import (
	"github.com/himanshuo/golang-learn/main"
	"testing"
)

func TestVerify(t *testing.T) {
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

