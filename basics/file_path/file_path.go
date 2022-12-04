// The filepath package provides functions to parse and
// construct file paths in a way that is portable between
// operating systems

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println(p) // dir1/dir2/filename

	// In addition to providing portability, Join will also normalize
	// paths by removing superfluous separators and directory changes.

	// takes care of extra slash and putting slash as per OS
	// fwd slash in Unix and bkwd slash in Windows
	fmt.Println(filepath.Join("dir1//", "filename"))
	// dir1/filename

	// resolves the operators and give final path
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))
	// dir1/filename

	fmt.Println("Dir(p): ", filepath.Dir(p))
	// Dir(p):  dir1/dir2
	fmt.Println("Base(p): ", filepath.Base(p))
	// Base(p):  filename
	// filepath.Split() // returns both dir and file

	// check whether a path is absolute
	fmt.Println(filepath.IsAbs("dir/file"))  // false
	fmt.Println(filepath.IsAbs("/dir/file")) // true

	// look for extension in filename
	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext) // .json
	// get file name with extension removed
	fmt.Println(strings.TrimSuffix(filename, ext)) // config

	// Rel finds a relative path between a base and target
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel) // t/file

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel) // ../c/t/file
}
