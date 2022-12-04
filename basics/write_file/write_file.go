package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func get_write_file_path(write_file string) string {
	pwd, err := os.Getwd()
	check(err)
	filePath := pwd + "/write_file/" + write_file + ".txt"
	return filePath
}

func main() {
	// how to dump a string (or just bytes) into a file.
	byte_data := []byte("hello\ngo\n")
	filepath := get_write_file_path("write_file")
	err := os.WriteFile(filepath, byte_data, 0644)
	check(err)
	// in write_file.txt
	// hello
	// go

	// create file for writing and put a immediate a deferred close
	// since we give it same file name as above it will over write
	file_pointer, err := os.Create(filepath)
	check(err)
	defer file_pointer.Close()

	data := []byte{115, 111, 109, 101, 10} // some
	bytes_written, err := file_pointer.Write(data)
	fmt.Printf("wrote %d bytes\n", bytes_written)

	bytes_written, err = file_pointer.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", bytes_written)

	// commit the current content of file from memory to disk
	file_pointer.Sync()

	// bufio provides buffered writers in addition to the
	// buffered readers we saw earlier.
	buffered_writer := bufio.NewWriter(file_pointer)
	bytes_written, err = buffered_writer.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", bytes_written)

	// write any leftover data in buffer to writer/file_pointer
	buffered_writer.Flush()
}

// wrote 5 bytes
// wrote 7 bytes
// wrote 9 bytes
