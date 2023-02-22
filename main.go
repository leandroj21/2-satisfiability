package main

import (
	"2satisfiability/src"
	"fmt"
	"strconv"
)

func isSatisfiable(path string) bool {
	// Create graph
	var graph src.Graph
	graph.CreateGraph(path)
	return false
}

func main() {
	pathBase := "./data/2sat"
	fmt.Println("Satisfiability of the 6 files:")
	for i := 1; i < 7; i++ {
		if isSatisfiable(pathBase + strconv.Itoa(i) + ".txt") {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
	}
}
