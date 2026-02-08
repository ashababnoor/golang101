package main

import (
	"fmt"
	"strings"
)

// A compact, runnable tutorial demonstrating common usages of the
// Go strings package. Run with `go run main.go`.

func main() {
	fmt.Println("Go strings package tutorial — quick examples")
	fmt.Println("---------------------------------------------")

	basicOperations()
	fmt.Println()

	searching()
	fmt.Println()

	modification()
	fmt.Println()

	splittingJoining()
}

func basicOperations() {
	fmt.Println("1) Basic operations")
	s := "Hello, World!"

	fmt.Println("Contains 'World':", strings.Contains(s, "World"))
	fmt.Println("HasPrefix 'Hello':", strings.HasPrefix(s, "Hello"))
	fmt.Println("HasSuffix '!'", strings.HasSuffix(s, "!"))
	fmt.Println("Index of 'World':", strings.Index(s, "World"))
	fmt.Println("Count of 'l':", strings.Count(s, "l"))
}

func searching() {
	fmt.Println("2) Searching and comparison")
	s1 := "golang"
	s2 := "GOLANG"

	fmt.Println("EqualFold:", strings.EqualFold(s1, s2))
	fmt.Println("Compare:", strings.Compare(s1, s2))

	// Case insensitive search
	fmt.Println("Contains (case insensitive):", strings.Contains(strings.ToLower(s1), strings.ToLower("LANG")))
}

func modification() {
	fmt.Println("3) Modification")
	s := "  hello world  "

	fmt.Println("TrimSpace:", "'"+strings.TrimSpace(s)+"'")
	fmt.Println("ToUpper:", strings.ToUpper(s))
	fmt.Println("ToLower:", strings.ToLower(s))
	fmt.Println("Replace:", strings.Replace(s, "world", "golang", 1))
	fmt.Println("ReplaceAll:", strings.ReplaceAll(s, "l", "L"))
}

func splittingJoining() {
	fmt.Println("4) Splitting and joining")
	s := "apple,banana,cherry,date"

	parts := strings.Split(s, ",")
	fmt.Println("Split:", parts)

	joined := strings.Join(parts, " | ")
	fmt.Println("Join:", joined)

	// Fields (splits on whitespace)
	fields := strings.Fields("hello   world\tgolang")
	fmt.Println("Fields:", fields)
}