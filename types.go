package cmap

import (
	"unsafe"
)

type ShardedMap struct {
	shards []*CMap
}

type Node struct {
	Key      []byte
	Value    []byte
	IsLeaf   bool
	Bitmap   uint32
	Children []*Node
}

type CMap struct {
	Root         unsafe.Pointer
	BitChunkSize int
	HashChunks   int
}
