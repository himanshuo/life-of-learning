package tests

import (
	"github.com/himanshuo/golang-learn/leetcode"
	"testing"
)

var testLengthOfLongestSubstring = []struct {
	str  string
	num int
}{  //good
	{"hi", 2},
	{"aa", 1},
	{"abcabc", 3},
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{"", 0},
	{"abcdef", 6},
	{"abaaabd", 3},
	{"cdeef", 3},
	{"eeeeaeee", 2},
	{"abcdearqlmnop", 12},
	{"abcarlmn", 7},
	{"cabbba", 3},
	{"abcbb", 3},
	{"abbba", 2},

}

func TestLengthOfLongestSubstring(t *testing.T) {
	for i, test := range testLengthOfLongestSubstring {
		ret := leetcode.LengthOfLongestSubstring(test.str)
		if ret != test.num {
			t.Errorf("#%d: %v : Incorrect result. Expected:%v Got:%v", i, test.str, test.num, ret)
		}
	}
}

