package cmap

import (
	"math"
	"math/bits"
)

func (cMap *cMap) calculateHashForCurrentLevel(key []byte, hash uint32, level int) uint32 {
	if level%cMap.hashChunks == 0 || hash == 0 {
		return Murmur32(key, uint32(level/cMap.hashChunks+1))
	}

	return hash
}

func (cMap *cMap) getSparseIndex(hash uint32, level int) int {
	return GetIndexForLevel(hash, cMap.bitChunkSize, level, cMap.hashChunks)
}

func (cMap *cMap) getPosition(bitMap uint32, hash uint32, level int) int {
	sparseIdx := GetIndexForLevel(hash, cMap.bitChunkSize, level, cMap.hashChunks)
	return calculateHammingWeight(bitMap & uint32((1<<sparseIdx)-1))
}

func GetIndexForLevel(hash uint32, chunkSize int, level int, hashChunks int) int {
	return GetIndex(hash, chunkSize, level%hashChunks)
}

func GetIndex(hash uint32, chunkSize int, level int) int {
	slots := int(math.Pow(float64(2), float64(chunkSize)))
	shiftSize := slots - (chunkSize * (level + 1))
	return int(hash >> shiftSize & uint32(slots-1))
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
