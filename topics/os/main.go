package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// A compact, runnable tutorial demonstrating common usages of the
// Go os package. Run with `go run main.go`.

func main() {
	fmt.Println("Go os package tutorial — quick examples")
	fmt.Println("----------------------------------------")

	environment()
	fmt.Println()

	fileOperations()
	fmt.Println()

	directoryOperations()
	fmt.Println()

	processInfo()
}

func environment() {
	fmt.Println("1) Environment variables")
	fmt.Println("HOME:", os.Getenv("HOME"))
	fmt.Println("PATH exists:", os.Getenv("PATH") != "")

	// Set and get
	os.Setenv("MY_VAR", "hello")
	fmt.Println("MY_VAR:", os.Getenv("MY_VAR"))

	// All env vars
	fmt.Println("Total env vars:", len(os.Environ()))
}

func fileOperations() {
	fmt.Println("2) File operations")
	// Create temp file
	tmpFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Write to file
	tmpFile.WriteString("Hello from os package!")
	tmpFile.Sync() // flush to disk

	// Get file info
	info, err := tmpFile.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("File size: %d bytes\n", info.Size())
	fmt.Printf("File mode: %s\n", info.Mode())
}

func directoryOperations() {
	fmt.Println("3) Directory operations")
	// Create temp dir
	tmpDir, err := os.MkdirTemp("", "example-dir-")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create subdirectory
	subdir := filepath.Join(tmpDir, "subdir")
	os.Mkdir(subdir, 0o755)

	// List directory
	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		panic(err)
	}
	fmt.Println("Directory contents:")
	for _, entry := range entries {
		fmt.Printf("  %s (dir: %t)\n", entry.Name(), entry.IsDir())
	}

	// Change directory
	oldWd, _ := os.Getwd()
	defer os.Chdir(oldWd)
	os.Chdir(tmpDir)
	fmt.Println("Changed to temp dir, current wd:", tmpDir)
}

func processInfo() {
	fmt.Println("4) Process information")
	fmt.Println("PID:", os.Getpid())
	fmt.Println("PPID:", os.Getppid())

	// Hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Hostname error:", err)
	} else {
		fmt.Println("Hostname:", hostname)
	}

	// Working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Getwd error:", err)
	} else {
		fmt.Println("Working directory:", wd)
	}
}