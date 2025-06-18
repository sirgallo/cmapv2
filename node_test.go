package cmap

import (
	"bytes"
	"testing"
)

func TestNode(t *testing.T) {
	t.Run("test node copy isolation", func(t *testing.T) {
		child1 := &node{isLeaf: true}
		child2 := &node{isLeaf: true}
		parent := &node{children: []*node{child1, child2}}
		clone := parent.copyNode()
		clone.children[0] = nil // mutate clone child

		if parent.children[0] == nil { // the original must not see that mutation
			t.Fatal("parent.children was mutated when clone.children was changed")
		}
	})

	t.Run("test node extend and shrink table", func(t *testing.T) {
		n := &node{children: []*node{}}

		leafA := NewLNode([]byte("a"), []byte("A"))
		n.extendTable(1<<0, 0, leafA)
		if len(n.children) != 1 || n.children[0] != leafA {
			t.Errorf("extendTable failed, got %v", n.children)
		}

		leafB := NewLNode([]byte("b"), []byte("B"))
		n.extendTable((1<<0)|(1<<1), 1, leafB) // set bits for positions 0 and 1 in bitmap = 0b11
		if len(n.children) != 2 || n.children[1] != leafB {
			t.Errorf("extendTable 2nd element failed, got %v", n.children)
		}

		n.shrinkTable(1<<1, 0) // clear bit 0, leaving only bit1 index 0 in new slice
		if len(n.children) != 1 || n.children[0] != leafB {
			t.Errorf("shrinkTable failed, got %v", n.children)
		}
	})

	t.Run("test key value immutability", func(t *testing.T) {
		m := NewMap()
		key := []byte("foo")
		val := []byte("bar")
		m.Put(key, val)

		key[0], val[0] = 'F', 'B' // corrupt the caller’s slices

		out := m.Get([]byte("foo"))
		if !bytes.Equal(out, []byte("bar")) {
			t.Fatalf("expected stored value to be ‘bar’, got %q", out)
		}
	})
}
