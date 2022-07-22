package problems

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Partition List.
//Memory Usage: 2.4 MB, less than 18.52% of Go online submissions for Partition List.

//  -------------   submission start

func partition(head *ListNode, x int) *ListNode {
	cur := head
	newHeadLeft, newHeadRight := &ListNode{}, &ListNode{}
	lastLeft, lastRight := newHeadLeft, newHeadRight
	for cur != nil {
		if cur.Val < x {
			lastLeft.Next = cur
			lastLeft = cur
		} else {
			lastRight.Next = cur
			lastRight = cur
		}
		cur = cur.Next
	}
	lastLeft.Next = newHeadRight.Next
	lastRight.Next = nil
	return newHeadLeft.Next
}

//  -------------   submission end
