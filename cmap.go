package cmap

import (
	"unsafe"
)

func NewShardedMap(shards int) *ShardedMap {
	s := &ShardedMap{shards: make([]*CMap, shards)}
	for i := range s.shards {
		s.shards[i] = NewCMap()
	}
	return s
}

func NewCMap() *CMap {
	rootNode := &Node{}
	rootNode.IsLeaf = false
	rootNode.Bitmap = 0
	rootNode.Children = []*Node{}

	return &CMap{
		BitChunkSize: 5,
		HashChunks:   6,
		Root:         unsafe.Pointer(rootNode),
	}
}

func (s *ShardedMap) getShard(key []byte) *CMap {
	h := Murmur32(key, 1) % uint32(len(s.shards))
	return s.shards[h]
}
