package src

var (
	count int
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
	pathToSource  string
	AmountOfNodes int // TODO: can be temp
}

// createNode create a node in the graph
func (g *Graph) createNode(position int) {
	idx := position
	// To support negated nodes
	if position < 0 {
		idx = (-1)*position + g.AmountOfNodes
	}
	if g.Nodes[idx] != nil {
		return
	}

	g.Nodes[idx] = new(Node)
	// Same label as its position
	g.Nodes[idx].Label = position
}

// CreateGraph create the graph from a list of pairs of nodes
func (g *Graph) CreateGraph(pairOfNodesList []intTuple, reverse bool) {
	var nodeFrom *Node
	for _, pair := range pairOfNodesList {
		from, to := pair[0], pair[1]
		if reverse {
			from, to = pair[1], pair[0]
		}

		// Create the nodes in the graph if they do not exist
		g.createNode(from)
		nodeFrom = g.Nodes[g.Get(from)]
		g.createNode(to)

		// Append neighbor
		nodeFrom.Neighbors = append(nodeFrom.Neighbors, to)
	}
}

// Get the index of an element by the state of connection
func (g *Graph) Get(label int) int {
	if label < 0 {
		return (-1)*label + g.AmountOfNodes
	}
	return label
}

func (g *Graph) dfsVisit(index int, rollback, reverse bool) {
	if index == 0 {
		//if reverse {
		//	// Insert path length into the max SCCs array
		//	insertPathLength(count)
		//}
		return
	}

	node := g.Nodes[index]
	if !rollback {
		g.Nodes[index].Visited = true

		// TODO: delete
		if reverse {
			count++
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
func (g *Graph) Dfs(reverse bool) {
	for idx, node := range g.Nodes {
		if node == nil {
			continue
		}

		if !node.Visited {
			g.dfsVisit(idx, false, reverse)
		}
	}
}

func (g *Graph) IsSatisfiable(edgesList []intTuple, amountOfNodes int) bool {
	// Run DFS
	g.Dfs(false)

	// Create reversed graph G
	reversedGraph := new(Graph)
	reversedGraph.Nodes = make([]*Node, 2*amountOfNodes+1)
	reversedGraph.AmountOfNodes = g.AmountOfNodes
	reversedGraph.CreateGraph(edgesList, false)

	// Pop one by one to make DFS to the top of stack
	for !stack.IsEmpty() {
		v, _ := stack.Pop()
		if !reversedGraph.Nodes[v].Visited {
			count = 0
			reversedGraph.dfsVisit(v, false, true)
		}
	}

	return false
}
