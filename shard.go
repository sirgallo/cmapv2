package cmap

import (
	"log"
)

func (sMap *shardedMap) Root(shards ...int) Node {
	if len(shards) != 1 {
		log.Fatalf("please provide the shard number")
	}
	return sMap.shards[shards[0]].Root()
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

func (s *shardedMap) getShard(key []byte) CMap {
	h := Murmur32(key, 1) % uint32(len(s.shards))
	return s.shards[h]
}
