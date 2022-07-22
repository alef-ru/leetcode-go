package problems

import "fmt"

// Types that are used in more than one exercises

type ListNode struct {
	Val  int
	Next *ListNode
}

type intSlice []int

func (slice intSlice) asList() *ListNode {
	var cur, head *ListNode
	for _, v := range slice {
		prev := cur
		cur = &ListNode{Val: v}
		if head == nil {
			head = cur
		} else {
			prev.Next = cur
		}
	}
	return head
}

func (head ListNode) asSlice() intSlice {
	result := make([]int, 0, 10)
	cur := &head
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}
	return result
}

func (head ListNode) String() string {
	return fmt.Sprint(head.asSlice())
}
