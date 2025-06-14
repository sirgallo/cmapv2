package cmap

import (
	"encoding/binary"
)

const (
	c32_1 = 0x85ebca6b
	c32_2 = 0xc2b2ae35
	c32_3 = 0xe6546b64
	c32_4 = 0x1b873593
	c32_5 = 0x5c4bcea9
)

func Murmur32(data []byte, seed uint32) uint32 {
	hash := seed
	length := uint32(len(data))
	total4ByteChunks := len(data) / 4
	for idx := range make([]int, total4ByteChunks) {
		startIdxOfChunk := idx * 4 
		endIdxOfChunk := (idx + 1) * 4
		chunk := binary.LittleEndian.Uint32(data[startIdxOfChunk:endIdxOfChunk])
		rotateRight32(&hash, chunk)
	}

	handleRemainingBytes32(&hash, data)

	hash ^= length
	hash ^= hash >> 16
	hash *= c32_4
	hash ^= hash >> 13
	hash *= c32_5
	hash ^= hash >> 16
	return hash
}

func rotateRight32(hash *uint32, chunk uint32) {
	chunk *= c32_1
	chunk = (chunk << 15) | (chunk >> 17) // Rotate right by 15
	chunk *= c32_2

	*hash ^= chunk
	*hash = (*hash << 13) | (*hash >> 19) // Rotate right by 13
	*hash = *hash * 5 + c32_3
}

func handleRemainingBytes32(hash *uint32, dataAsBytes []byte) {
	remaining := dataAsBytes[len(dataAsBytes)-len(dataAsBytes) % 4:]
	if len(remaining) > 0 {
		var chunk uint32
		
		switch len(remaining) {
		case 3:
			chunk |= uint32(remaining[2]) << 16
			fallthrough
		case 2:
			chunk |= uint32(remaining[1]) << 8
			fallthrough
		case 1:
			chunk |= uint32(remaining[0])
			chunk *= c32_1
			chunk = (chunk << 15) | (chunk >> 17) // Rotate right by 15
			chunk *= c32_2
			*hash ^= chunk
		}
	}
}
