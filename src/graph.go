package src

type Edge struct {
	Label int
	// true = not negated, false = negated
	State bool
}

type Node struct {
	Label     int
	Visited   bool
	Previous  int
	Neighbors []*Edge
}

type Graph struct {
	NodesNotNegated []*Node
	NodesNegated    []*Node
}

// Get an element by the state of connection
func (g *Graph) Get(edge Edge) *Node {
	if edge.State {
		return g.NodesNotNegated[edge.Label]
	}
	return g.NodesNegated[edge.Label]
}
