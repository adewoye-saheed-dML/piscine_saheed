package piscine

// ListPushFront function as before
func ListPushFront(l *List, data interface{}) {
	// Create a new node with the provided data
	newNode := &NodeL{
		Data: data,
		Next: l.Head,
	}

	// Update the head of the list to the new node
	l.Head = newNode

	// If the list was empty, update the Tail pointer to the new node
	if l.Tail == nil {
		l.Tail = newNode
	}
}
