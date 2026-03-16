package main

import (
	"fmt"

	"github.com/mattschofield/go-knapsack"
)

// Set this to the number of slots available in the given year.
var capacity int64 = 1000

var items = []item{
	// Add each item (candidate) here.
}

type item struct {
	name   string
	value  int64
	weight int64
}

func (i item) Name() string {
	return i.name
}

func (i item) Value() int64 {
	return i.value
}

func (i item) Weight() int64 {
	return i.weight
}

func main() {
	var packables []knapsack.Packable
	for _, item := range items {
		packables = append(packables, item)
	}

	picks := knapsack.Knapsack(packables, capacity)
	printSummary(items, picks, capacity)
}

func printSummary(items []item, picks []int64, capacity int64) {
	fmt.Printf("Selected the following names (capacity: %d):", capacity)
	for i, pick := range picks {
		fmt.Printf("%d. Name: %s\n   Weight: %d\n", i+1, items[pick].Name(), items[pick].Weight())
	}
}
