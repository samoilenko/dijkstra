package main

import (
	"fmt"

	"github.com/samoilenko/swiss"
)

type VisitedVertex struct {
	Path         string // name of all vertices beginning from the entry vertex
	Weight       int32  // weight to this vertex from the start vertex
	IsCalculated bool   // it shows if path, weight from this vertex to other were calculated
}

type Dijkstra struct {
	graph   *Graph
	visited *swiss.Map[string, *VisitedVertex] // visited/calculated graph's vertices
	heap    *HeapMin
}

func (d *Dijkstra) inspectNeighbors(current string) {
	neighbors, _ := d.graph.Vertices.Get(current)
	currentVertex, _ := d.visited.Get(current)

	for neighborVertexName, neighborWeight := range neighbors.Iterator() {
		destinationWeight := currentVertex.Weight + neighborWeight

		// if a vertex was visited earlier and new weight will be bigger that existing one
		// then this path will be longer and it will be skipped
		visitedNeighbor, ok := d.visited.Get(neighborVertexName)

		if !ok {
			visitedNeighbor = &VisitedVertex{}
			d.visited.Put(neighborVertexName, visitedNeighbor)
		} else if destinationWeight >= visitedNeighbor.Weight {
			continue
		}

		// accumulate weight
		visitedNeighbor.Weight = destinationWeight

		// increase path
		visitedNeighbor.Path = currentVertex.Path + neighborVertexName

		d.heap.Add(neighborVertexName, destinationWeight)
	}
}

func (d *Dijkstra) Calculate(from string) (weight int32, path string, err error) {
	if _, ok := d.graph.Vertices.Get(from); !ok {
		return 0, "", fmt.Errorf("%s does not exist in Graph", from)
	}

	nextVertexName := from
	for {
		currentVertex, ok := d.visited.Get(nextVertexName)
		if !ok {
			currentVertex = &VisitedVertex{Path: nextVertexName}
			d.visited.Put(nextVertexName, currentVertex)
		}

		currentVertex.IsCalculated = true
		d.inspectNeighbors(nextVertexName)

		newVertexSourceName := d.heap.GetRoot()

		// if all vertices is calculated, return results
		if newVertexSourceName == "" {
			return currentVertex.Weight, currentVertex.Path, nil
		} else {
			nextVertexName = newVertexSourceName
		}
	}
}

func newDijkstra(graph *Graph) *Dijkstra {
	return &Dijkstra{
		graph:   graph,
		visited: swiss.NewMap[string, *VisitedVertex](100),
		heap:    &HeapMin{tree: make([]*Item, 0)},
	}
}
