package main

/*

use command : `go run .`
not `go run main.go`

-------------------- DEFINITION: Prototype Pattern

Prototype is a creational design pattern that allows cloning
objects, even complex ones, without coupling to their specific classes.

-------------------- APPLICATION

- Use the Prototype pattern when your code shouldn’t depend 
on the concrete classes of objects that you need to copy.

-------------------- IMPLEMENTATION
- Create the prototype interface and declare the clone method in it. 
Or just add the method to all classes of an existing class hierarchy, 
if you have one.

- Every class must explicitly override the cloning method and use its own 
class name along with the new operator. Otherwise, the cloning method may 
produce an object of a parent class.

*/

import "fmt"

func main() {
	file1 := File{name: "File1"}
	file2 := File{name: "File2"}
	file3 := File{name: "File3"}

	folder1 := Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}
	folder2 := Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}

// Printing hierarchy for Folder2
//   Folder2
//     Folder1
//         File1
//     File2
//     File3

// Printing hierarchy for clone Folder
//   Folder2_clone
//     Folder1_clone
//         File1_clone
//     File2_clone
//     File3_clone
