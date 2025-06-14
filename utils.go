package cmap

import (
	"fmt"
	"math"
	"math/bits"
	"sync/atomic"
	"unsafe"
)

func (cMap *CMap) CalculateHashForCurrentLevel(key []byte, hash uint32, level int) uint32 {
	if level % 6 != 0 {
		return hash
	}
	
	currChunk := level / cMap.HashChunks
	seed := uint32(currChunk + 1)
	return Murmur32(key, seed)
}

func (cMap *CMap) getSparseIndex(hash uint32, level int) int {
	return GetIndexForLevel(hash, cMap.BitChunkSize, level, cMap.HashChunks)
}

func (cMap *CMap) getPosition(bitMap uint32, hash uint32, level int) int {
	sparseIdx := GetIndexForLevel(hash, cMap.BitChunkSize, level, cMap.HashChunks)
	mask := uint32((1 << sparseIdx) - 1)
	isolatedBits := bitMap & mask
	return calculateHammingWeight(isolatedBits)
}

func GetIndexForLevel(hash uint32, chunkSize int, level int, hashChunks int) int {
	updatedLevel := level % hashChunks
	return GetIndex(hash, chunkSize, updatedLevel)
}

func GetIndex(hash uint32, chunkSize int, level int) int {
	slots := int(math.Pow(float64(2), float64(chunkSize)))
	shiftSize := slots - (chunkSize * (level + 1))
	mask := uint32(slots - 1)
	return int(hash >> shiftSize & mask)
}

func calculateHammingWeight(bitmap uint32) int {
	return bits.OnesCount32(bitmap)
}

func SetBit(bitmap uint32, position int) uint32 {
	return bitmap ^ (1 << position)
}

func IsBitSet(bitmap uint32, position int) bool {
	return (bitmap & (1 << position)) != 0
}

func (cMap *CMap) PrintChildren() {
	cMap.printChildrenRecursive(&cMap.Root, 0)
}

func (cMap *CMap) printChildrenRecursive(node *unsafe.Pointer, level int) {
	currNode := (*Node)(atomic.LoadPointer(node))
	if currNode == nil {
		return
	}

	for idx, child := range currNode.Children {
		if child != nil {
			fmt.Printf("Level: %d, Index: %d, Key: %s, Value: %v\n", level, idx, child.Key, child.Value)

			childPtr := unsafe.Pointer(child)
			cMap.printChildrenRecursive(&childPtr, level+1)
		}
	}
}
