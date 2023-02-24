package main

import (
	"2satisfiability/src"
	"fmt"
	"strconv"
)

func isSatisfiable(path string) bool {
	amountOfNodes, edgesList := src.ReadFile(path)

	// Create graph G_rev first
	graph := new(src.Graph)
	graph.AmountOfNodes = amountOfNodes
	graph.Nodes = make([]*src.Node, 2*amountOfNodes+1)
	graph.CreateGraph(edgesList, true)

	return graph.IsSatisfiable(edgesList, amountOfNodes)
}

func main() {
	pathBase := "./data/2sat"
	fmt.Println("Satisfiability of the 6 files:")
	for i := 1; i <= 6; i++ {
		if isSatisfiable(pathBase + strconv.Itoa(i) + ".txt") {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
	}
}
