package bst

import (
	"fmt"
	"math/rand"
	"testing"
)

func (root *rbNode) inOrder(h int) {
	if root == nil {
		return
	}

	te := root.entry.(*testEntry)

	root.left.inOrder(h + 1)
	fmt.Printf("%d: %d, black: %t\n", h, te.value, root.isBlack())
	root.right.inOrder(h + 1)
}

func (root *rbNode) validate(t *testing.T, parent *rbNode, currentBlackNodes int, expectedBlackNodes int) {
	if root == nil {
		if currentBlackNodes != expectedBlackNodes {
			t.Errorf("Black depth path doesn't match")
		}
		return
	}

	te := root.entry.(*testEntry)
	if parent == nil && !root.isBlack() {
		t.Errorf("Parent '%d' doesn't black", te.value)
	}
	if !root.isBlack() && (!root.left.isBlack() || !root.right.isBlack()) {
		t.Errorf("Red node '%d' doesn't have black children", te.value)
	}

	var countBlack int
	if root.isBlack() {
		countBlack++
	}

	root.left.validate(t, root, currentBlackNodes+countBlack, expectedBlackNodes)
	root.right.validate(t, root, currentBlackNodes+countBlack, expectedBlackNodes)
}

func (root *rbNode) countBlackNodes() int {
	if root == nil {
		return 0
	}
	if root.isBlack() {
		return root.left.countBlackNodes() + 1
	}
	return root.left.countBlackNodes()
}

func Test_rbNode_insert(t *testing.T) {
	var root *rbNode
	nodes := 300
	rnd := rand.New(rand.NewSource(int64(1000)))

	for i := 0; i < nodes; i++ {
		root, _ = root.insert(nil, nil, newTestEntry(rnd.Int()))
	}
	root.validate(t, nil, 0, root.countBlackNodes())
}

func BenchmarkRedBlackTreeInsertEmptyTree(b *testing.B) {
	var root *rbNode

	rnd := rand.New(rand.NewSource(int64(b.N)))
	list := make([]*testEntry, b.N)

	for i := 0; i < b.N; i++ {
		list[i] = newTestEntry(rnd.Int())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		root, _ = root.insert(nil, nil, list[i])
	}
}

func BenchmarkRedBlackTreeInsert(b *testing.B) {
	var root *rbNode

	rnd := rand.New(rand.NewSource(int64(b.N)))
	list := make([]*testEntry, b.N)
	for i := 0; i < b.N; i++ {
		root, _ = root.insert(nil, nil, newTestEntry(rnd.Int()))
	}

	for i := 0; i < b.N; i++ {
		list[i] = newTestEntry(rnd.Int())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		root, _ = root.insert(nil, nil, list[i])
	}
}
