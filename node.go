package cmap

import (
	"fmt"
)

func (cMap *cMap) NewLNode(key []byte, value []byte) *node {
	n := cMap.pool.GetNode()
	n.isLeaf = true
	n.key = key
	n.value = value
	return n
}

func (cMap *cMap) NewINode() *node {
	n := cMap.pool.GetNode()
	n.isLeaf = false
	n.bitmap = 0
	n.children = []*node{}
	return n
}

func (cMap *cMap) CopyNode(n *node) *node {
	nodeCopy := cMap.pool.GetNode()
	nodeCopy.key = n.key
	nodeCopy.value = n.value
	nodeCopy.isLeaf = n.isLeaf
	nodeCopy.bitmap = n.bitmap
	nodeCopy.children = make([]*node, len(n.children))

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

func (n *node) Value() []byte {
	return n.value
}

func (n *node) IsLeaf() bool {
	return n.isLeaf
}

func (n *node) Bitmap() uint32 {
	return n.bitmap
}

func (n *node) Children() []*node {
	return n.children
}

func (n *node) Child(pos int) *node {
	return n.children[pos]
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
