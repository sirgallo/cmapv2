package cmap

import (
	"bytes"
	"runtime"
	"sync/atomic"
	"unsafe"
)

func (cMap *cMap) Root(shard ...int) Node {
	return (*node)(atomic.LoadPointer(&cMap.root))
}

func (cMap *cMap) Put(key []byte, value []byte) bool {
	for {
		completed := cMap.putRecursive(&cMap.root, key, value, 0, 0)
		if completed {
			return true
		}

		runtime.Gosched()
	}
}

func (cMap *cMap) putRecursive(nodePtr *unsafe.Pointer, key []byte, value []byte, hash uint32, level int) bool {
	hash = cMap.calculateHashForCurrentLevel(key, hash, level)
	index := cMap.getSparseIndex(hash, level)
	currNode := (*node)(atomic.LoadPointer(nodePtr))
	nodeCopy := cMap.CopyNode(currNode)

	if !IsBitSet(nodeCopy.Bitmap(), index) {
		newLeaf := cMap.NewLNode(key, value)
		nodeCopy.bitmap = SetBit(nodeCopy.Bitmap(), index)
		pos := cMap.getPosition(nodeCopy.Bitmap(), hash, level)
		nodeCopy.children = cMap.ExtendTable(nodeCopy.children, nodeCopy.Bitmap(), pos, newLeaf)
		return cMap.compareAndSwap(nodePtr, currNode, nodeCopy)
	} else {
		pos := cMap.getPosition(nodeCopy.Bitmap(), hash, level)
		if nodeCopy.Child(pos).IsLeaf() {
			if bytes.Equal(key, nodeCopy.Child(pos).Key()) {
				nodeCopy.Child(pos).value = value
				return cMap.compareAndSwap(nodePtr, currNode, nodeCopy)
			} else {
				newINode := cMap.NewINode()
				iNodePtr := unsafe.Pointer(newINode)
				cMap.putRecursive(&iNodePtr, nodeCopy.Child(pos).Key(), nodeCopy.Child(pos).Value(), 0, level+1)
				cMap.putRecursive(&iNodePtr, key, value, hash, level+1)
				nodeCopy.children[pos] = (*node)(atomic.LoadPointer(&iNodePtr))
				return cMap.compareAndSwap(nodePtr, currNode, nodeCopy)
			}
		} else {
			childPtr := unsafe.Pointer(nodeCopy.Child(pos))
			cMap.putRecursive(&childPtr, key, value, hash, level+1)
			nodeCopy.children[pos] = (*node)(atomic.LoadPointer(&childPtr))
			return cMap.compareAndSwap(nodePtr, currNode, nodeCopy)
		}
	}
}

func (cMap *cMap) Get(key []byte) []byte {
	return cMap.getRecursive((*node)(atomic.LoadPointer(&cMap.root)), key, 0, 0)
}

func (cMap *cMap) getRecursive(node *node, key []byte, hash uint32, level int) []byte {
	hash = cMap.calculateHashForCurrentLevel(key, hash, level)
	if !IsBitSet(node.Bitmap(), cMap.getSparseIndex(hash, level)) {
		return nil
	} else {
		pos := cMap.getPosition(node.Bitmap(), hash, level)
		if node.Child(pos).IsLeaf() && bytes.Equal(key, node.Child(pos).Key()) {
			return node.Child(pos).Value()
		} else {
			return cMap.getRecursive(node.Child(pos), key, hash, level+1)
		}
	}
}

func (cMap *cMap) Delete(key []byte) bool {
	for {
		completed := cMap.deleteRecursive(&cMap.root, key, 0, 0)
		if completed {
			return true
		}

		runtime.Gosched()
	}
}

func (cMap *cMap) deleteRecursive(nodePtr *unsafe.Pointer, key []byte, hash uint32, level int) bool {
	hash = cMap.calculateHashForCurrentLevel(key, hash, level)
	index := cMap.getSparseIndex(hash, level)
	currNode := (*node)(atomic.LoadPointer(nodePtr))
	nodeCopy := cMap.CopyNode(currNode)

	if !IsBitSet(nodeCopy.Bitmap(), index) {
		return true
	} else {
		pos := cMap.getPosition(nodeCopy.Bitmap(), hash, level)
		if nodeCopy.Child(pos).IsLeaf() {
			if bytes.Equal(key, nodeCopy.Child(pos).Key()) {
				nodeCopy.bitmap = SetBit(nodeCopy.Bitmap(), index)
				nodeCopy.children = cMap.ShrinkTable(nodeCopy.children, nodeCopy.Bitmap(), pos)
				return cMap.compareAndSwap(nodePtr, currNode, nodeCopy)
			}

			return false
		} else {
			childPtr := unsafe.Pointer(nodeCopy.Child(pos))
			cMap.deleteRecursive(&childPtr, key, hash, level+1)
			popCount := calculateHammingWeight(nodeCopy.Bitmap())
			if popCount == 0 {
				nodeCopy.bitmap = SetBit(nodeCopy.Bitmap(), index)
				nodeCopy.children = cMap.ShrinkTable(nodeCopy.children, nodeCopy.Bitmap(), pos)
			}

			return cMap.compareAndSwap(nodePtr, currNode, nodeCopy)
		}
	}
}

func (cMap *cMap) compareAndSwap(nodePtr *unsafe.Pointer, currNode *node, nodeCopy *node) bool {
	if atomic.CompareAndSwapPointer(nodePtr, unsafe.Pointer(currNode), unsafe.Pointer(nodeCopy)) {
		return true
	} else {
		cMap.pool.PutNode(nodeCopy)
		return false
	}
}
