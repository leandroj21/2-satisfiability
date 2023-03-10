package src

var (
	stack Stack
)

type Node struct {
	Label     int
	Visited   bool
	Previous  int
	Neighbors []int
}

type Graph struct {
	Nodes         []*Node
	AmountOfNodes int // TODO: can be temp
	dfsPath       map[int]bool
}

// createNode create a node in the graph
func (g *Graph) createNode(position int) *Node {
	idx := position
	// To support negated nodes
	if position < 0 {
		idx = (-1)*position + g.AmountOfNodes
	}
	if g.Nodes[idx] != nil {
		return g.Nodes[idx]
	}

	g.Nodes[idx] = new(Node)
	// Same label as its position
	g.Nodes[idx].Label = position
	return g.Nodes[idx]
}

// CreateGraph create the graph from a list of pairs of nodes
func (g *Graph) CreateGraph(pairOfNodesList []intTuple, reverse bool) {
	for _, variables := range pairOfNodesList {
		var1, var2 := variables[0], variables[1]
		if reverse {
			var1, var2 = -variables[0], -variables[1]
		}

		/* Create the nodes in the graph if they do not exist
		 Connections are made using disjunctive syllogism:
		-var1 -> var2
		-var2 -> var1
		*/
		from1 := g.createNode(-var1)
		g.createNode(var2)
		from2 := g.createNode(-var2)
		g.createNode(var1)

		// Append neighbor
		from1.Neighbors = append(from1.Neighbors, var2)
		from2.Neighbors = append(from2.Neighbors, var1)
	}
}

// Get the index of an element by the state of connection
func (g *Graph) Get(label int) int {
	if label < 0 {
		return (-1)*label + g.AmountOfNodes
	}
	return label
}

// dfsVisit returns true if in the path does not exist a pair [b] and [~b], false otherwise
func (g *Graph) dfsVisit(index int, rollback, reverse bool) {
	if index == 0 {
		return
	}

	node := g.Nodes[index]
	if !rollback {
		g.Nodes[index].Visited = true

		if reverse {
			g.dfsPath[node.Label] = true
		}
	}

	// Look for the next node
	allVisited := true
	for _, neighbor := range node.Neighbors {
		neighborIndex := g.Get(neighbor)
		if !g.Nodes[neighborIndex].Visited {
			allVisited = false

			// Continue to the next node
			g.Nodes[neighborIndex].Previous = index
			g.dfsVisit(neighborIndex, false, reverse)
		}
	}

	// Go back to the previous node
	if allVisited {
		if !reverse {
			// Push to stack since all neighbors were visited
			stack.Push(index)
		}
		g.dfsVisit(node.Previous, true, reverse)
	}
	return
}

// Dfs of the graph
func (g *Graph) Dfs() {
	for idx, node := range g.Nodes {
		if node == nil {
			continue
		}

		if !node.Visited {
			g.dfsVisit(idx, false, false)
		}
	}
}

// checkCollisions returns (true, varTag) if there is a contradiction between
// [b] and [~b]. (false, 0) otherwise
func (g *Graph) checkCollisions() (contradictions []int) {
	for k := range g.dfsPath {
		if _, exists := g.dfsPath[-k]; exists {
			contradictions = append(contradictions, Abs(k))
		}
	}
	return contradictions
}

// IsSatisfiable returns (true, 0) if the clauses are satisfiable. Otherwise
// (false, varTag), where varTag is the variable with a contradiction
func (g *Graph) IsSatisfiable(edgesList []intTuple, amountOfNodes int) (contradictions []int) {
	// Run DFS
	g.Dfs()

	// Create reversed graph G
	reversedGraph := new(Graph)
	reversedGraph.Nodes = make([]*Node, 2*amountOfNodes+1)
	reversedGraph.AmountOfNodes = g.AmountOfNodes
	reversedGraph.CreateGraph(edgesList, false)

	// Pop one by one to make DFS to the top of stack
	for !stack.IsEmpty() {
		v, _ := stack.Pop()
		if !reversedGraph.Nodes[v].Visited {
			reversedGraph.dfsPath = make(map[int]bool)
			reversedGraph.dfsVisit(v, false, true)

			if contradictionsFound := reversedGraph.checkCollisions(); len(contradictionsFound) > 0 {
				contradictions = append(contradictions, contradictionsFound...)
			}
		}
	}

	return
}
