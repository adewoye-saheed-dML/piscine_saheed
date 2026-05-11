package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: nil}

	if l.Head == nil {
		// If list is empty, new node becomes both head and tail
		l.Head = newNode
		l.Tail = newNode
	} else {
		// Otherwise, add to end and update tail
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
