package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"os"
	"slices"
	"strings"
	"text/tabwriter"

	"github.com/mattschofield/go-knapsack"
)

// Set this to the number of slots available in the given year.
var capacity int64 = 1000

var items = []item{
	// TODO
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
	seen := make(map[string]struct{})
	for _, item := range items {
		if _, alreadySeen := seen[item.Name()]; alreadySeen {
			// If there are duplicate names (which is possible)
			// manually append some differentiable suffix to each duplicate
			// name, to aid identification on selection.
			log.Fatalf("Name %q already exists: please append some unique suffix to each duplicate name\n", item.Name())
		}
		seen[item.Name()] = struct{}{}
	}

	selected := make(map[string]int)
	for {
		picks := permuteAndSelect(items)
		fingerprint := fingerprintFrom(picks)
		selected[fingerprint] += 1
		if selected[fingerprint] == 3 {
			fmt.Printf("[%d] %s - winner!\n", selected[fingerprint], fingerprint)
			// Stop when the first fingerprint to be selected 3 times is found.
			printSummary(picks)
			break
		} else {
			fmt.Printf("[%d] %s - try again\n", selected[fingerprint], fingerprint)
		}
	}
}

func permuteAndSelect(items []item) []item {
	perm := make([]item, len(items))
	for i, v := range rand.Perm(len(items)) {
		perm[v] = items[i]
	}

	var packables []knapsack.Packable
	for _, item := range perm {
		packables = append(packables, item)
	}

	picks := knapsack.Knapsack(packables, capacity)

	var pickedItems []item
	for _, i := range picks {
		pickedItems = append(pickedItems, perm[i])
	}
	return pickedItems
}

func fingerprintFrom(picks []item) string {
	var selection []string
	for _, pick := range picks {
		selection = append(selection, pick.Name())
	}
	slices.Sort(selection)
	hasher := md5.New()
	hasher.Write([]byte(strings.Join(selection, "")))
	return hex.EncodeToString(hasher.Sum(nil))
}

func printSummary(picks []item) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Printf("\nSelected the following names (capacity: %d):\n", capacity)
	for i, pick := range picks {
		fmt.Fprintf(w, "%d. \tName: %s\tWeight: %d\n", i+1, pick.Name(), pick.Weight())
	}
	w.Flush()
}
