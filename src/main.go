package main

import (
	"fmt"
)

func main() {
	csvRecords := readCsvFile("../candy-bars.csv")
	fmt.Printf("found %d records\n",  len(csvRecords))
	candyBars := fromCsv(csvRecords)
	forest := constructForest(candyBars)
	printForest(forest)
	// fmt.Println(forest)
}