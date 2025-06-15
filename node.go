package cmap

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
	nodeCopy.key = n.Key()
	nodeCopy.value = n.Value()
	nodeCopy.isLeaf = n.IsLeaf()
	nodeCopy.bitmap = n.Bitmap()
	nodeCopy.children = make([]*node, len(n.Children()))

	copy(nodeCopy.children, n.Children())
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