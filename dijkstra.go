package main

import (
	"fmt"
)

type VisitedVertex struct {
	Path         string // name of all vertices beginning from the entry vertex
	Weight       int32  // weight to this vertex from the start vertex
	IsCalculated bool   // it shows if path, weight from this vertex to other were calculated
}

type Dijkstra struct {
	graph   *Graph
	visited map[string]*VisitedVertex // visited/calculated graph's vertices
	heap    *HeapMin
}

func (d *Dijkstra) inspectNeighbors(current string) {
	for neighborVertexName, neighborWeight := range d.graph.Vertices[current] {
		destinationWeight := d.visited[current].Weight + neighborWeight

		// if a vertex was visited earlier and new weight will be bigger that existing one
		// then this path will be longer and it will be skipped
		neighborVertex, ok := d.visited[neighborVertexName]
		if !ok {
			neighborVertex = &VisitedVertex{}
			d.visited[neighborVertexName] = neighborVertex
		} else if d.visited[current].Weight+neighborWeight >= neighborVertex.Weight {
			continue
		}

		// accumulate weight
		neighborVertex.Weight = destinationWeight

		// increase path
		neighborVertex.Path = d.visited[current].Path + neighborVertexName
		d.heap.Add(neighborVertexName, destinationWeight)
	}
}

func (d *Dijkstra) Calculate(from string) (weight int32, path string, err error) {
	if _, ok := d.graph.Vertices[from]; !ok {
		return 0, "", fmt.Errorf("%s does not exist in Graph", from)
	}

	nextVertexName := from
	for {
		currentVertex, ok := d.visited[nextVertexName]
		if !ok {
			currentVertex = &VisitedVertex{Path: nextVertexName}
			d.visited[nextVertexName] = currentVertex
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
		visited: make(map[string]*VisitedVertex),
		heap:    &HeapMin{tree: make([]*Item, 0)},
	}
}
