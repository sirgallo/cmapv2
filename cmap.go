package cmap

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
	m := cMap{
		bitChunkSize: 5,
		hashChunks:   6,
	}

	m.root.Store(&node{
		isLeaf:   false,
		bitmap:   0,
		children: []*node{},
	})

	return &m
}
