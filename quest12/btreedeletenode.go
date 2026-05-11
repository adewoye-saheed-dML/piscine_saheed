package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	// Helper: find minimum node starting from 'n'
	minNode := func(n *TreeNode) *TreeNode {
		for n.Left != nil {
			n = n.Left
		}
		return n
	}

	var deleteNode func(curr, target *TreeNode) *TreeNode
	deleteNode = func(curr, target *TreeNode) *TreeNode {
		if curr == nil {
			return nil
		}

		if target.Data < curr.Data {
			curr.Left = deleteNode(curr.Left, target)
			if curr.Left != nil {
				curr.Left.Parent = curr
			}
		} else if target.Data > curr.Data {
			curr.Right = deleteNode(curr.Right, target)
			if curr.Right != nil {
				curr.Right.Parent = curr
			}
		} else {
			// Found node to delete

			// Case 1: No children
			if curr.Left == nil && curr.Right == nil {
				return nil
			}

			// Case 2: One child (right)
			if curr.Left == nil {
				curr.Right.Parent = curr.Parent
				return curr.Right
			}

			// Case 3: One child (left)
			if curr.Right == nil {
				curr.Left.Parent = curr.Parent
				return curr.Left
			}

			// Case 4: Two children
			// Find successor (min of right subtree)
			successor := minNode(curr.Right)
			curr.Data = successor.Data
			curr.Right = deleteNode(curr.Right, successor)
			if curr.Right != nil {
				curr.Right.Parent = curr
			}
		}

		return curr
	}

	newRoot := deleteNode(root, node)
	if newRoot != nil {
		newRoot.Parent = nil
	}
	return newRoot
}
