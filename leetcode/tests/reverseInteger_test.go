package tests

import (
	"github.com/himanshuo/golang-learn/leetcode"
	"testing"
)

var testReverseInteger = []struct {
	input int
	output int
}{  
	{1,1},
	{10,1},
	{100,1},
	{112,211},
	{-1,-1},
	{0,0},
	{-1231,-1321},
	{9909,9099},
	{12000,21},
	{-200,-2},
}

func TestReverseInteger(t *testing.T) {
	for i, test := range testReverseInteger {
		ret := leetcode.ReverseInteger(test.input)
		if ret != test.output {
			t.Errorf("#%d: %v : Incorrect result. Expected:%v Got:%v", i, test.input, test.output, ret)
		}
	}
}

