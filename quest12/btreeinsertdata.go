package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	newNode := &TreeNode{Data: data}

	// If the tree is empty, the new node becomes the root
	if root == nil {
		return newNode
	}

	current := root
	for {
		if data < current.Data {
			if current.Left == nil {
				current.Left = newNode
				newNode.Parent = current
				break
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = newNode
				newNode.Parent = current
				break
			}
			current = current.Right
		}
	}
	return root
}
