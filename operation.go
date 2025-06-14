package cmap

import (
	"bytes"
	"runtime"
	"sync/atomic"
	"unsafe"
)

func (cMap *CMap) Put(key []byte, value []byte) bool {
	for {
		completed := cMap.putRecursive(&cMap.Root, key, value, 0, 0)
		if completed { return true }
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
		nodeCopy.Children = ExtendTable(nodeCopy.Children, nodeCopy.Bitmap, pos, newLeaf)
		return cMap.compareAndSwap(node, currNode, nodeCopy)
	} else {
		pos := cMap.getPosition(nodeCopy.Bitmap, hash, level)
		childNode := nodeCopy.Children[pos]
		if childNode.IsLeaf {
			if bytes.Equal(key, childNode.Key) {
				nodeCopy.Children[pos].Value = value
				return cMap.compareAndSwap(node, currNode, nodeCopy)
			} else {
				newINode := cMap.NewINode()
				iNodePtr := unsafe.Pointer(newINode)
				cMap.putRecursive(&iNodePtr, childNode.Key, childNode.Value, 0, level+1)
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

func (cMap *CMap) Get(key []byte) []byte {
	return cMap.getRecursive(&cMap.Root, key, 0, 0)
}

func (cMap *CMap) getRecursive(node *unsafe.Pointer, key []byte, hash uint32, level int) []byte {
	hash = cMap.CalculateHashForCurrentLevel(key, hash, level)
	index := cMap.getSparseIndex(hash, level)
	currNode := (*Node)(atomic.LoadPointer(node))

	if !IsBitSet(currNode.Bitmap, index) {
		return nil
	} else {
		pos := cMap.getPosition(currNode.Bitmap, hash, level)
		childNode := currNode.Children[pos]
		if childNode.IsLeaf && bytes.Equal(key, childNode.Key) {
			return childNode.Value
		} else {
			childPtr := unsafe.Pointer(currNode.Children[pos])
			return cMap.getRecursive(&childPtr, key, hash, level+1)
		}
	}
}

func (cMap *CMap) Delete(key []byte) bool {
	for {
		completed := cMap.deleteRecursive(&cMap.Root, key, 0, 0)
		if completed { return true }
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
		childNode := nodeCopy.Children[pos]
		if childNode.IsLeaf {
			if bytes.Equal(key, childNode.Key) {
				nodeCopy.Bitmap = SetBit(nodeCopy.Bitmap, index)
				nodeCopy.Children = ShrinkTable(nodeCopy.Children, nodeCopy.Bitmap, pos)
				return cMap.compareAndSwap(node, currNode, nodeCopy)
			}

			return false
		} else {
			childPtr := unsafe.Pointer(nodeCopy.Children[pos])
			cMap.deleteRecursive(&childPtr, key, hash, level+1)
			popCount := calculateHammingWeight(nodeCopy.Bitmap)
			if popCount == 0 {
				nodeCopy.Bitmap = SetBit(nodeCopy.Bitmap, index)
				nodeCopy.Children = ShrinkTable(nodeCopy.Children, nodeCopy.Bitmap, pos)
			}

			return cMap.compareAndSwap(node, currNode, nodeCopy)
		}
	}
}

func (cMap *CMap) compareAndSwap(node *unsafe.Pointer, currNode *Node, nodeCopy *Node) bool {
	return atomic.CompareAndSwapPointer(node, unsafe.Pointer(currNode), unsafe.Pointer(nodeCopy))
}
