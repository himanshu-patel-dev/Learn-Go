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
	// The easiest way to create a temporary file is by calling
	// os.CreateTemp. It creates a file and opens it for reading
	// and writing. We provide "" as the first argument, so
	// os.CreateTemp will create the file in the default location for our OS.
	f, err := os.CreateTemp("", "sample")
	check(err)
	fmt.Println("Temp file name:", f.Name())
	// Temp file name: /tmp/sample428631613

	// Display the name of the temporary file. On Unix-based OSes the
	// directory will likely be /tmp. The file name starts with the
	// prefix given as the second argument to os.CreateTemp and the
	// rest is chosen automatically to ensure that concurrent calls
	// will always create different file names

	// Clean up the file after we’re done. The OS is likely to clean
	// up temporary files by itself after some time, but it’s good
	// practice to do this explicitly
	defer os.Remove(f.Name())

	// write some more data
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// If we intend to write many temporary files, we may prefer
	// to create a temporary directory. os.MkdirTemp’s arguments
	// are the same as CreateTemp’s, but it returns a directory
	// name rather than an open file.
	dname, err := os.MkdirTemp("", "tempdir")
	check(err)
	fmt.Println("Temp dir name:", dname)
	// Temp dir name: /tmp/sampledir2045362464
	// delete the temp dir after we are done
	defer os.RemoveAll(dname)

	// synthesize temporary file names by prefixing them with
	// our temporary directory
	fname := filepath.Join(dname, "file1")
	fmt.Println("File path inside temp dir: ", fname)
	// File path inside temp dir:  /tmp/tempdir3191488193/file1
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
