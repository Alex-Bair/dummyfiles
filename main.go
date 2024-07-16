package main

import (
	"os"
	"path/filepath"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var fileContent = []byte("Howdy, world!")
const topDirectory = "tmp"
// Change directoryCount for however many nested directories you need.
const directoryCount = 5

// createNestedDirectory builds a single nested directory with the specified depth at the passed in path. It creates a single file.tmp at each level of nesting.
func createNestedDirectory(depth int, currentPath string) {
	// create directory
	dirErr := os.MkdirAll(currentPath, os.ModePerm)
	check(dirErr)

	// create file in current directory
	fileName := filepath.Join(currentPath, "file.txt")
	fileErr := os.WriteFile(fileName, fileContent, os.ModePerm)
	check(fileErr)

	// recurse to next depth
	nextDepth := depth - 1
	if nextDepth > 0 {
		nextPath := filepath.Join(currentPath, strconv.Itoa(nextDepth))
		createNestedDirectory(nextDepth, nextPath)
	}
}

// createNestedDirectories creates `count` number of nested directories at the passed in path
func createNestedDirectories(count int, currentPath string) {
	// remove directory if it already exists
	os.RemoveAll(currentPath)

	// create top level directory
	dirErr := os.MkdirAll(currentPath, os.ModePerm)
	check(dirErr)

	// create nested directories within top level
	for i := 1; i <= count; i++ {
		nextPath := filepath.Join(currentPath, strconv.Itoa(i))
		createNestedDirectory(i, nextPath)
	}
}

func main() {
	createNestedDirectories(directoryCount, topDirectory)
}