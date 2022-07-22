package problems

// Types that used in more than one excercise

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(l []int) *ListNode {
	var cur, head *ListNode
	for _, v := range l {
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

func list2slice(head *ListNode) []int {
	result := make([]int, 0, 10)
	cur := head
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}
	return result
}
