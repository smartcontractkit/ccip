package graph

import "sync"

type MultiGraph interface {
	Get(id string) (Graph, bool)
	GetOrCreate(id string) Graph
	Add(id string, g Graph) bool
}

type multiGraph struct {
	graphs map[string]Graph
	lock   sync.RWMutex
}

func NewMultiGraph() MultiGraph {
	return &multiGraph{
		graphs: make(map[string]Graph),
	}
}

func (g *multiGraph) Get(id string) (Graph, bool) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	gr, ok := g.graphs[id]
	return gr, ok
}

func (g *multiGraph) GetOrCreate(id string) Graph {
	g.lock.Lock()
	defer g.lock.Unlock()

	if gr, ok := g.graphs[id]; ok {
		return gr
	}
	gr := NewGraph()
	g.graphs[id] = gr

	return gr
}

func (g *multiGraph) Add(id string, graph Graph) bool {
	g.lock.Lock()
	defer g.lock.Unlock()

	if _, ok := g.graphs[id]; ok {
		return false
	}
	g.graphs[id] = graph

	return true
}
