package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func get_read_file_path() string {
	pwd, err := os.Getwd()
	checkError(err)
	filePath := pwd + "/read_file/read_file.txt"
	return filePath
}

func main() {
	// Perhaps the most basic file reading task is slurping
	// a file’s entire contents into memory.
	filepath := get_read_file_path()
	data, err := os.ReadFile(filepath)
	checkError(err)
	fmt.Println(string(data))

	// to control what parts of a file are read.
	// open a file to obtain an os.File value.
	file_pointer, err := os.Open(filepath)
	checkError(err)

	// Read some bytes from the beginning of the file.
	// Allow up to 5 to be read but also note how many
	// actually were read.
	byte_chunk := make([]byte, 5)
	bytes_read, err := file_pointer.Read(byte_chunk)
	checkError(err)
	fmt.Printf("\n%d Bytes: %s\n", bytes_read, byte_chunk[:bytes_read])

	// You can also Seek to a known location in the file and
	// Read from there. 0 - means relative to starting of file
	read_position, err := file_pointer.Seek(6, 0)
	checkError(err)
	byte_chunk = make([]byte, 2)
	// read enough bytes which are less than or equal to the size
	// of output byte array
	bytes_read, err = file_pointer.Read(byte_chunk)
	checkError(err)
	fmt.Printf("\n%d Bytes: %s - Pointer : %d\n",
		bytes_read, byte_chunk[:bytes_read], read_position)

	// There is no built-in rewind, but Seek(0, 0) accomplishes this
	_, err = file_pointer.Seek(0, 0)
	checkError(err)

	// The bufio package implements a buffered reader that may be useful
	// both for its efficiency with many small reads and because of the
	// additional reading methods it provides
	read_pointer := bufio.NewReader(file_pointer)
	// return n bytes from file without moving the file pointer
	read_bytes, err := read_pointer.Peek(5)
	checkError(err)
	fmt.Printf("\n%d Bytes: %s\n", 5, read_bytes)

	// close the file pointer in the end
	file_pointer.Close()
}

// abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRST

// 5 Bytes: abcde

// 2 Bytes: gh - Pointer : 6

// 5 Bytes: abcde
