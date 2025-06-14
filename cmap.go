package cmap

import (
	"math"
	"unsafe"
)

func NewCMap() *CMap {
	bitChunkSize := 5
	hashChunks := int(math.Pow(float64(2), float64(bitChunkSize))) / bitChunkSize
	
	rootNode := &Node{}
	rootNode.IsLeaf = false
	rootNode.Bitmap = 0
	rootNode.Children = []*Node{}

	return &CMap{
		BitChunkSize: bitChunkSize,
		HashChunks: hashChunks,
		Root: unsafe.Pointer(rootNode),
	}
}
