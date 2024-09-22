package main

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestDijkstra(t *testing.T) {
	graph := NewGraph()
	dijkstra := newDijkstra(graph)

	graph.AddVertex("D", "A", 4)
	graph.AddVertex("D", "E", 2)
	graph.AddVertex("A", "E", 4)
	graph.AddVertex("A", "C", 5)
	graph.AddVertex("E", "G", 5)
	graph.AddVertex("E", "C", 4)
	graph.AddVertex("C", "G", 5)
	graph.AddVertex("C", "F", 5)
	graph.AddVertex("C", "B", 2)
	graph.AddVertex("G", "F", 5)
	graph.AddVertex("B", "F", 2)

	weight, path, err := dijkstra.Calculate("D")
	if err != nil {
		t.Error(err)
		return
	}

	const expectedWeight = 10
	if weight != expectedWeight {
		t.Errorf("expected weight: %d, actual weight: %d", expectedWeight, weight)
		return
	}

	expectedPath := "DECBF"
	if expectedPath != path {
		t.Errorf("expected path: %s, actual path: %s", string(expectedPath), string(path))
		return
	}
}

func FuzzCalculation(f *testing.F) {
	graph := NewGraph()
	dijkstra := newDijkstra(graph)

	graph.AddVertex("D", "A", 4)
	graph.AddVertex("D", "E", 2)
	graph.AddVertex("A", "E", 4)
	graph.AddVertex("A", "C", 4)
	graph.AddVertex("E", "G", 5)
	graph.AddVertex("E", "G", 1)
	graph.AddVertex("E", "C", 4)
	graph.AddVertex("C", "G", 5)
	graph.AddVertex("C", "F", 5)
	graph.AddVertex("C", "B", 2)
	graph.AddVertex("G", "F", 5)
	graph.AddVertex("B", "F", 2)

	f.Add("D")
	f.Fuzz(func(t *testing.T, vertexName string) {
		dijkstra.Calculate(vertexName)
	})
}

func FuzzAdd(f *testing.F) {
	graph := NewGraph()

	f.Add("D", "A", 4)
	f.Add("D", "E", 2)
	f.Add("A", "E", 4)
	f.Add("A", "C", 4)
	f.Add("E", "G", 5)
	f.Add("E", "G", 1)
	f.Add("E", "C", 4)
	f.Add("C", "G", 5)
	f.Add("C", "F", 5)
	f.Add("C", "B", 2)
	f.Add("G", "F", 5)
	f.Add("B", "F", 2)

	f.Fuzz(func(t *testing.T, vertexAName, vertexBName string, weight int) {
		graph.AddVertex(vertexAName, vertexBName, int32(weight))
	})
}

func BenchmarkDijkstra(b *testing.B) {
	graph := NewGraph()
	dijkstra := newDijkstra(graph)

	// Generate fixtures
	numEdges := 100_000
	numNodes := 100_000
	var startVertex string
	for i := 0; i < numEdges; i++ {
		from := "A" + strconv.Itoa(rand.Intn(numNodes))
		for i := 0; i < rand.Intn(1000); i++ {
			to := "A" + strconv.Itoa(rand.Intn(numNodes))
			if from != to {
				weight := int32(rand.Intn(100) + 1) // weight between 1 and 100
				graph.AddVertex(
					from,
					to,
					weight,
				)
				startVertex = from
			}
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, err := dijkstra.Calculate(startVertex)
		if err != nil {
			b.Fatal(err)
		}
	}
}
