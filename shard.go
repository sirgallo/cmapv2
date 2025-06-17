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
	return sMap.getShard(key).Put(key, value)
}

func (sMap *shardedMap) Get(key []byte) []byte {
	return sMap.getShard(key).Get(key)
}

func (sMap *shardedMap) Delete(key []byte) bool {
	return sMap.getShard(key).Delete(key)
}

func (s *shardedMap) getShard(key []byte) CMap {
	return s.shards[Murmur64(key, 1)%uint64(len(s.shards))]
}
