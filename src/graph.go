package src

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	Label int
	// true = not negated, false = negated
	State bool
}

type Node struct {
	Label     int
	Visited   bool
	Previous  int
	Neighbors []Edge
}

type Graph struct {
	NodesNotNegated []*Node
	NodesNegated    []*Node
}

// Get an element by the state of connection
func (g *Graph) Get(label int, state bool) *Node {
	if state {
		return g.NodesNotNegated[label]
	}
	return g.NodesNegated[label]
}

func (g *Graph) Insert(node *Node, state bool) {
	if g.Get(node.Label, state) == nil {
		if state {
			g.NodesNotNegated[node.Label] = node
		} else {
			g.NodesNegated[node.Label] = node
		}
	}
}

// CreateGraph from a file
func (g *Graph) CreateGraph(path string) {
	// Open file
	file, err := os.Open(path)
	CheckError(err)
	defer func(file *os.File) {
		err := file.Close()
		CheckError(err)
	}(file)

	amountOfNodes := 0
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	_, _ = fmt.Sscanf(line, "%d", &amountOfNodes)

	// Create slices
	g.NodesNegated = make([]*Node, amountOfNodes+1, amountOfNodes+1)
	g.NodesNotNegated = make([]*Node, amountOfNodes+1, amountOfNodes+1)

	for scanner.Scan() {
		lineStr := scanner.Text()
		connections := strings.Fields(lineStr)

		from, _ := strconv.Atoi(connections[0])
		to, _ := strconv.Atoi(connections[1])

		nodeFrom := Node{Label: Abs(from)}
		nodeTo := Node{Label: Abs(to)}
		edge := Edge{Label: to, State: to > 0}
		nodeFrom.Neighbors = append(nodeFrom.Neighbors, edge)

		g.Insert(&nodeFrom, from > 0)
		g.Insert(&nodeTo, to > 0)
	}
}
