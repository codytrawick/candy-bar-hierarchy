package main

import (
	"fmt"
	"strings"
)

type Node struct {
	bitmask int
	ingredients []string
	candyBars []string
	children []Node
}

func compare(a int, b int) bool {
	return (a != b) && (a & b == a);
}

func constructForest(candyBars []CandyBar) []Node {
	forest := []Node {}
	// First pass to consolidate nodes
	for i, candyBar := range candyBars {
	  index := contains(forest, candyBar.bitmask)
		if index == -1 {
			forest = append(forest, Node{ bitmask: candyBar.bitmask, ingredients: candyBar.ingredients, candyBars: []string { candyBars[i].name }})
		} else {
			match := forest[index]
			match.candyBars = append(match.candyBars, candyBars[i].name)
			forest[index] = match
		}
	}
	return forest
}

func contains(nodes []Node, element int) int {
	for i, node := range nodes {
		if node.bitmask == element {
			return i
		}
	}
	return -1
}

func printNode(node Node) {
	fmt.Printf("- %s and %d children\n", strings.Join(node.candyBars[:], ", "), len(node.children))
}