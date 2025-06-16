package cmap

import (
	"fmt"
)

func (cMap *cMap) NewLNode(key []byte, value []byte) *node {
	n := cMap.pool.GetNode()
	n.setIsLeaf(true)
	n.setKey(key)
	n.setValue(value)
	return n
}

func (cMap *cMap) NewINode() *node {
	n := cMap.pool.GetNode()
	n.setIsLeaf(false)
	n.setBitmap(0)
	n.setChildren([]*node{})
	return n
}

func (cMap *cMap) CopyNode(n *node) *node {
	nodeCopy := cMap.pool.GetNode()
	nodeCopy.setKey(n.Key())
	nodeCopy.setValue(n.Value())
	nodeCopy.setIsLeaf(n.IsLeaf())
	nodeCopy.setBitmap(n.Bitmap())
	nodeCopy.setChildren(make([]*node, len(n.children)))

	copy(nodeCopy.children, n.children)
	return nodeCopy
}

func (cMap *cMap) ExtendTable(orig []*node, bitMap uint32, pos int, newNode *node) []*node {
	tableSize := calculateHammingWeight(bitMap)
	newTable := make([]*node, tableSize)

	copy(newTable[:pos], orig[:pos])
	newTable[pos] = newNode
	copy(newTable[pos+1:], orig[pos:])
	return newTable
}

func (cMap *cMap) ShrinkTable(orig []*node, bitMap uint32, pos int) []*node {
	tableSize := calculateHammingWeight(bitMap)
	newTable := make([]*node, tableSize)

	copy(newTable[:pos], orig[:pos])
	copy(newTable[pos:], orig[pos+1:])
	return newTable
}

func (n *node) Key() []byte {
	return n.key
}

func (n *node) setKey(key []byte) {
	n.key = key
}

func (n *node) Value() []byte {
	return n.value
}

func (n *node) setValue(value []byte) {
	n.value = value
}

func (n *node) IsLeaf() bool {
	return n.isLeaf
}

func (n *node) setIsLeaf(isLeaf bool) {
	n.isLeaf = isLeaf
}

func (n *node) Bitmap() uint32 {
	return n.bitmap
}

func (n *node) setBitmap(bitmap uint32) {
	n.bitmap = bitmap
}

func (n *node) Children() []*node {
	return n.children
}

func (n *node) setChildren(children []*node) {
	n.children = children
}

func (n *node) Child(pos int) *node {
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
