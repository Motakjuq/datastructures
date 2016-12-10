package bst

const rbMaxReCheckBalance int8 = 2 // amount of times for re-checking the balance if something change in the tree

// rbNode contain the necessary structure for the tree creation
type rbNode struct {
	entry  Entry
	red    bool
	parent *rbNode
	left   *rbNode
	right  *rbNode
}

// NewRedBlack create a new empty AVL tree
func NewRedBlack() BST {
	return &redBlackTree{}
}

type redBlackTree struct {
	root *rbNode
}

// Insert an Entry into the tree, if the Entry already exists, it's override
func (rb *redBlackTree) Insert(entry Entry) {
	rb.root, _ = rb.root.insert(nil, nil, entry)
}

// Search func
func (rb *redBlackTree) Search(entry Entry) Entry {
	return rb.root.search(entry)
}

// Delete the specified Entry
func (rb *redBlackTree) Delete(entry Entry) {
	//avl.root = avl.root.delete(entry)
}

//  _     _____  _    _   _      _____ _   _ _____ _      ______ _   _ _   _ _____
// | |   |  _  || |  | | | |    |  ___| | | |  ___| |     |  ___| | | | \ | /  __ \
// | |   | | | || |  | | | |    | |__ | | | | |__ | |     | |_  | | | |  \| | /  \/
// | |   | | | || |/\| | | |    |  __|| | | |  __|| |     |  _| | | | | . ` | |
// | |___\ \_/ /\  /\  / | |____| |___\ \_/ / |___| |____ | |   | |_| | |\  | \__/\
// \_____/\___/  \/  \/  \_____/\____/ \___/\____/\_____/ \_|    \___/\_| \_/\____/

func (rb *rbNode) leftRotation() *rbNode {
	parent := rb.right
	rb.right = parent.left
	parent.left = rb

	return parent
}

func (rb *rbNode) rightRotation() *rbNode {
	parent := rb.left
	rb.left = parent.right
	parent.right = rb

	return parent
}

func (rb *rbNode) leftRightRotation() *rbNode {
	rb.left = rb.left.leftRotation()
	return rb.rightRotation()
}

func (rb *rbNode) rightLeftRotation() *rbNode {
	rb.right = rb.right.rightRotation()
	return rb.leftRotation()
}

func (rb *rbNode) isBlack() bool {
	if rb == nil || !rb.red {
		return true
	}
	return false
}

func (rb *rbNode) isRed() bool {
	return !rb.isBlack()
}

// insert func
func (rb *rbNode) insert(grandParent *rbNode, parent *rbNode, data Entry) (*rbNode, int8) { // root, inserted
	if rb == nil {
		if parent == nil {
			return &rbNode{entry: data}, rbMaxReCheckBalance
		}
		return &rbNode{entry: data, red: true}, rbMaxReCheckBalance
	}

	compareResult := data.Compare(rb.entry)
	var reCheckTimes int8 = -1
	if compareResult < 0 { // left insert
		rb.left, reCheckTimes = rb.left.insert(parent, rb, data)
	} else if compareResult > 0 { // right insert
		rb.right, reCheckTimes = rb.right.insert(parent, rb, data)
	} else { // equals replace
		rb.entry = data
		return rb, -1
	}

	if reCheckTimes >= 0 {
		var treeChanged bool
		rb, treeChanged = rb.rebalance(grandParent, parent)
		if treeChanged {
			return rb, rbMaxReCheckBalance
		}
	}

	return rb, reCheckTimes - 1
}

func (rb *rbNode) rebalance(grandParent, parent *rbNode) (*rbNode, bool) {
	if !rb.isBlack() && (!rb.left.isBlack() || !rb.right.isBlack()) && !parent.left.isBlack() && !parent.right.isBlack() {
		if grandParent != nil {
			parent.red = true
		} else {
			parent.red = false
		}
		parent.left.red = false
		parent.right.red = false
		return rb, true
	}
	var rotated bool
	if rb.left != nil && rb.right.isBlack() && !rb.left.isBlack() {
		if rb.left.left != nil && !rb.left.left.isBlack() {
			rotated = true
			rb = rb.rightRotation()
		} else if rb.left.right != nil && !rb.left.right.isBlack() {
			rotated = true
			rb = rb.leftRightRotation()
		}
	} else if rb.right != nil && rb.left.isBlack() && !rb.right.isBlack() {
		if rb.right.right != nil && !rb.right.right.isBlack() {
			rotated = true
			rb = rb.leftRotation()
		} else if rb.right.left != nil && !rb.right.left.isBlack() {
			rotated = true
			rb = rb.rightLeftRotation()
		}
	}
	// if rb.left != nil && !rb.left.isBlack() && rb.left.left != nil && !rb.left.left.isBlack() && rb.right.isBlack() {
	// 	rotated = true
	// 	rb = rb.rightRotation()
	// } else if rb.left != nil && !rb.left.isBlack() && rb.left.right != nil && !rb.left.right.isBlack() && rb.right.isBlack() {
	// 	rotated = true
	// 	rb = rb.leftRightRotation()
	// } else if rb.right != nil && !rb.right.isBlack() && rb.right.right != nil && !rb.right.right.isBlack() && rb.left.isBlack() {
	// 	rotated = true
	// 	rb = rb.leftRotation()
	// } else if rb.right != nil && !rb.right.isBlack() && rb.right.left != nil && !rb.right.left.isBlack() && rb.left.isBlack() {
	// 	rotated = true
	// 	rb = rb.rightLeftRotation()
	// }

	if rotated {
		rb.red = false
		rb.left.red = true
		rb.right.red = true
		return rb, true
	}

	return rb, false
}

// search returns the defined Entry or nil otherwise
func (rb *rbNode) search(data Entry) Entry {
	if rb != nil {
		node := rb
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
