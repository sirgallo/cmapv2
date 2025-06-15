package cmap

import (
	"bytes"
	"runtime"
	"sync/atomic"
	"unsafe"
)

func (sMap *ShardedMap) Put(key []byte, value []byte) bool {
	shard := sMap.getShard(key)
	return shard.Put(key, value)
}

func (cMap *CMap) Put(key []byte, value []byte) bool {
	for {
		completed := cMap.putRecursive(&cMap.Root, key, value, 0, 0)
		if completed {
			return true
		}

		runtime.Gosched()
	}
}

func (cMap *CMap) putRecursive(node *unsafe.Pointer, key []byte, value []byte, hash uint32, level int) bool {
	hash = cMap.CalculateHashForCurrentLevel(key, hash, level)
	index := cMap.getSparseIndex(hash, level)
	currNode := (*Node)(atomic.LoadPointer(node))
	nodeCopy := cMap.CopyNode(currNode)

	if !IsBitSet(nodeCopy.Bitmap, index) {
		newLeaf := cMap.NewLNode(key, value)
		nodeCopy.Bitmap = SetBit(nodeCopy.Bitmap, index)
		pos := cMap.getPosition(nodeCopy.Bitmap, hash, level)
		nodeCopy.Children = cMap.ExtendTable(nodeCopy.Children, nodeCopy.Bitmap, pos, newLeaf)
		return cMap.compareAndSwap(node, currNode, nodeCopy)
	} else {
		pos := cMap.getPosition(nodeCopy.Bitmap, hash, level)
		if nodeCopy.Children[pos].IsLeaf {
			if bytes.Equal(key, nodeCopy.Children[pos].Key) {
				nodeCopy.Children[pos].Value = value
				return cMap.compareAndSwap(node, currNode, nodeCopy)
			} else {
				newINode := cMap.NewINode()
				iNodePtr := unsafe.Pointer(newINode)
				cMap.putRecursive(&iNodePtr, nodeCopy.Children[pos].Key, nodeCopy.Children[pos].Value, 0, level+1)
				cMap.putRecursive(&iNodePtr, key, value, hash, level+1)

				nodeCopy.Children[pos] = (*Node)(atomic.LoadPointer(&iNodePtr))
				return cMap.compareAndSwap(node, currNode, nodeCopy)
			}
		} else {
			childPtr := unsafe.Pointer(nodeCopy.Children[pos])
			cMap.putRecursive(&childPtr, key, value, hash, level+1)
			nodeCopy.Children[pos] = (*Node)(atomic.LoadPointer(&childPtr))
			return cMap.compareAndSwap(node, currNode, nodeCopy)
		}
	}
}

func (sMap *ShardedMap) Get(key []byte) []byte {
	shard := sMap.getShard(key)
	return shard.Get(key)
}

func (cMap *CMap) Get(key []byte) []byte {
	root := (*Node)(atomic.LoadPointer(&cMap.Root))
	return cMap.getRecursive(root, key, 0, 0)
}

func (cMap *CMap) getRecursive(node *Node, key []byte, hash uint32, level int) []byte {
	hash = cMap.CalculateHashForCurrentLevel(key, hash, level)
	if !IsBitSet(node.Bitmap, cMap.getSparseIndex(hash, level)) {
		return nil
	} else {
		pos := cMap.getPosition(node.Bitmap, hash, level)
		if node.Children[pos].IsLeaf && bytes.Equal(key, node.Children[pos].Key) {
			return node.Children[pos].Value
		} else {
			return cMap.getRecursive(node.Children[pos], key, hash, level+1)
		}
	}
}

func (sMap *ShardedMap) Delete(key []byte) bool {
	shard := sMap.getShard(key)
	return shard.Delete(key)
}

func (cMap *CMap) Delete(key []byte) bool {
	for {
		completed := cMap.deleteRecursive(&cMap.Root, key, 0, 0)
		if completed {
			return true
		}

		runtime.Gosched()
	}
}

func (cMap *CMap) deleteRecursive(node *unsafe.Pointer, key []byte, hash uint32, level int) bool {
	hash = cMap.CalculateHashForCurrentLevel(key, hash, level)
	index := cMap.getSparseIndex(hash, level)
	currNode := (*Node)(atomic.LoadPointer(node))
	nodeCopy := cMap.CopyNode(currNode)

	if !IsBitSet(nodeCopy.Bitmap, index) {
		return true
	} else {
		pos := cMap.getPosition(nodeCopy.Bitmap, hash, level)
		if nodeCopy.Children[pos].IsLeaf {
			if bytes.Equal(key, nodeCopy.Children[pos].Key) {
				nodeCopy.Bitmap = SetBit(nodeCopy.Bitmap, index)
				nodeCopy.Children = cMap.ShrinkTable(nodeCopy.Children, nodeCopy.Bitmap, pos)
				return cMap.compareAndSwap(node, currNode, nodeCopy)
			}

			return false
		} else {
			childPtr := unsafe.Pointer(nodeCopy.Children[pos])
			cMap.deleteRecursive(&childPtr, key, hash, level+1)
			popCount := calculateHammingWeight(nodeCopy.Bitmap)
			if popCount == 0 {
				nodeCopy.Bitmap = SetBit(nodeCopy.Bitmap, index)
				nodeCopy.Children = cMap.ShrinkTable(nodeCopy.Children, nodeCopy.Bitmap, pos)
			}

			return cMap.compareAndSwap(node, currNode, nodeCopy)
		}
	}
}

func (cMap *CMap) compareAndSwap(node *unsafe.Pointer, currNode *Node, nodeCopy *Node) bool {
	if atomic.CompareAndSwapPointer(node, unsafe.Pointer(currNode), unsafe.Pointer(nodeCopy)) {
		return true
	} else {
		cMap.pool.putNode(nodeCopy)
		return false
	}
}
