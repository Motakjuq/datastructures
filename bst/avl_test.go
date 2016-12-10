package bst

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestNewAVL(t *testing.T) {
	tests := []struct {
		name string
		want BST
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := NewAVL(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewAVL() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_avlTree_Insert(t *testing.T) {
	type args struct {
		entry Entry
	}
	tests := []struct {
		name string
		avl  *avlTree
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.avl.Insert(tt.args.entry)
	}
}

func Test_avlTree_Search(t *testing.T) {
	type args struct {
		entry Entry
	}
	tests := []struct {
		name string
		avl  *avlTree
		args args
		want Entry
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.avl.Search(tt.args.entry); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. avlTree.Search() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_avlTree_Delete(t *testing.T) {
	type args struct {
		entry Entry
	}
	tests := []struct {
		name string
		avl  *avlTree
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.avl.Delete(tt.args.entry)
	}
}

func Test_avlNode_getHeight(t *testing.T) {
	tests := []struct {
		name string
		n    *avlNode
		want int
	}{
		{
			name: "Nil node height",
			n:    nil,
			want: 0,
		},
		{
			name: "Simple node height",
			n:    &avlNode{height: 13},
			want: 13,
		},
	}
	for _, tt := range tests {
		if got := tt.n.getHeight(); got != tt.want {
			t.Errorf("%q. avlNode.getHeight() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_avlNode_updateHeight(t *testing.T) {
	tests := []struct {
		name   string
		n      *avlNode
		result int
	}{
		{
			name:   "Simple node",
			n:      &avlNode{height: 13},
			result: 1,
		},
		{
			name:   "Simple node, right child",
			n:      &avlNode{height: 13, right: &avlNode{height: 1}},
			result: 2,
		},
		{
			name:   "Simple node, left child",
			n:      &avlNode{height: 13, left: &avlNode{height: 1}},
			result: 2,
		},
		{
			name:   "Simple node, two children",
			n:      &avlNode{height: 13, left: &avlNode{height: 1}, right: &avlNode{height: 1}},
			result: 2,
		},
		{
			name:   "Simple node, heavy right",
			n:      &avlNode{height: 13, left: &avlNode{height: 1}, right: &avlNode{height: 2}},
			result: 3,
		},
		{
			name:   "Simple node, heavy left",
			n:      &avlNode{height: 13, left: &avlNode{height: 2}, right: &avlNode{height: 1}},
			result: 3,
		},
	}
	for _, tt := range tests {
		tt.n.updateHeight()
		if tt.n.height != tt.result {
			t.Errorf("%q. avlNode.getHeight() = %v, want %v", tt.name, tt.n.height, tt.result)
		}
	}
}

func Test_avlNode_balanceFactor(t *testing.T) {
	tests := []struct {
		name   string
		n      *avlNode
		result int
	}{
		{
			name:   "Simple node",
			n:      &avlNode{height: 13},
			result: 0,
		},
		{
			name:   "Simple node, right child",
			n:      &avlNode{height: 13, right: &avlNode{height: 1}},
			result: 1,
		},
		{
			name:   "Simple node, left child",
			n:      &avlNode{height: 13, left: &avlNode{height: 1}},
			result: -1,
		},
		{
			name:   "Simple node, two children",
			n:      &avlNode{height: 13, left: &avlNode{height: 1}, right: &avlNode{height: 1}},
			result: 0,
		},
		{
			name:   "Simple node, heavy right",
			n:      &avlNode{height: 13, left: &avlNode{height: 1}, right: &avlNode{height: 2}},
			result: 1,
		},
		{
			name:   "Simple node, heavy left",
			n:      &avlNode{height: 13, left: &avlNode{height: 2}, right: &avlNode{height: 1}},
			result: -1,
		},
		{
			name:   "Simple node, unbalance right",
			n:      &avlNode{height: 13, left: &avlNode{height: 1}, right: &avlNode{height: 3}},
			result: 2,
		},
		{
			name:   "Simple node, unbalance left",
			n:      &avlNode{height: 13, left: &avlNode{height: 3}, right: &avlNode{height: 1}},
			result: -2,
		},
	}
	for _, tt := range tests {
		if got := tt.n.balanceFactor(); got != tt.result {
			t.Errorf("%q. avlNode.balanceFactor() = %v, want %v", tt.name, got, tt.result)
		}
	}
}

func Test_avlNode_leftRotation(t *testing.T) {
	tests := []struct {
		name string
		n    *avlNode
		want *avlNode
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.n.leftRotation(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. avlNode.leftRotation() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_avlNode_rightRotation(t *testing.T) {
	tests := []struct {
		name string
		n    *avlNode
		want *avlNode
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.n.rightRotation(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. avlNode.rightRotation() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_avlNode_leftRightRotation(t *testing.T) {
	tests := []struct {
		name string
		n    *avlNode
		want *avlNode
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.n.leftRightRotation(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. avlNode.leftRightRotation() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_avlNode_rightLeftRotation(t *testing.T) {
	tests := []struct {
		name string
		n    *avlNode
		want *avlNode
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.n.rightLeftRotation(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. avlNode.rightLeftRotation() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_avlNode_rebalance(t *testing.T) {
	tests := []struct {
		name        string
		n           *avlNode
		wantRoot    *avlNode
		wantUpdated bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		gotRoot, gotUpdated := tt.n.rebalance()
		if !reflect.DeepEqual(gotRoot, tt.wantRoot) {
			t.Errorf("%q. avlNode.rebalance() gotRoot = %v, want %v", tt.name, gotRoot, tt.wantRoot)
		}
		if gotUpdated != tt.wantUpdated {
			t.Errorf("%q. avlNode.rebalance() gotUpdated = %v, want %v", tt.name, gotUpdated, tt.wantUpdated)
		}
	}
}

func Test_avlNode_insert(t *testing.T) {
	type args struct {
		data Entry
	}
	tests := []struct {
		name  string
		n     *avlNode
		args  args
		want  *avlNode
		want1 bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, got1 := tt.n.insert(tt.args.data)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. avlNode.insert() got = %v, want %v", tt.name, got, tt.want)
		}
		if got1 != tt.want1 {
			t.Errorf("%q. avlNode.insert() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func Test_avlNode_insertIter(t *testing.T) {
	type args struct {
		data Entry
	}
	tests := []struct {
		name string
		n    *avlNode
		args args
		want *avlNode
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.n.insertIter(tt.args.data); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. avlNode.insertIter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_avlNode_search(t *testing.T) {
	type args struct {
		data Entry
	}
	tests := []struct {
		name string
		n    *avlNode
		args args
		want Entry
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.n.search(tt.args.data); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. avlNode.search() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_avlNode_extractDeepLeftNode(t *testing.T) {
	tests := []struct {
		name     string
		n        *avlNode
		wantRoot *avlNode
		wantNode *avlNode
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		gotRoot, gotNode := tt.n.extractDeepLeftNode()
		if !reflect.DeepEqual(gotRoot, tt.wantRoot) {
			t.Errorf("%q. avlNode.extractDeepLeftNode() gotRoot = %v, want %v", tt.name, gotRoot, tt.wantRoot)
		}
		if !reflect.DeepEqual(gotNode, tt.wantNode) {
			t.Errorf("%q. avlNode.extractDeepLeftNode() gotNode = %v, want %v", tt.name, gotNode, tt.wantNode)
		}
	}
}

func Test_avlNode_delete(t *testing.T) {
	type args struct {
		data Entry
	}
	tests := []struct {
		name string
		n    *avlNode
		args args
		want *avlNode
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.n.delete(tt.args.data); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. avlNode.delete() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func BenchmarkAVLTreeInsertEmptyTree(b *testing.B) {
	var root *avlNode

	rnd := rand.New(rand.NewSource(int64(b.N)))
	list := make([]*testEntry, b.N)

	for i := 0; i < b.N; i++ {
		list[i] = newTestEntry(rnd.Int())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		root, _ = root.insert(list[i])
	}
}

func BenchmarkAVLTreeInsert(b *testing.B) {
	var root *avlNode

	rnd := rand.New(rand.NewSource(int64(b.N)))
	list := make([]*testEntry, b.N)
	for i := 0; i < b.N; i++ {
		root, _ = root.insert(newTestEntry(rnd.Int()))
	}

	for i := 0; i < b.N; i++ {
		list[i] = newTestEntry(rnd.Int())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		root, _ = root.insert(list[i])
	}
}

func BenchmarkAVLTreeInsertIterEmptyTree(b *testing.B) {
	var root *avlNode

	rnd := rand.New(rand.NewSource(int64(b.N)))
	list := make([]*testEntry, b.N)

	for i := 0; i < b.N; i++ {
		list[i] = newTestEntry(rnd.Int())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		root = root.insertIter(list[i])
	}
}

func BenchmarkAVLTreeInsertIter(b *testing.B) {
	var root *avlNode

	rnd := rand.New(rand.NewSource(int64(b.N)))
	list := make([]*testEntry, b.N)
	for i := 0; i < b.N; i++ {
		root = root.insertIter(newTestEntry(rnd.Int()))
	}

	for i := 0; i < b.N; i++ {
		list[i] = newTestEntry(rnd.Int())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		root = root.insertIter(list[i])
	}
}
