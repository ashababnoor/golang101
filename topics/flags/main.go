package main

import (
	"flag"
	"fmt"
	"os"
)

// A compact, runnable tutorial demonstrating common usages of the
// Go flag package. Run with `go run main.go -name=Gopher -n 3 -v`

func main() {
	fmt.Println("Go flag package tutorial — quick examples")
	fmt.Println("------------------------------------------")

	basicFlags()
	fmt.Println()

	differentTypes()
	fmt.Println()

	customUsage()
	fmt.Println()

	subcommands()
}

func basicFlags() {
	fmt.Println("1) Basic flag usage")
	name := flag.String("name", "world", "name to greet")
	n := flag.Int("n", 1, "number of greetings")
	v := flag.Bool("v", false, "verbose output")
	flag.Parse()

	for i := 0; i < *n; i++ {
		if *v {
			fmt.Printf("(%d/%d) Hello %s\n", i+1, *n, *name)
		} else {
			fmt.Printf("Hello %s\n", *name)
		}
	}
}

func differentTypes() {
	fmt.Println("2) Different flag types")
	var (
		str    = flag.String("string", "default", "a string flag")
		num    = flag.Int("int", 42, "an int flag")
		float  = flag.Float64("float", 3.14, "a float flag")
		bool   = flag.Bool("bool", false, "a bool flag")
	)

	flag.Parse()
	fmt.Printf("String: %s, Int: %d, Float: %f, Bool: %t\n", *str, *num, *float, *bool)
}

func customUsage() {
	fmt.Println("3) Custom usage function")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Custom help message\n\n")
		flag.PrintDefaults()
	}

	help := flag.Bool("help", false, "show help")
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	fmt.Println("Run with -help to see custom usage")
}

func subcommands() {
	fmt.Println("4) Subcommands example")
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <command> [args...]")
		fmt.Println("Commands: greet, count")
		return
	}

	switch os.Args[1] {
	case "greet":
		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
		name := greetCmd.String("name", "world", "name to greet")
		greetCmd.Parse(os.Args[2:])
		fmt.Printf("Hello, %s!\n", *name)

	case "count":
		countCmd := flag.NewFlagSet("count", flag.ExitOnError)
		upto := countCmd.Int("upto", 5, "count up to this number")
		countCmd.Parse(os.Args[2:])
		for i := 1; i <= *upto; i++ {
			fmt.Println(i)
		}

	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
	}
}
