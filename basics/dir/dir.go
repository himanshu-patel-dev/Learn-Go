package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dirPath := "dir/subdir"
	// clean the path and remove any dir present to that
	// Mkdir do not give error of existing dir
	os.RemoveAll(dirPath)

	err := os.Mkdir(dirPath, 0755)
	check(err)
	// os.RemoveAll will delete a whole directory tree (similarly to rm -rf)
	// defer os.RemoveAll(dirPath) // comment out to see dir and files created

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}
	filePath := "dir/subdir/file"
	createEmptyFile(filePath)

	// We can create a hierarchy of directories, including parents with MkdirAll.
	// This is similar to the command-line mkdir -p
	dirPath2 := "dir/subdir/parent/child"
	err = os.MkdirAll(dirPath2, 0755)
	check(err)

	createEmptyFile("dir/subdir/parent/file2")
	createEmptyFile("dir/subdir/parent/file3")
	createEmptyFile("dir/subdir/parent/child/file4")

	// ReadDir lists directory contents, returning a slice of
	// os.DirEntry objects
	c, err := os.ReadDir("dir/subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	// Listing subdir/parent
	// child true
	// file2 false
	// file3 false

	// Chdir lets us change the current working directory, similarly to cd
	err = os.Chdir("dir/subdir/parent/child")
	check(err)
	// now we are in dir/subdir/parent/child directory

	c, err = os.ReadDir(".")
	check(err)
	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	// Listing subdir/parent/child
	//   file4 false

	// cd back to where we started
	err = os.Chdir("../../../../")
	check(err)

	// We can also visit a directory recursively, including all its
	// sub-directories. Walk accepts a callback function to handle every
	// file or directory visited.
	fmt.Println("Visiting dir")
	err = filepath.Walk("dir", visit)
	check(err)
	// Visiting dir
	// dir true
	// dir/dir.go false
	// dir/subdir true
	// dir/subdir/file false
	// dir/subdir/parent true
	// dir/subdir/parent/child true
	// dir/subdir/parent/child/file4 false
	// dir/subdir/parent/file2 false
	// dir/subdir/parent/file3 false
}

// visit is called for every file or directory found recursively
// by filepath.Walk.
func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Printf("%v -- %v\n", path, info.IsDir())
	return nil
}

// $ go run dir/dir.go

// Before execution

// dir/
// └── dir.go

// After execution

// dir/
// ├── dir.go
// └── subdir
//     ├── file
//     └── parent
//         ├── child
//         │   └── file4
//         ├── file2
//         └── file3
