package piscine

func ListMerge(l1 *List, l2 *List) {
	if l2 == nil || l2.Head == nil {
		return // nothing to merge
	}
	if l1 == nil {
		return // destination list is nil
	}
	if l1.Head == nil {
		// l1 is empty, so just point l1 to l2
		l1.Head = l2.Head
		l1.Tail = l2.Tail
	} else {
		// l1 is not empty: attach l2 at the end
		l1.Tail.Next = l2.Head
		l1.Tail = l2.Tail
	}
}
