package main

import (
	"fmt"
	"sort"
)

// A compact, runnable tutorial demonstrating common usages of the
// Go sort package. Run with `go run main.go`.

func main() {
	fmt.Println("Go sort package tutorial — quick examples")
	fmt.Println("------------------------------------------")

	basicSort()
	fmt.Println()

	reverseSort()
	fmt.Println()

	customSort()
	fmt.Println()

	searching()
}

func basicSort() {
	fmt.Println("1) Basic sorting")
	ints := []int{3, 1, 4, 1, 5, 9, 2, 6}
	strings := []string{"zebra", "apple", "banana", "cherry"}

	fmt.Println("Before sort - ints:", ints)
	sort.Ints(ints)
	fmt.Println("After sort - ints:", ints)

	fmt.Println("Before sort - strings:", strings)
	sort.Strings(strings)
	fmt.Println("After sort - strings:", strings)
}

func reverseSort() {
	fmt.Println("2) Reverse sorting")
	ints := []int{3, 1, 4, 1, 5, 9, 2, 6}

	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("Reverse sorted ints:", ints)
}

func customSort() {
	fmt.Println("3) Custom sorting")
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}

	// Sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("Sorted by age:")
	for _, p := range people {
		fmt.Printf("  %s: %d\n", p.Name, p.Age)
	}

	// Sort by name
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})

	fmt.Println("Sorted by name:")
	for _, p := range people {
		fmt.Printf("  %s: %d\n", p.Name, p.Age)
	}
}

func searching() {
	fmt.Println("4) Searching sorted slices")
	ints := []int{1, 2, 3, 4, 5, 6, 9}

	// Search returns the index where to insert the value
	idx := sort.SearchInts(ints, 7)
	fmt.Println("Index to insert 7:", idx)

	// Check if exists
	if idx < len(ints) && ints[idx] == 7 {
		fmt.Println("7 found at index", idx)
	} else {
		fmt.Println("7 not found, would insert at", idx)
	}

	// Custom search
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	searchAge := 25
	idx2 := sort.Search(len(people), func(i int) bool {
		return people[i].Age >= searchAge
	})

	if idx2 < len(people) && people[idx2].Age == searchAge {
		fmt.Printf("Person with age %d: %s\n", searchAge, people[idx2].Name)
	}
}