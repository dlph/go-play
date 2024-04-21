package partitionlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.com/problems/partition-list
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == nil {
		return head // nothing to swap with
	}

	if head.Val >= x {
		tail := partition(head.Next, x)

		if tail == nil {
			return head
		}

		//  don't swap
		if tail.Val >= x {
			head.Next = tail
			return head
		}

		// swap
		head.Next = tail.Next
		tail.Next = partition(head, x)

		return tail
	}

	tail := partition(head.Next, x)
	head.Next = tail

	return head
}
