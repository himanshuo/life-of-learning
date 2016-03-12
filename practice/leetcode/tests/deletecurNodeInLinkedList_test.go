package tests

import (
	"github.com/himanshuo/golang-learn/leetcode"
	"testing"
)

var a LinkedNode
var b LinkedNode
var c LinkedNode
var d LinkedNode
var e LinkedNode

var testDeleteCurNodeInLinkedList = []struct {
	list []LinkedNode
	toDelete LinkedNode
}{  
	{[]LinkedNode{
		LinkedNode{
			0,
			LinkedNode{1,nil}
			}
		},1},
}

func teardown(){
	
}
func TestReverseInteger(t *testing.T) {
	for i, test := range testReverseInteger {
		ret := leetcode.ReverseInteger(test.input)
		if ret != test.output {
			t.Errorf("#%d: %v : Incorrect result. Expected:%v Got:%v", i, test.input, test.output, ret)
		}
	}
}

