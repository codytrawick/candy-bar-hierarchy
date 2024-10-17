package main

import (
	// "fmt"
)

type Node struct {
	bitmask int
	ingredients []string
	candyBars []
	children []Node
}

func compare(a int, b int) bool {
	return (a != b) && (a & b == a);
}

func constructForest(candyBars []CandyBar) []Node {
	forest := []Node {}
	for i, candyBar := range candyBars {
	  index := contains(forest, candyBar.bitmask)
		if index == -1 {
			forest = append(forest, Node{ bitmask: candyBar.bitmask, ingredients: candyBar.ingredients, candyBars: []CandyBar { candyBars[i] }})
		} else {
			// fmt.Printf("match %s\n", candyBar.name)
			match := forest[index]
			match.candyBars = append(match.candyBars, candyBars[i])
			forest[index] = match
			// fmt.Println(match)
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