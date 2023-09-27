package main

import "fmt"

type Folder struct {
	name     string
	children []Inode
}

func (f *Folder) print(indent string) {
	fmt.Println(indent + f.name)
	for _, child := range f.children {
		child.print(indent + indent)
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildern []Inode
	for _, child := range f.children {
		copy := child.clone()
		tempChildern = append(tempChildern, copy)
	}
	cloneFolder.children = tempChildern
	return cloneFolder
}
