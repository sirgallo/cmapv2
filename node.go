package cmap

func (cMap *CMap) NewLNode(key []byte, value []byte) *Node {
	node := cMap.pool.getNode()
	node.IsLeaf = true
	node.Key = key
	node.Value = value
	return node
}

func (cMap *CMap) NewINode() *Node {
	node := cMap.pool.getNode()
	node.IsLeaf = false
	node.Bitmap = 0
	node.Children = []*Node{}
	return node
}

func (cMap *CMap) CopyNode(node *Node) *Node {
	nodeCopy := cMap.pool.getNode()
	nodeCopy.Key = node.Key
	nodeCopy.Value = node.Value
	nodeCopy.IsLeaf = node.IsLeaf
	nodeCopy.Bitmap = node.Bitmap
	nodeCopy.Children = make([]*Node, len(node.Children))

	copy(nodeCopy.Children, node.Children)
	return nodeCopy
}

func (cMap *CMap) ExtendTable(orig []*Node, bitMap uint32, pos int, newNode *Node) []*Node {
	tableSize := calculateHammingWeight(bitMap)
	newTable := make([]*Node, tableSize)

	copy(newTable[:pos], orig[:pos])
	newTable[pos] = newNode
	copy(newTable[pos+1:], orig[pos:])
	return newTable
}

func (cMap *CMap) ShrinkTable(orig []*Node, bitMap uint32, pos int) []*Node {
	tableSize := calculateHammingWeight(bitMap)
	newTable := make([]*Node, tableSize)

	copy(newTable[:pos], orig[:pos])
	copy(newTable[pos:], orig[pos+1:])
	return newTable
}
