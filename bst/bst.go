package bst

// Entry interface defines the comparation method, 1 if greate than entry, -1 if is lower and 0 if it's equal
type Entry interface {
	Compare(entry Entry) int
}

// BST interface
type BST interface {
	Insert(entry Entry)
	Search(entry Entry) Entry
	Delete(entry Entry)
	// PullFirst() Entry
	// PullLast() Entry
}