package piscine

func ListReverse(l *List) {
	var prev *NodeL = nil
	current := l.Head
	l.Tail = l.Head // The old head becomes the new tail

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev // The last processed node becomes the new head
}
