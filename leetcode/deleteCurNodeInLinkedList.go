package leetcode


func DeleteCurNodeInLinkedList(cur ListNode) {
	current := cur
	for ;current.next.next != nul ; {
		current.val = current.next.val
		current = current.next
	}
}
