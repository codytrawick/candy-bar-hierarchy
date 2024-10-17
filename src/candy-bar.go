package main

import (
	"fmt"
	"slices"
)

type CandyBar struct {
	name string
	ingredients []string
	bitmask int
}

func getBitmask(candyBar CandyBar, headers []string) int {
	bitmask := 0
	for i, v := range headers {
		if slices.Contains(candyBar.ingredients, v) {
			bitmask = bitmask | (1 << i - 1)
		}
	}
	return bitmask
}

func getIngredients(record []string, headers []string) []string {
	ingredients := []string {}
	for i, ingredient := range record {
		if ingredient == "1" {
			ingredients = append(ingredients, headers[i])
		}
	}
	return ingredients
}

func fromCsv(records [][]string) []CandyBar {
	headers := records[0]
	fmt.Println(headers)
	candyBars := []CandyBar {}
	for _, v := range records[1:] {
		candyBar := CandyBar{ name: v[0], ingredients: getIngredients(v, headers) }
		candyBar.bitmask = getBitmask(candyBar, headers)
		candyBars = append(candyBars, candyBar)
	}
	return candyBars
}