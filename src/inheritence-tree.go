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

func lessThan(a int, b int) bool {
	// fmt.Printf("%x & %x %d\n", a, b, a & b)
	return (a != b) && (a & b == a);
}

func constructForest(candyBars []CandyBar) []Node {
	forest := consolidateEquivalentNodes(candyBars)

	forest = pruneForest(forest)
	for i, node := range forest {
		if len(node.children) > 0 {
			forest[i].children = pruneForest(node.children)
		}
	}
	return forest
}

func pruneForest(forest []Node) []Node {
	searchIndex := 0
	for true {
		if searchIndex == len(forest) {
			break;
		}
		first := forest[searchIndex]
		found := false
		for i, node := range forest {
			if lessThan(node.bitmask, first.bitmask) {
				match := forest[i]
				match.children = append(match.children, first)
				forest[i] = match
				if searchIndex == 0 {
					forest = forest[1:]
				} else {
					forest = append(forest[:searchIndex], forest[searchIndex+1:]...)
				}
				found = true
				break;
			}
		}
		if found {
			searchIndex = 0
		} else {
			searchIndex = searchIndex + 1
		}
	}

	return forest
}

func consolidateEquivalentNodes(candyBars []CandyBar) []Node {
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

func printForest(forest []Node) {
	for _, node := range forest {
		printNode(node, 0)
	}
}

func printNode(node Node, depth int) {
	fmt.Printf("%s- %s with %d ingredients\n", strings.Repeat(" ", depth), strings.Join(node.candyBars[:], ", "), len(node.ingredients))
	for _, child := range node.children {
		printNode(child, depth + 2)
	}
}