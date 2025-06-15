package cmap

import (
	"sync"
)

func newPool() *Pool {
	np := sync.Pool{
		New: func() any {
			return &Node{}
		},
	}

	return &Pool{
		node: &np,
	}
}

func (p *Pool) getNode() *Node {
	return p.node.Get().(*Node)
}

func (p *Pool) putNode(node *Node) {
	p.node.Put(p.resetNode(node))
}

func (p *Pool) resetNode(node *Node) *Node {
	node.Key = nil
	node.Value = nil
	node.IsLeaf = false
	node.Bitmap = 0
	node.Children = nil

	return node
}
