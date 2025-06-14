package cmap

import (
	"unsafe"
)

func NewCMap() *CMap {	
	rootNode := &Node{}
	rootNode.IsLeaf = false
	rootNode.Bitmap = 0
	rootNode.Children = []*Node{}

	return &CMap{
		BitChunkSize: 5,
		HashChunks: 6,
		Root: unsafe.Pointer(rootNode),
	}
}
