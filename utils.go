package cmap

import (
	"math"
	"math/bits"
)

func (cMap *cMap) calculateHashForCurrentLevel(key []byte, hash uint32, level int) uint32 {
	if level%cMap.hashChunks == 0 || hash == 0 {
		currChunk := level / cMap.hashChunks
		seed := uint32(currChunk + 1)
		return Murmur32(key, seed)
	}

	return hash
}

func (cMap *cMap) getSparseIndex(hash uint32, level int) int {
	return GetIndexForLevel(hash, cMap.bitChunkSize, level, cMap.hashChunks)
}

func (cMap *cMap) getPosition(bitMap uint32, hash uint32, level int) int {
	sparseIdx := GetIndexForLevel(hash, cMap.bitChunkSize, level, cMap.hashChunks)
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
