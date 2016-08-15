// 'directedgraph_test.go'.
// Chris Shiels.


package directedgraph


import (
    "sort"
    "testing"
)


func expectarraystringsequal(t *testing.T,
                             name string,
                             actual, expected []string) {
    if len(actual) != len(expected) {
        t.Errorf("%s: %v, want %v", name, actual, expected)
        return
    }

    i := 0
    for i = 0; i < len(expected); i++ {
        if actual[i] != expected[i] {
            t.Errorf("%s: %v, want %v", name, actual, expected)
            return
        }
    }
}


func expectarraystringsnotequal(t *testing.T,
                                name string,
                                actual, expected []string) {
    if len(actual) != len(expected) {
        return
    }

    i := 0
    for i = 0; i < len(expected); i++ {
        if actual[i] != expected[i] {
            return
        }
    }
    t.Errorf("%s: %v, want %v", name, actual, expected)
}


func Test_ShouldInstantiate(t *testing.T) {
    dg := NewDirectedGraph()
    if dg == nil {
        t.Errorf("Test_ShouldInstantiate: %v, want %v",
                 dg,
                 nil)
    }
}


func Test_ShouldManageVertices(t *testing.T) {
    dg := NewDirectedGraph()
    dg.AddVertex("a")
    dg.AddVertex("a")
    vertices := dg.Vertices()
    sort.Strings(vertices)
    expectarraystringsequal(t,
                            "Test_ShouldManageVertices",
                            vertices,
                            []string{ "a" })
}


func Test_ShouldManageEdges(t *testing.T) {
    dg := NewDirectedGraph()
    dg.AddVertex("a")
    dg.AddVertex("b")
    dg.AddEdge("a", "b")
    vertices := dg.Vertices()
    edges := dg.Edges("a")
    sort.Strings(vertices)
    sort.Strings(edges)
    expectarraystringsequal(t,
                            "Test_ShouldManageEdges",
                            vertices,
                            []string{ "a", "b" })
    expectarraystringsequal(t,
                            "Test_ShouldManageEdges",
                            edges,
                            []string{ "b" })
}


func Test_ShouldIgnoreSubsequentlyAddingTheSameVertex(t *testing.T) {
    dg := NewDirectedGraph()
    dg.AddVertex("a")
    dg.AddVertex("b")
    dg.AddEdge("a", "b")
    dg.AddVertex("a")
    vertices := dg.Vertices()
    edges := dg.Edges("a")
    sort.Strings(vertices)
    sort.Strings(edges)
    expectarraystringsequal(t,
                            "Test_ShouldIgnoreSubsequentlyAddingTheSameVertex",
                            vertices,
                            []string{ "a", "b" })
    expectarraystringsequal(t,
                            "Test_ShouldIgnoreSubsequentlyAddingTheSameVertex",
                            edges,
                            []string{ "b" })
}


func Test_ShouldTopologicallySortAnEmptyGraph(t *testing.T) {
    dg := NewDirectedGraph()
    expectarraystringsequal(t,
                            "Test_ShouldTopologicallySortAnEmptyGraph",
                            dg.TopologicalSort(),
                            []string{})
}


func Test_ShouldTopologicallySortASimpleAcyclicGraph(t *testing.T) {
    dg := NewDirectedGraph()
    dg.AddVertex("a")
    dg.AddVertex("b")
    dg.AddVertex("c")
    dg.AddVertex("d")
    dg.AddVertex("e")
    dg.AddEdge("a", "b")
    dg.AddEdge("b", "c")
    dg.AddEdge("c", "d")
    dg.AddEdge("d", "e")
    expectarraystringsequal(t,
                            "Test_ShouldTopologicallySortASimpleAcyclicGraph",
                            dg.TopologicalSort(),
                            []string{ "e", "d", "c", "b", "a" })
}


func Test_ShouldNotDetectCyclesInAnEmptyGraph(t *testing.T) {
    dg := NewDirectedGraph()
    expectarraystringsequal(t,
                            "Test_ShouldNotDetectCyclesInAnEmptyGraph",
                            dg.Cycle(),
                            []string{})
}


func Test_ShouldNotDetectCyclesWhenTheGraphIsACyclic(t *testing.T) {
    dg := NewDirectedGraph()
    dg.AddVertex("a")
    dg.AddVertex("b")
    dg.AddVertex("c")
    dg.AddVertex("d")
    dg.AddVertex("e")
    dg.AddEdge("a", "b")
    dg.AddEdge("b", "c")
    dg.AddEdge("c", "d")
    dg.AddEdge("d", "e")
    expectarraystringsequal(t,
                            "Test_ShouldNotDetectCyclesWhenTheGraphIsACyclic",
                            dg.Cycle(),
                            []string{})
}


func Test_ShouldDetectCyclesFromTheLastVertexToTheFirstVertex(t *testing.T) {
    dg := NewDirectedGraph()
    dg.AddVertex("a")
    dg.AddVertex("b")
    dg.AddVertex("c")
    dg.AddVertex("d")
    dg.AddVertex("e")
    dg.AddEdge("a", "b")
    dg.AddEdge("b", "c")
    dg.AddEdge("c", "d")
    dg.AddEdge("d", "e")
    dg.AddEdge("e", "a")
    expectarraystringsnotequal(t,
                               "Test_ShouldDetectCyclesFromTheLastVertexToTheFirstVertex",
                               dg.Cycle(),
                               []string{})
}


func Test_ShouldDetectCyclesFromWithinTheVertexChain(t *testing.T) {
    dg := NewDirectedGraph()
    dg.AddVertex("a")
    dg.AddVertex("b")
    dg.AddVertex("c")
    dg.AddVertex("d")
    dg.AddVertex("e")
    dg.AddEdge("a", "b")
    dg.AddEdge("b", "c")
    dg.AddEdge("c", "d")
    dg.AddEdge("d", "c")
    dg.AddEdge("d", "e")
    dg.AddEdge("e", "a")
    expectarraystringsnotequal(t,
                               "Test_ShouldDetectCyclesFromWithinTheVertexChain",
                               dg.Cycle(),
                               []string{})
}
