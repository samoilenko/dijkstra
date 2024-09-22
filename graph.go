package main

type Graph struct {
	Vertices map[string]map[string]int32 // graph's vertices. graph vertex with their relations
}

func (g *Graph) AddVertex(vertexNameA, vertexNameB string, weight int32) {
	vertexA, ok := g.Vertices[vertexNameA]
	if !ok {
		vertexA = make(map[string]int32)
		g.Vertices[vertexNameA] = vertexA
	}
	vertexA[vertexNameB] = weight

	vertexB, ok := g.Vertices[vertexNameB]
	if !ok {
		vertexB = make(map[string]int32)
		g.Vertices[vertexNameB] = vertexB
	}
	vertexB[vertexNameA] = weight
	g.Vertices[vertexNameB] = vertexB
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: make(map[string]map[string]int32),
	}
}
