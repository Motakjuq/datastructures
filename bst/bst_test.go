package bst

type testEntry struct {
	value int
}

func (t *testEntry) Compare(entry Entry) int {
	e2 := entry.(*testEntry)
	return t.value - e2.value
}

func newTestEntry(value int) *testEntry {
	return &testEntry{value}
}
