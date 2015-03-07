package transportdata

import (
	"strings"
	"testing"
)

func setupDjikstraGraph() graph {
	graph := make(graph)
	graph.connectVertices("milsons point", "wynyard")
	graph.connectVertices("wynyard", "town hall")
	graph.connectVertices("town hall", "central")
	graph.connectVertices("town hall", "martin place")
	graph.connectVertices("martin place", "kings cross")
	graph.connectVertices("kings cross", "edgecliff")
	graph.connectVertices("edgecliff", "bondi")
	graph.connectVertices("central", "museum")
	graph.connectVertices("museum", "st james")
	graph.connectVertices("st james", "circular quay")
	graph.connectVertices("circular quay", "wynyard")
	return graph
}

func formatPath(path []vertex) string {
	vertexIds := []string{}
	for _, vertex := range path {
		vertexIds = append(vertexIds, vertex.vertexId)
	}
	return strings.Join(vertexIds, " -> ")
}

func TestSuccesfulPathSearch(t *testing.T) {
	graph := setupDjikstraGraph()
	start := graph["milsons point"]
	end := graph["st james"]
	expected := "milsons point -> wynyard -> circular quay -> st james"

	shortestPathSearch(graph, start)
	path := formatPath(end.pathFromStart())

	if path != expected {
		t.Fatalf("Expected %#v but got %#v", expected, path)
	}
}

func TestImpossibleRoute(t *testing.T) {
	graph := setupDjikstraGraph()
	graph.addVertex("nowhere")
	start := graph["milsons point"]
	end := graph["nowhere"]

	shortestPathSearch(graph, start)
	path := formatPath(end.pathFromStart())

	if path != "" {
		t.Fatalf("Expected empty path but got %#v", path)
	}
}

func BenchmarkDjikstra(b *testing.B) {
	graph := setupDjikstraGraph()
	start := graph["milsons point"]

	for n := 0; n < b.N; n++ {
		shortestPathSearch(graph, start)
	}
}