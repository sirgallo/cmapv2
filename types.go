package cmap

import (
	"sync"
	"unsafe"
)

type Node struct {
	Key	[]byte
	Value []byte
	IsLeaf bool
	Bitmap uint32
	Children []*Node
}

type CMap struct {
	Root unsafe.Pointer
	BitChunkSize int
	HashChunks int
	// Pool *Pool
}

type Pool struct {
	size int64
	pool *sync.Pool
	tablePoolMap *sync.Map
}
