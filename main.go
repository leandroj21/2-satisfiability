package main

import (
	"2satisfiability/src"
	"fmt"
	"strconv"
)

func isSatisfiable(path string) (bool, int) {
	amountOfNodes, edgesList := src.ReadFile(path)

	// Create graph G_rev first
	graph := new(src.Graph)
	graph.AmountOfNodes = amountOfNodes
	graph.Nodes = make([]*src.Node, 2*amountOfNodes+1)
	graph.CreateGraph(edgesList, true)

	return graph.IsSatisfiable(edgesList, amountOfNodes)
}

func printErrors(errors map[int]int) {
	fmt.Println("\nFails:")
	for testcase, varTag := range errors {
		fmt.Printf("2-sat%d (bit %d) failed due to a contradiction of: %d y %d\n",
			testcase,
			testcase,
			varTag,
			-varTag,
		)
	}
}

func main() {
	// errors[num_of_testcase] = variable that caused error
	errors := make(map[int]int)
	pathBase := "./data/2sat"
	fmt.Println("Satisfiability of the 6 testcases:")
	for i := 1; i <= 6; i++ {
		if satisfiable, varTagError := isSatisfiable(pathBase + strconv.Itoa(i) + ".txt"); satisfiable {
			fmt.Print(1)
		} else {
			errors[i] = varTagError
			fmt.Print(0)
		}
	}
	printErrors(errors)
}
