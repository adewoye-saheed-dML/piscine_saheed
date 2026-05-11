package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

// helper: merge two sorted lists
func merge(l1, l2 *NodeI) *NodeI {
	dummy := &NodeI{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Data <= l2.Data {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}

	return dummy.Next
}

// helper: split the list into two halves
func split(head *NodeI) (*NodeI, *NodeI) {
	if head == nil || head.Next == nil {
		return head, nil
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow.Next
	slow.Next = nil

	return head, mid
}

// ListSort sorts a linked list of NodeI in ascending order
func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	left, right := split(l)
	left = ListSort(left)
	right = ListSort(right)

	return merge(left, right)
}
