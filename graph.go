package main

import (
	"github.com/samoilenko/swiss"
)

type Graph struct {
	Vertices *swiss.Map[string, *swiss.Map[string, int32]] // graph's vertices. graph vertex with their relations
}

func (g *Graph) AddVertex(vertexNameA, vertexNameB string, weight int32) {
	vertexA, ok := g.Vertices.Get(vertexNameA)
	if !ok {
		vertexA = swiss.NewMap[string, int32](0)
		g.Vertices.Put(vertexNameA, vertexA)
	}
	vertexA.Put(vertexNameB, weight)

	vertexB, ok := g.Vertices.Get(vertexNameB)
	if !ok {
		vertexB = swiss.NewMap[string, int32](0)
		g.Vertices.Put(vertexNameB, vertexB)
	}
	vertexB.Put(vertexNameA, weight)
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: swiss.NewMap[string, *swiss.Map[string, int32]](1000),
	}
}
