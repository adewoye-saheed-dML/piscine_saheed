package piscine

// ListLast returns the data of the last node in the linked list.
func ListLast(l *List) interface{} {
	// If the list is empty, return nil
	if l.Head == nil {
		return nil
	}

	// Start from the head and traverse the list
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	// Return the data of the last node
	return current.Data
}
