package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}

	// Case 1: Empty list or new node should be new head
	if l == nil || data_ref <= l.Data {
		newNode.Next = l
		return newNode
	}

	// Case 2: Insert somewhere after head
	current := l
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode

	return l
}
