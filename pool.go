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
	node.key = nil
	node.value = nil
	node.isLeaf = false
	node.bitmap = 0
	node.children = nil

	return node
}
