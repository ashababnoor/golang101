package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// A compact, runnable tutorial demonstrating common usages of the
// Go file system operations. Run with `go run main.go`.

func main() {
	fmt.Println("Go file operations tutorial — quick examples")
	fmt.Println("---------------------------------------------")

	basicReadWrite()
	fmt.Println()

	directoryOperations()
	fmt.Println()

	filepathExamples()
	fmt.Println()

	tempFiles()
}

func basicReadWrite() {
	fmt.Println("1) Basic read/write files")
	dir := filepath.Join(".", "tmp_demo")
	os.MkdirAll(dir, 0o755)
	defer os.RemoveAll(dir)

	file := filepath.Join(dir, "example.txt")
	content := []byte("Hello file system!\n")

	// Write file
	if err := os.WriteFile(file, content, 0o644); err != nil {
		panic(err)
	}

	// Read file
	b, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote then read: %s", string(b))

	// Stat and info
	fi, _ := os.Stat(file)
	fmt.Println("Size:", fi.Size())
}

func directoryOperations() {
	fmt.Println("2) Directory operations")
	dir := filepath.Join(".", "tmp_demo2")
	os.MkdirAll(dir, 0o755)
	defer os.RemoveAll(dir)

	subdir := filepath.Join(dir, "subdir")
	os.Mkdir(subdir, 0o755)

	file := filepath.Join(subdir, "file.txt")
	os.WriteFile(file, []byte("content"), 0o644)

	// List directory
	entries, _ := os.ReadDir(dir)
	for _, entry := range entries {
		fmt.Println("Entry:", entry.Name(), "IsDir:", entry.IsDir())
	}
}

func filepathExamples() {
	fmt.Println("3) Filepath operations")
	path := "/home/user/docs/file.txt"

	fmt.Println("Base:", filepath.Base(path))
	fmt.Println("Dir:", filepath.Dir(path))
	fmt.Println("Ext:", filepath.Ext(path))

	// Join paths
	full := filepath.Join("home", "user", "file.txt")
	fmt.Println("Joined:", full)

	// Clean path
	dirty := "/home/user/../user/./file.txt"
	clean := filepath.Clean(dirty)
	fmt.Println("Cleaned:", clean)
}

func tempFiles() {
	fmt.Println("4) Temporary files and directories")
	// Create temp file
	tmpFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	tmpFile.WriteString("temporary content")
	fmt.Println("Temp file:", tmpFile.Name())

	// Create temp dir
	tmpDir, err := os.MkdirTemp("", "example-dir-")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpDir)
	fmt.Println("Temp dir:", tmpDir)
}