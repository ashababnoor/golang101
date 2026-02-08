package main

import (
	"fmt"
	"regexp"
)

// A compact, runnable tutorial demonstrating common usages of the
// Go regexp package. Run with `go run main.go`.

func main() {
	fmt.Println("Go regexp package tutorial — quick examples")
	fmt.Println("--------------------------------------------")

	basicMatching()
	fmt.Println()

	submatches()
	fmt.Println()

	replaceExamples()
	fmt.Println()

	compileAndReuse()
}

func basicMatching() {
	fmt.Println("1) Basic matching")
	re := regexp.MustCompile(`\b\w+@\w+\.\w+\b`)
	s := "Contact: alice@example.com or bob@company.org"

	if re.MatchString(s) {
		fmt.Println("Found email in string")
	}

	// Find first match
	match := re.FindString(s)
	fmt.Println("First match:", match)
}

func submatches() {
	fmt.Println("2) Submatches and groups")
	re := regexp.MustCompile(`\b([A-Za-z]+)@(\w+\.\w+)\b`)
	s := "Contact: alice@example.com or bob@company.org"

	matches := re.FindAllStringSubmatch(s, -1)
	for _, m := range matches {
		fmt.Printf("Full: %s, User: %s, Domain: %s\n", m[0], m[1], m[2])
	}
}

func replaceExamples() {
	fmt.Println("3) Replace operations")
	re := regexp.MustCompile(`\b\d{3}-\d{3}-\d{4}\b`)
	s := "Call me at 123-456-7890 or 098-765-4321"

	// Replace with placeholder
	result := re.ReplaceAllString(s, "[PHONE]")
	fmt.Println("Replaced:", result)

	// Replace with function
	result2 := re.ReplaceAllStringFunc(s, func(match string) string {
		return "(" + match + ")"
	})
	fmt.Println("Wrapped:", result2)
}

func compileAndReuse() {
	fmt.Println("4) Compile once, reuse")
	re := regexp.MustCompile(`\b\w+\b`)

	words := []string{"hello", "world", "golang"}
	for _, word := range words {
		if re.MatchString(word) {
			fmt.Println("Word:", word)
		}
	}

	// Split using regexp
	re2 := regexp.MustCompile(`\s+`)
	parts := re2.Split("hello   world\tgolang", -1)
	fmt.Println("Split:", parts)
}