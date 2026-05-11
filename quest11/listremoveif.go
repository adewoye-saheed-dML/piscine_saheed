package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil || l.Head == nil {
		return
	}

	// Remove nodes from the beginning that match data_ref
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	// If the list is now empty, update Tail as well
	if l.Head == nil {
		l.Tail = nil
		return
	}

	// Remove nodes in the middle
	prev := l.Head
	curr := l.Head.Next

	for curr != nil {
		if curr.Data == data_ref {
			prev.Next = curr.Next
			if curr == l.Tail {
				l.Tail = prev
			}
			curr = prev.Next
		} else {
			prev = curr
			curr = curr.Next
		}
	}
}
