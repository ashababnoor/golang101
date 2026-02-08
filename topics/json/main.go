package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// A compact, runnable tutorial demonstrating common usages of the
// Go encoding/json package. Run with `go run main.go`.

func main() {
	fmt.Println("Go encoding/json package tutorial — quick examples")
	fmt.Println("---------------------------------------------------")

	basicMarshalUnmarshal()
	fmt.Println()

	structTags()
	fmt.Println()

	streaming()
	fmt.Println()

	customMarshaling()
}

func basicMarshalUnmarshal() {
	fmt.Println("1) Basic marshal and unmarshal")
	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	u := User{ID: 42, Name: "Ada Lovelace"}

	// Marshal to JSON
	b, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON output:")
	fmt.Println(string(b))

	// Unmarshal back
	var u2 User
	if err := json.Unmarshal(b, &u2); err != nil {
		panic(err)
	}
	fmt.Printf("Decoded struct: %+v\n", u2)
}

func structTags() {
	fmt.Println("2) Struct tags and options")
	type Person struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email,omitempty"`
		Internal string `json:"-"`
	}

	p1 := Person{ID: 1, Name: "Alice", Email: "alice@example.com", Internal: "secret"}
	p2 := Person{ID: 2, Name: "Bob"}

	b1, _ := json.MarshalIndent(p1, "", "  ")
	b2, _ := json.MarshalIndent(p2, "", "  ")

	fmt.Println("With email:")
	fmt.Println(string(b1))
	fmt.Println("Without email (omitted):")
	fmt.Println(string(b2))
}

func streaming() {
	fmt.Println("3) Streaming with Encoder/Decoder")
	type Item struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	items := []Item{
		{"first", 1},
		{"second", 2},
		{"third", 3},
	}

	// Encode to stdout
	enc := json.NewEncoder(os.Stdout)
	fmt.Println("Streaming encode:")
	for _, item := range items {
		enc.Encode(item)
	}
}

func customMarshaling() {
	fmt.Println("4) Custom marshaling")
	ct := CustomTime{Time: "2023-03-18T15:30:00Z"}
	b, _ := json.Marshal(ct)
	fmt.Println("Custom marshaled:", string(b))
}

// CustomTime demonstrates custom JSON marshaling
type CustomTime struct {
	Time string
}

// MarshalJSON implements json.Marshaler
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"custom_time": ct.Time})
}

