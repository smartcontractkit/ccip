package graph

import "sync"

type GraphContainer interface {
	Get(id string) (Graph, bool)
	GetOrCreate(id string) Graph
	Set(id string, g Graph)
}

type graphContainer struct {
	graphs map[string]Graph
	lock   sync.RWMutex
}

func NewGraphContainer() GraphContainer {
	return &graphContainer{
		graphs: make(map[string]Graph),
	}
}

func (c *graphContainer) Get(id string) (Graph, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	g, ok := c.graphs[id]
	return g, ok
}

func (c *graphContainer) GetOrCreate(id string) Graph {
	c.lock.Lock()
	defer c.lock.Unlock()

	g, ok := c.graphs[id]
	if !ok {
		g = NewGraph()
		c.graphs[id] = g
	}

	return g
}

func (c *graphContainer) Set(id string, graph Graph) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.graphs[id] = graph
}
