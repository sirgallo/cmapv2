package cmap

import (
	"sync"
	"unsafe"
)

type CMap interface {
	Root(shard ...int) Node
	Put(key []byte, value []byte) bool
	Get(key []byte) []byte
	Delete(key []byte) bool
}

type cMap struct {
	root         unsafe.Pointer
	bitChunkSize int
	hashChunks   int
	pool         Pool
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

type Pool interface {
	GetNode() *node
	PutNode(n *node)
}

type pool struct {
	node *sync.Pool
}
