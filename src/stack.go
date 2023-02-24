package src

import "fmt"

type Stack struct {
	items []int
}

// Push a new value onto the stack
func (stack *Stack) Push(num int) {
	stack.items = append(stack.items, num)
}

// IsEmpty check if stack is empty
func (stack *Stack) IsEmpty() bool {
	return len(stack.items) == 0
}

// Pop remove and return top element of stack. Return false if stack is empty.
func (stack *Stack) Pop() (int, bool) {
	if stack.IsEmpty() {
		return -1, false
	}

	// The element to extract will be the final one
	indexToExtract := len(stack.items) - 1
	element := stack.items[indexToExtract]

	// Remove the element from the stack
	stack.items = stack.items[:indexToExtract]

	return element, true
}

// Print the items in stack
func (stack *Stack) Print() {
	fmt.Println(stack.items)
}

// Clear the stack
func (stack *Stack) Clear() {
	stack.items = nil
}
