package cmap

import (
	"sync"
)

func newPool() Pool {
	np := sync.Pool{
		New: func() any {
			return &node{}
		},
	}

	return &pool{
		node: &np,
	}
}

func (p *pool) GetNode() *node {
	return p.node.Get().(*node)
}

func (p *pool) PutNode(node *node) {
	p.node.Put(p.resetNode(node))
}

func (p *pool) resetNode(node *node) *node {
	node.setKey(nil)
	node.setValue(nil)
	node.setIsLeaf(false)
	node.setBitmap(0)
	node.setChildren(nil)
	return node
}
