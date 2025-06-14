package cmap

import (
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
}

