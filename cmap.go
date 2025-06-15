package cmap

import (
	"unsafe"
)

func NewShardedMap(shards int) CMap {
	s := &shardedMap{shards: make([]CMap, shards)}
	for i := range s.shards {
		s.shards[i] = NewCMap()
	}
	return s
}

func NewCMap() CMap {
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
