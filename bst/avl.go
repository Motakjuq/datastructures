package bst

// AvlNode contain the necessary structure for the tree creation
type AvlNode struct {
	value  int
	height int
	left   *AvlNode
	right  *AvlNode
}

func (n *AvlNode) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *AvlNode) updateHeight() {
	if n.left.getHeight() > n.right.getHeight() {
		n.height = n.left.getHeight() + 1
	} else {
		n.height = n.right.getHeight() + 1
	}
}

func (n *AvlNode) balanceFactor() int {
	return -1*n.left.getHeight() + n.right.getHeight()
}

func (n *AvlNode) leftRotation() *AvlNode {
	parent := n.right
	n.right = parent.left
	parent.left = n

	n.updateHeight()
	parent.updateHeight()
	return parent
}

func (n *AvlNode) rightRotation() *AvlNode {
	parent := n.left
	n.left = parent.right
	parent.right = n

	n.updateHeight()
	parent.updateHeight()
	return parent
}

func (n *AvlNode) leftRightRotation() *AvlNode {
	n.left = n.left.leftRotation()
	parent := n.rightRotation()
	return parent
}

func (n *AvlNode) rightLeftRotation() *AvlNode {
	n.right = n.right.rightRotation()
	parent := n.leftRotation()
	return parent
}

func (n *AvlNode) rebalance() *AvlNode {
	n.updateHeight()
	balance := n.balanceFactor()
	if balance == -2 { // heavy left
		if n.left.balanceFactor() > 0 {
			n = n.leftRightRotation()
		} else {
			n = n.rightRotation()
		}
	}
	if balance == 2 {
		if n.right.balanceFactor() < 0 {
			n = n.rightLeftRotation()
		} else {
			n = n.leftRotation()
		}
	}
	return n
}

// Insert func
func (n *AvlNode) Insert(data int) *AvlNode {
	if n == nil {
		return &AvlNode{value: data, height: 1}
	}

	if data < n.value { // left insert
		n.left = n.left.Insert(data)
	} else if n.value < data { // right insert
		n.right = n.right.Insert(data)
	} else { // equals replace
		n.value = data
		return n
	}

	return n.rebalance()
}

// Contains func
func (n *AvlNode) Contains(data int) bool {
	if n != nil {
		node := n
		for node != nil {
			if node.value == data {
				// break
				return true
			}
			if data < node.value {
				node = node.left
			} else {
				node = node.right
			}
		}
	}
	return false
}

// return the new root of the tree and the removed node
func (n *AvlNode) extractDeepLeftNode() (root, node *AvlNode) {
	if n != nil {
		if n.left == nil {
			root := n.right
			n.right = nil
			return root, n
		}
		parent := n
		node = n.left
		list := make([]*AvlNode, 0, n.height)
		list = append(list, parent)
		i := 1
		for node.left != nil {
			i++
			parent = node
			node = node.left
			list = append(list, parent)
		}

		parent.left = node.right
		node.right = nil
		for i > 0 {
			i--
			list[i] = list[i].rebalance()
			if i >= 1 {
				list[i-1].left = list[i]
			}
		}
		return list[0], node
	}
	return nil, nil
}

// Delete func
func (n *AvlNode) Delete(data int) *AvlNode {
	if n != nil {
		list := make([]*AvlNode, 0, n.height)
		node := n
		for node != nil {
			if node.value == data {
				if node.left != nil && node.right != nil { // find the replace for the root node
					rightNode, replace := node.right.extractDeepLeftNode()
					replace.left = node.left
					replace.right = rightNode
					list = append(list, replace)
					node.left = nil
					node.right = nil
				} else if node.left == nil && node.right == nil { // is leaf
					if len(list) == 0 {
						return nil
					}
					parent := list[len(list)-1]
					if node.value < parent.value {
						parent.left = nil
					} else {
						parent.right = nil
					}
				} else if node.left != nil && node.right == nil { // only has left children
					list = append(list, node.left)
					node.left = nil
				} else { // only has right children
					list = append(list, node.right)
					node.right = nil
				}
				break
			}

			list = append(list, node)
			if data < node.value {
				node = node.left
			} else {
				node = node.right
			}
		}
		// it doesn't have the node
		if node == nil {
			return n
		}

		if len(list) > 0 {
			for i := len(list) - 1; i >= 0; i-- {
				list[i] = list[i].rebalance()
				if i > 0 {
					if list[i].value < list[i-1].value {
						list[i-1].left = list[i]
					} else {
						list[i-1].right = list[i]
					}
				}
			}
			return list[0]
		}
	}
	return nil
}
