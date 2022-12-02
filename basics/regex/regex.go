package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match) // true

	emailMatch, _ := regexp.MatchString("([a-zA-Z]+)@([a-zA-Z]{3,5}).([in|com|us])", "first@gmail.com")
	fmt.Println(emailMatch) // true

	// make a regex compiler which we can use each time with new string
	r, _ := regexp.Compile("p([a-z]+)ch")
	// match given string with r regex compiler
	fmt.Println(r.MatchString("peach")) // true

	// peach - grabbing the first element in given string which matches patterns
	fmt.Println(r.FindString("peach punch"))

	// get the index of matched word we got in FindString
	fmt.Println("idx:", r.FindStringIndex("peach punch")) // idx: [0 5]

	// gives the matching string and matching sub string in reg exp
	fmt.Println(r.FindStringSubmatch("peach punch")) // [peach ea]

	// get index of submatch
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // [0 5 1 3]

	// get all string that matches the pattern
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	// [peach punch pinch]

	// get sub string match index for all matches
	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	// all: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	// limit the number of matches
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	// [peach punch]

	// provide []byte arguments instead of string
	fmt.Println(r.Match([]byte("peach")))
	// true

	// MustCompile panics instead of returning an error,
	// which makes it safer to use for global variables.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)
	// regexp: p([a-z]+)ch

	// return copy of first string by replacing the matched sub string with
	// given second string
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	// a <fruit>

	// The Func variant allows you to transform matched text with a given function
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
	// a PEACH
}
