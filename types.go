package cmap

import (
	"sync/atomic"
)

type CMap interface {
	Root(shard ...int) Node
	Put(key []byte, value []byte) bool
	Get(key []byte) []byte
	Delete(key []byte) bool
}

type cMap struct {
	root         atomic.Pointer[node]
	bitChunkSize int
	hashChunks   int
}

type shardedMap struct {
	shards []CMap
}

type Node interface {
	Key() []byte
	Value() []byte
	IsLeaf() bool
	Bitmap() uint32
	Children() []Node
	Child(pos int) Node
	PrintChildren()
}

type node struct {
	key      []byte
	value    []byte
	isLeaf   bool
	bitmap   uint32
	children []*node
}
