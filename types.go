package cmap

import (
	"sync"
	"unsafe"
)

type CMap interface {
	Root() Node
	Put(key []byte, value []byte) bool
	Get(key []byte) []byte
	Delete(key []byte) bool
	CalculateHashForCurrentLevel(key []byte, hash uint32, level int) uint32
}

type shardedMap struct {
	shards []CMap
}

type cMap struct {
	root         unsafe.Pointer
	bitChunkSize int
	hashChunks   int
	pool         Pool
}

type Node interface {
	Key() []byte
	Value() []byte
	IsLeaf() bool
	Bitmap() uint32
	Children() []*node
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
