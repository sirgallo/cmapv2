package cmap

func (cMap *CMap) NewLNode(key []byte, value []byte) *Node {
	return &Node{
		IsLeaf: true,
		Key: key,
		Value: value,
	}
}

func (cMap *CMap) NewINode() *Node {
	return &Node{
		IsLeaf: false,
		Bitmap: 0,
		Children: []*Node{},
	}
}

func (cMap *CMap) CopyNode(node *Node) *Node {
	nodeCopy := &Node{}
	nodeCopy.Key = node.Key
	nodeCopy.Value = node.Value
	nodeCopy.IsLeaf = node.IsLeaf
	nodeCopy.Bitmap = node.Bitmap
	nodeCopy.Children = make([]*Node, len(node.Children))

	copy(nodeCopy.Children, node.Children)
	return nodeCopy
}

func ExtendTable(orig []*Node, bitMap uint32, pos int, newNode *Node) []*Node {
	tableSize := calculateHammingWeight(bitMap)
	newTable := make([]*Node, tableSize)

	copy(newTable[:pos], orig[:pos])
	newTable[pos] = newNode
	copy(newTable[pos+1:], orig[pos:])
	return newTable
}

func ShrinkTable(orig []*Node, bitMap uint32, pos int) []*Node {
	tableSize := calculateHammingWeight(bitMap)
	newTable := make([]*Node, tableSize)

	copy(newTable[:pos], orig[:pos])
	copy(newTable[pos:], orig[pos+1:])
	return newTable
}
