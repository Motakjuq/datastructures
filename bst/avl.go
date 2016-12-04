package bst

// Entry interface defines the comparation method, 1 if greate than entry, -1 if is lower and 0 if it's equal
type Entry interface {
	Compare(entry Entry) int
}

// avlNode contain the necessary structure for the tree creation
type avlNode struct {
	entry  Entry
	height int
	left   *avlNode
	right  *avlNode
}

// BST interface
type BST interface {
	Insert(entry Entry)
	Search(entry Entry) Entry
	Delete(entry Entry)
}

// NewAVL create a new empty AVL tree
func NewAVL() BST {
	return &avlTree{}
}

type avlTree struct {
	root *avlNode
}

// Insert an Entry into the tree, if the Entry already exists, it's override
func (avl *avlTree) Insert(entry Entry) {
	avl.root, _ = avl.root.insert(entry)
}

// Search func
func (avl *avlTree) Search(entry Entry) Entry {
	return avl.root.search(entry)
}

// Delete the specified Entry
func (avl *avlTree) Delete(entry Entry) {
	avl.root = avl.root.delete(entry)
}

//  _     _____  _    _   _      _____ _   _ _____ _      ______ _   _ _   _ _____
// | |   |  _  || |  | | | |    |  ___| | | |  ___| |     |  ___| | | | \ | /  __ \
// | |   | | | || |  | | | |    | |__ | | | | |__ | |     | |_  | | | |  \| | /  \/
// | |   | | | || |/\| | | |    |  __|| | | |  __|| |     |  _| | | | | . ` | |
// | |___\ \_/ /\  /\  / | |____| |___\ \_/ / |___| |____ | |   | |_| | |\  | \__/\
// \_____/\___/  \/  \/  \_____/\____/ \___/\____/\_____/ \_|    \___/\_| \_/\____/

func (n *avlNode) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *avlNode) updateHeight() {
	if n.left.getHeight() > n.right.getHeight() {
		n.height = n.left.getHeight() + 1
	} else {
		n.height = n.right.getHeight() + 1
	}
}

func (n *avlNode) balanceFactor() int {
	return -1*n.left.getHeight() + n.right.getHeight()
}

func (n *avlNode) leftRotation() *avlNode {
	parent := n.right
	n.right = parent.left
	parent.left = n

	n.updateHeight()
	parent.updateHeight()
	return parent
}

func (n *avlNode) rightRotation() *avlNode {
	parent := n.left
	n.left = parent.right
	parent.right = n

	n.updateHeight()
	parent.updateHeight()
	return parent
}

func (n *avlNode) leftRightRotation() *avlNode {
	n.left = n.left.leftRotation()
	return n.rightRotation()
}

func (n *avlNode) rightLeftRotation() *avlNode {
	n.right = n.right.rightRotation()
	return n.leftRotation()
}

func (n *avlNode) rebalance() (root *avlNode, updated bool) {
	height := n.height
	n.updateHeight()
	if height != n.height {
		updated = true
	}
	balance := n.balanceFactor()
	if balance == -2 { // heavy left
		if n.left.balanceFactor() > 0 {
			return n.leftRightRotation(), true
		}
		return n.rightRotation(), true
	}
	if balance == 2 {
		if n.right.balanceFactor() < 0 {
			return n.rightLeftRotation(), true
		}
		return n.leftRotation(), true
	}
	return n, updated
}

// insert func
func (n *avlNode) insert(data Entry) (*avlNode, bool) {
	if n == nil {
		return &avlNode{entry: data, height: 1}, true
	}
	var updated bool
	compareResult := data.Compare(n.entry)
	if compareResult < 0 { // left insert
		n.left, updated = n.left.insert(data)
	} else if compareResult > 0 { // right insert
		n.right, updated = n.right.insert(data)
	} else { // equals replace
		n.entry = data
		return n, updated
	}
	if updated {
		return n.rebalance()
	}
	return n, updated
}

// insert func
func (n *avlNode) insertIter(data Entry) *avlNode {
	if n == nil {
		return &avlNode{entry: data, height: 1}
	}

	var node *avlNode
	compResult := data.Compare(n.entry)
	if compResult < 0 { // left insert
		node = n.left
	} else if compResult > 0 { // right insert
		node = n.right
	} else { // equals replace
		n.entry = data
		return n
	}

	list := make([]*avlNode, 0, n.height)
	list = append(list, n)
	for node != nil {
		list = append(list, node)
		compResult = data.Compare(node.entry)
		if compResult < 0 { // left insert
			node = node.left
		} else if compResult > 0 { // right insert
			node = node.right
		} else { // equals replace
			node.entry = data
			return node
		}
		if node == nil {
			break
		}
	}

	if node == nil {
		node = &avlNode{entry: data, height: 1}
		compResult := data.Compare(list[len(list)-1].entry)
		if compResult < 0 { // left insert
			list[len(list)-1].left = node
		} else if compResult > 0 { // right insert
			list[len(list)-1].right = node
		}
	}

	//var updated bool
	size := len(list) - 2
	var updated bool
	for i := len(list) - 2; i >= 0; i-- {
		list[i], updated = list[i].rebalance()
		if i > 0 && updated {
			if list[i].entry.Compare(list[i-1].entry) < 0 {
				list[i-1].left = list[i]
			} else {
				list[i-1].right = list[i]
			}
		}
		if !updated && i < size {
			break
		}
	}
	return list[0]
}

// search returns the defined Entry or nil otherwise
func (n *avlNode) search(data Entry) Entry {
	if n != nil {
		node := n
		for node != nil {
			if node.entry.Compare(data) == 0 {
				return node.entry
			}
			if data.Compare(node.entry) < 0 {
				node = node.left
			} else {
				node = node.right
			}
		}
	}
	return nil
}

// return the new root of the tree and the removed node
func (n *avlNode) extractDeepLeftNode() (root, node *avlNode) {
	if n != nil {
		if n.left == nil {
			root := n.right
			n.right = nil
			return root, n
		}
		parent := n
		node = n.left
		list := make([]*avlNode, 0, n.height)
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
			list[i], _ = list[i].rebalance()
			if i >= 1 {
				list[i-1].left = list[i]
			}
		}
		return list[0], node
	}
	return nil, nil
}

// iterative delete
func (n *avlNode) delete(data Entry) *avlNode {
	if n != nil {
		list := make([]*avlNode, 0, n.height)
		node := n
		for node != nil {
			// if node.entry == data {
			if node.entry.Compare(data) == 0 {
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
					if node.entry.Compare(parent.entry) < 0 {
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
			if data.Compare(node.entry) < 0 {
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
				list[i], _ = list[i].rebalance()
				if i > 0 {
					if list[i].entry.Compare(list[i-1].entry) < 0 {
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
