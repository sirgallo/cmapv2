package cmap

import (
	"unsafe"
)

func NewMap(shards ...int) CMap {
	if len(shards) == 1 && shards[0] > 1 {
		return newShardedMap(shards[0])
	}
	return newCMap()
}

func newShardedMap(shards int) CMap {
	s := &shardedMap{shards: make([]CMap, shards)}
	for i := range s.shards {
		s.shards[i] = newCMap()
	}
	return s
}

func newCMap() CMap {
	nPool := newPool()
	rootNode := nPool.GetNode()
	rootNode.isLeaf = false
	rootNode.bitmap = 0
	rootNode.children = []*node{}

	return &cMap{
		bitChunkSize: 5,
		hashChunks:   6,
		root:         unsafe.Pointer(rootNode),
		pool:         nPool,
	}
}
