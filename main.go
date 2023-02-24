package main

import (
	"2satisfiability/src"
	"fmt"
	"strconv"
)

func isSatisfiable(path string) bool {
	amountOfNodes, edgesList, lineCount := src.ReadFile(path)

	// Create graph G_rev first
	graph := new(src.Graph)
	graph.AmountOfNodes = lineCount
	graph.Nodes = make([]*src.Node, 2*amountOfNodes+1)
	graph.CreateGraph(edgesList, true)

	return graph.IsSatisfiable(edgesList, amountOfNodes)
}

func main() {
	pathBase := "./data/2sat"
	fmt.Println("Satisfiability of the 6 files:")
	// TODO: change 2 to 7
	for i := 1; i < 2; i++ {
		if isSatisfiable(pathBase + strconv.Itoa(i) + ".txt") {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
	}
}
