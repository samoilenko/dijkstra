package main

import (
	"fmt"
	"math"
)

type VisitedVertex struct {
	Path         string // name of all vertices beginning from the entry vertex
	Weight       int32  // weight to this vertex from the start vertex
	IsCalculated bool   // it shows if path, weight from this vertex to other were calculated
}

type Dijkstra struct {
	graph   *Graph
	visited map[string]*VisitedVertex // visited/calculated graph's vertices
}

func (d *Dijkstra) getNextVertexName() string {
	var vertexName string
	var min int32 = math.MaxInt32
	for visitedVertexName, visitedVertex := range d.visited {
		if !visitedVertex.IsCalculated && min > visitedVertex.Weight {
			vertexName = visitedVertexName
			min = visitedVertex.Weight
		}
	}

	return vertexName
}

func (d *Dijkstra) inspectNeighbors(current string) {
	for neighborVertexName, neighborWeight := range d.graph.Vertices[current] {
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
		neighborVertex.Weight = d.visited[current].Weight + neighborWeight

		// increase path
		neighborVertex.Path = d.visited[current].Path + neighborVertexName
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

		// finds new vertex to calculate
		newVertexSource := d.getNextVertexName()

		// if all vertices is calculated, return results
		if newVertexSource == "" {
			return currentVertex.Weight, currentVertex.Path, nil
		} else {
			nextVertexName = newVertexSource
		}
	}
}

func newDijkstra(graph *Graph) *Dijkstra {
	return &Dijkstra{
		graph:   graph,
		visited: make(map[string]*VisitedVertex),
	}
}
