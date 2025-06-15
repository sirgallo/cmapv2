package cmap

func (s *shardedMap) getShard(key []byte) CMap {
	h := Murmur32(key, 1) % uint32(len(s.shards))
	return s.shards[h]
}

func (sMap *shardedMap) Root() Node {
	return nil
}

func (sMap *shardedMap) Put(key []byte, value []byte) bool {
	shard := sMap.getShard(key)
	return shard.Put(key, value)
}

func (sMap *shardedMap) Get(key []byte) []byte {
	shard := sMap.getShard(key)
	return shard.Get(key)
}

func (sMap *shardedMap) Delete(key []byte) bool {
	shard := sMap.getShard(key)
	return shard.Delete(key)
}

func (sMap *shardedMap) CalculateHashForCurrentLevel(key []byte, hash uint32, level int) uint32 {
	return 0
}
