package cmap

import (
	"sync"
)

func newPool() *Pool {
	size := int64(0)
	np := &Pool{ size: size }
	np.pool = &sync.Pool { 
		New: func() any { return np.resetNode(&Node{}) },
	}

	np.tablePoolMap = &sync.Map{}
	return np
}

func (p *Pool) getNode() *Node {
	node := p.pool.Get().(*Node)
	return node
}

func (p *Pool) putNode(node *Node) {
	p.pool.Put(p.resetNode(node))
}

func (p *Pool) getTable(tableSize int) []*Node {
	pool, _ := p.tablePoolMap.LoadOrStore(tableSize, &sync.Pool{
		New: func() any {
			return make([]*Node, tableSize)
		},
	})
	return pool.(*sync.Pool).Get().([]*Node)
}

func (p *Pool) putTable(tableSize int, table []*Node) {
	pool, _ := p.tablePoolMap.LoadOrStore(tableSize, &sync.Pool{
		New: func() any {
			return make([]*Node, tableSize)
		},
	})
	pool.(*sync.Pool).Put(table)
}

func (p *Pool) resetNode(node *Node) *Node {
	node.Key = nil
	node.Value = nil
	node.IsLeaf = false
	node.Bitmap = 0
	p.putTable(len(node.Children), node.Children)
	node.Children = nil

	return node
}
