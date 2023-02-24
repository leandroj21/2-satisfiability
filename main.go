package main

import (
	"2satisfiability/src"
	"fmt"
	"strconv"
	"time"
)

// isSatisfiable creates the graph and check if it is 2-satisfiable
func isSatisfiable(path string) (bool, int) {
	amountOfNodes, edgesList := src.ReadFile(path)

	// Create graph G_rev first
	graph := new(src.Graph)
	graph.AmountOfNodes = amountOfNodes
	graph.Nodes = make([]*src.Node, 2*amountOfNodes+1)
	graph.CreateGraph(edgesList, true)

	return graph.IsSatisfiable(edgesList, amountOfNodes)
}

func printContradictions(errors map[int]int) {
	fmt.Println("\n\nContradictions:")
	for testcase, varTag := range errors {
		fmt.Printf("2-sat%d (bit %d) failed due to a contradiction of: %d and %d\n",
			testcase,
			testcase,
			varTag,
			-varTag,
		)
	}
}

func printTiming(errors map[int]time.Duration, start time.Time) {
	fmt.Println("\nTiming:")
	for testcase, totalTime := range errors {
		fmt.Printf("   Testing of bit %d took %v\n", testcase, totalTime)
	}
	fmt.Printf("   Finish time: %v\n", time.Since(start))
}

func main() {
	start := time.Now()
	// contradictions[num_of_testcase] = variable that caused error
	contradictions := make(map[int]int)
	timing := make(map[int]time.Duration)

	fmt.Println("Satisfiability of the 6 testcases:")
	pathBase := "./data/2sat"
	for i := 1; i <= 6; i++ {
		startLoop := time.Now()
		if satisfiable, varTagError := isSatisfiable(pathBase + strconv.Itoa(i) + ".txt"); satisfiable {
			fmt.Print(1)
		} else {
			contradictions[i] = varTagError
			fmt.Print(0)
		}
		timing[i] = time.Since(startLoop)
	}

	printContradictions(contradictions)
	printTiming(timing, start)
}
