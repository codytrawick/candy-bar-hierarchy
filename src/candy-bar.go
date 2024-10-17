package main

import "fmt"

type CandyBar struct {
	name string
	ingredients []string
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
		candyBars = append(candyBars, CandyBar{ name: v[0], ingredients: getIngredients(v, headers) })
		// fmt.Printf("Record %d %s\n", i, v[0])
	}
	return candyBars
}