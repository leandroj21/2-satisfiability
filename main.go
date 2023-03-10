package main

import (
	"2satisfiability/src"
	"fmt"
	"strconv"
	"time"
)

// isSatisfiable creates the graph and check if it is 2-satisfiable
func isSatisfiable(path string) []int {
	amountOfNodes, edgesList := src.ReadFile(path)

	// Create graph G_rev first
	graph := new(src.Graph)
	graph.AmountOfNodes = amountOfNodes
	graph.Nodes = make([]*src.Node, 2*amountOfNodes+1)
	graph.CreateGraph(edgesList, true)

	return graph.IsSatisfiable(edgesList, amountOfNodes)
}

func printContradictions(errors [6][]int) {
	fmt.Println("\n\nContradictions:")
	for testcase, tags := range errors {
		// Ignore satisfiable problems
		if len(tags) == 0 {
			continue
		}

		fmt.Printf("\t2-sat%d.txt (bit %d) failed due to a contradiction of: \n", testcase+1, testcase+1)
		printed := make(map[int]bool)
		for _, tag := range tags {
			if _, wasPrinted := printed[src.Abs(tag)]; wasPrinted {
				continue
			}

			printed[src.Abs(tag)] = true
			fmt.Printf("\t\t%d and %d\n", tag, -tag)
		}
	}
}

func printTiming(errors []time.Duration, start time.Time) {
	fmt.Println("\nTiming:")
	for testcase, totalTime := range errors {
		fmt.Printf("   Testing of bit %d took %v\n", testcase+1, totalTime)
	}
	fmt.Printf("   Finish time: %v\n", time.Since(start))
}

func main() {
	start := time.Now()
	// contradictions[num_of_testcase] = variable that caused error
	var contradictions [6][]int
	var timing []time.Duration

	fmt.Println("Satisfiability of the 6 testcases:")
	pathBase := "./data/2sat"
	for i := 1; i <= 6; i++ {
		startLoop := time.Now()
		if contradictionsList := isSatisfiable(pathBase + strconv.Itoa(i) + ".txt"); len(contradictionsList) == 0 {
			fmt.Print(1)
		} else {
			contradictions[i-1] = contradictionsList
			fmt.Print(0)
		}
		timing = append(timing, time.Since(startLoop))
	}

	printContradictions(contradictions)
	printTiming(timing, start)
}
