package main

import (
	"fmt"
)

func main() {
	csvRecords := readCsvFile("../candy-bars.csv")
	fmt.Printf("found %d records\n",  len(csvRecords))
	candyBars := fromCsv(csvRecords)
	fmt.Println(candyBars)
}