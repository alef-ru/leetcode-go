package problems

// First run:
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List II.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Reverse Linked List II.
// 😲
//
// Second run:
//Runtime: 4 ms, faster than 12.15% of Go online submissions for Reverse Linked List II.
//Memory Usage: 2.2 MB, less than 12.96% of Go online submissions for Reverse Linked List II.
//  -------------   submission start

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	var ln, bln, prev *ListNode = nil, nil, nil // left and one before left nodes
	cur := head
loop:
	for i := 1; ; i++ {
		switch {
		case i < left-1:
			cur = cur.Next
		case i == left-1:
			bln = cur
			cur = cur.Next
		case i == left:
			ln = cur
			prev, cur = cur, cur.Next
		case i > left && i < right:
			cur.Next, prev, cur = prev, cur, cur.Next
		case i == right:
			if bln == nil {
				head = cur
			} else {
				bln.Next = cur
			}
			ln.Next = cur.Next
			cur.Next = prev
			break loop
		}
	}
	return head
}

// --------------   submission end
