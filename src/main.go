package main

import (
	"fmt"
)

func main() {
	records := readCsvFile("../candy-bars.csv")
	fmt.Println(records)
}