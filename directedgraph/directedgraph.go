// 'directedgraph.go'.
// Chris Shiels.


package directedgraph


import (
)


type DirectedGraph struct {
    graph map[string]map[string]bool
}


func NewDirectedGraph() *DirectedGraph {
    return &DirectedGraph{ graph: make(map[string]map[string]bool) }
}


func (d *DirectedGraph) AddVertex(name string) {
    if d.graph[name] == nil {
        d.graph[name] = make(map[string]bool)
    }
}


func (d *DirectedGraph) Vertices() (vertices []string) {
    vertices = make([]string, len(d.graph))
    i := 0
    for key := range d.graph {
        vertices[i] = key
        i++
    }
    return vertices
}


func (d *DirectedGraph) AddEdge(namefrom, nameto string) {
    d.graph[namefrom][nameto] = true
}


func (d *DirectedGraph) Edges(name string) (edges []string) {
    edges = make([]string, len(d.graph[name]))
    i := 0
    for key := range d.graph[name] {
        edges[i] = key
        i++
    }
    return edges
}


func (d *DirectedGraph) TopologicalSort() (vertices []string) {

    var recurse func(name string, stack []string) []string
    recurse = func(name string, stack []string) []string {
        for _, v := range stack {
            if name == v {
                return stack
            }
        }
        for vertex := range d.graph[name] {
            stack = recurse(vertex, stack)
        }
        stack = append(stack, name)
        return stack
    }

    stack := make([]string, 0)
    for vertex := range d.graph {
        stack = recurse(vertex, stack)
    }
    vertices = stack
    return vertices
}


func (d *DirectedGraph) Cycle() (vertices []string) {

    var recurse func(name string, stack []string) ([]string, bool)
    recurse = func(name string, stack []string) ([]string, bool) {
        var cycle bool
        stack = append(stack, name)
        for _, v := range stack[0:len(stack) - 1] {
            if name == v {
                return stack, true
            }
        }
        for vertex := range d.graph[name] {
            stack, cycle = recurse(vertex, stack)
            if cycle {
                return stack, cycle
            }
        }
        stack = stack[0:len(stack) - 1]
        return stack, false
    }

    stack := make([]string, 0)
    var cycle bool
    for vertex := range d.graph {
        stack, cycle = recurse(vertex, stack)
        if cycle {
            break
        }
    }
    vertices = stack
    return vertices
}
