package cmap

import (
	"fmt"
	"slices"
)

func NewLNode(key []byte, value []byte) *node {
	keyCopy := slices.Clone(key)
	valueCopy := slices.Clone(value)
	return &node{
		isLeaf: true,
		key:    keyCopy,
		value:  valueCopy,
	}
}

func NewINode() *node {
	return &node{
		isLeaf:   false,
		bitmap:   0,
		children: []*node{},
	}
}

func (n *node) copyNode() *node {
	childrenCopy := make([]*node, len(n.children))
	copy(childrenCopy, n.children)
	key := slices.Clone(n.Key())
	value := slices.Clone(n.Value())
	return &node{
		key:      key,
		value:    value,
		isLeaf:   n.IsLeaf(),
		bitmap:   n.Bitmap(),
		children: childrenCopy,
	}
}

func (n *node) extendTable(bitMap uint32, pos int, newNode *node) {
	newTable := make([]*node, calculateHammingWeight(bitMap))
	copy(newTable[:pos], n.children[:pos])
	newTable[pos] = newNode
	copy(newTable[pos+1:], n.children[pos:])
	n.children = newTable
}

func (n *node) shrinkTable(bitMap uint32, pos int) {
	newTable := make([]*node, calculateHammingWeight(bitMap))
	copy(newTable[:pos], n.children[:pos])
	copy(newTable[pos:], n.children[pos+1:])
	n.children = newTable
}

func (n *node) Key() []byte {
	return n.key
}

func (n *node) Value() []byte {
	return n.value
}

func (n *node) IsLeaf() bool {
	return n.isLeaf
}

func (n *node) Bitmap() uint32 {
	return n.bitmap
}

func (n *node) setBitmap(bitmap uint32) {
	n.bitmap = bitmap
}

func (n *node) Children() []Node {
	var children []Node
	for _, child := range n.children {
		children = append(children, child)
	}

	return children
}

func (n *node) Child(pos int) Node {
	return n.children[pos]
}

func (n *node) setChild(child *node, pos int) {
	n.children[pos] = child
}

func (n *node) PrintChildren() {
	if n == nil {
		return
	}
	n.printChildrenRecursive(0)
}

func (n *node) printChildrenRecursive(level int) {
	for idx, child := range n.children {
		if child != nil {
			fmt.Printf("Level: %d, Index: %d, Key: %s, Value: %v\n", level, idx, child.Key(), child.Value())
			child.printChildrenRecursive(level + 1)
		}
	}
}
