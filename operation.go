package cmap

import (
	"bytes"
	"runtime"
)

func (cMap *cMap) Root(shard ...int) Node {
	return cMap.root.Load()
}

func (cMap *cMap) Put(key []byte, value []byte) bool {
	for {
		root := cMap.root.Load()
		if cMap.compareAndSwap(root, cMap.putRecursive(root, key, value, 0, 0)) {
			return true
		}
		runtime.Gosched()
	}
}

func (cMap *cMap) putRecursive(currNode *node, key []byte, value []byte, hash uint32, level int) *node {
	hash = cMap.calculateHashForCurrentLevel(key, hash, level)
	index := cMap.getSparseIndex(hash, level)

	var nodeCopy *node
	if currNode == nil {
		nodeCopy = NewINode()
	} else {
		nodeCopy = currNode.copyNode()
	}

	pos := cMap.getPosition(nodeCopy.Bitmap(), hash, level)
	if !IsBitSet(nodeCopy.Bitmap(), index) {
		nodeCopy.setBitmap(SetBit(nodeCopy.Bitmap(), index))
		nodeCopy.extendTable(nodeCopy.Bitmap(), pos, NewLNode(key, value))
	} else {
		if nodeCopy.Child(pos).IsLeaf() {
			if bytes.Equal(key, nodeCopy.Child(pos).Key()) {
				nodeCopy.setChild(NewLNode(key, value), pos)
			} else {
				nodeCopy.setChild(
					cMap.putRecursive(
						cMap.putRecursive(nil, nodeCopy.Child(pos).Key(), nodeCopy.Child(pos).Value(), 0, level+1), key, value, hash, level+1), pos)
			}
		} else {
			nodeCopy.setChild(
				cMap.putRecursive(nodeCopy.children[pos], key, value, hash, level+1), pos)
		}
	}
	return nodeCopy
}

func (cMap *cMap) Get(key []byte) []byte {
	return cMap.getRecursive(cMap.root.Load(), key, 0, 0)
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
			return cMap.getRecursive(node.children[pos], key, hash, level+1)
		}
	}
}

func (cMap *cMap) Delete(key []byte) bool {
	for {
		root := cMap.root.Load()
		if cMap.compareAndSwap(root, cMap.deleteRecursive(root, key, 0, 0)) {
			return true
		}
		runtime.Gosched()
	}
}

func (cMap *cMap) deleteRecursive(currNode *node, key []byte, hash uint32, level int) *node {
	hash = cMap.calculateHashForCurrentLevel(key, hash, level)
	index := cMap.getSparseIndex(hash, level)
	nodeCopy := currNode.copyNode()

	if IsBitSet(nodeCopy.Bitmap(), index) {
		pos := cMap.getPosition(nodeCopy.Bitmap(), hash, level)
		if nodeCopy.Child(pos).IsLeaf() {
			if bytes.Equal(key, nodeCopy.Child(pos).Key()) {
				nodeCopy.setBitmap(ClearBit(nodeCopy.Bitmap(), index))
				nodeCopy.shrinkTable(nodeCopy.Bitmap(), pos)
			}
		} else {
			childCopy := cMap.deleteRecursive(nodeCopy.children[pos], key, hash, level+1)
			if calculateHammingWeight(childCopy.Bitmap()) == 0 {
				nodeCopy.setBitmap(ClearBit(nodeCopy.Bitmap(), index))
				nodeCopy.shrinkTable(nodeCopy.Bitmap(), pos)
			} else {
				nodeCopy.setChild(childCopy, pos)
			}
		}
	}
	return nodeCopy
}

func (cMap *cMap) compareAndSwap(currNode *node, nodeCopy *node) bool {
	return cMap.root.CompareAndSwap(currNode, nodeCopy)
}
