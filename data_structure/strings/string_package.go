package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("------------ Len ------------")
	fmt.Println(len("abcd")) // 4

	fmt.Println("------------ Clone String ------------")
	// clone a string
	sampleString := "HelloWorld"
	cloneString := strings.Clone(sampleString) + " #"
	fmt.Println(cloneString) // HelloWorld #

	fmt.Println("------------ Compare Strings ------------")
	// compare strings
	a := "abc"
	b := "abc"
	if a == b {
		fmt.Println("a = b")
	}
	a = "abc"
	b = "cba"
	if a < b {
		fmt.Println("a < b")
	}
	if b > a {
		fmt.Println("b > a")
	}
	if a != b {
		fmt.Println("a != b")
	}

	fmt.Println("------------ Contains ------------")
	// str containes a subStr completely - Contains(str, subStr)
	fmt.Println(strings.Contains("seafood", "food")) // true
	fmt.Println(strings.Contains("seafood", ""))     // true
	fmt.Println(strings.Contains("", ""))            // true

	fmt.Println("------------ ContainsAny ------------")
	// str contains any Unicode code points in chars are within s.
	fmt.Println(strings.ContainsAny("team", "i"))  // false - i not in team
	fmt.Println(strings.ContainsAny("team", "az")) // true - a in team
	fmt.Println(strings.ContainsAny("team", ""))   // false - no char to match
	fmt.Println(strings.ContainsAny("", ""))       // false - no match

	fmt.Println("------------ ContainsRune ------------")
	// str contains any Unicode code points r
	fmt.Println(strings.ContainsRune("team", 97))        // true - code point for "a" is 97
	fmt.Println(strings.ContainsRune("team", 98))        // false - code point for "b" is 98
	fmt.Println(strings.ContainsRune("team", rune('t'))) // true - code point for "t" = rune('t)
	for _, unicode := range "man" {
		isPresent := strings.ContainsRune("team", unicode)
		if isPresent {
			fmt.Println(string(unicode))
		}
	}
	// print - m and a, these runes are in "team"

	fmt.Println("------------ Count ------------")
	// counts the number of non-overlapping instances of substr in s
	// if substr is empty then it returns
	// 1 + the number of Unicode code points in s i.e. = 1+len(s)
	fmt.Println(strings.Count("Cheese", "e")) // 3
	fmt.Println(strings.Count("Five", ""))    // 5

	fmt.Println("------------ EqualFold ------------")
	// reports whether s and t are case insensitive equal
	fmt.Println(strings.EqualFold("AbCd", "abcd"))     // true
	fmt.Println(strings.EqualFold("AbC123", "aBc123")) // true

	fmt.Println("------------ FieldFunc ------------")
	// FieldsFunc splits the string s at each run of Unicode code
	// points c satisfying f(c) returns an array of slices of s

	// makes no guarantees about the order in which it calls f(c)
	// and assumes that f always returns the same value for a given c
	f := func(c rune) bool {
		return string(c) == "."
	}
	fmt.Println(strings.FieldsFunc(".first.second.third.", f)) // [first second third]
	fmt.Println(strings.FieldsFunc("first.second.third", f))   // [first second third]
	f = func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.FieldsFunc("foo1;foo2;foo3", f)) // [foo1 foo2 foo3]

	fmt.Println("------------ HasPrefix & HasSuffix ------------")
	// HasPrefix tests whether the string s begins with prefix
	fmt.Println(strings.HasPrefix("Gopher", "Go")) // true
	fmt.Println(strings.HasPrefix("Gopher", "go")) // false
	fmt.Println(strings.HasPrefix("Gopher", ""))   // true

	fmt.Println(strings.HasSuffix("Gopher", "er")) // true
	fmt.Println(strings.HasSuffix("Gopher", "eR")) // false
	fmt.Println(strings.HasSuffix("Gopher", ""))   // true

	fmt.Println("------------ Index ------------")
	// return the index of first instance of substr in s or -1
	fmt.Println(strings.Index("ShinChanShinChan", "Chan")) // 4
	fmt.Println(strings.Index("ShinChanShinChan", "Chi"))  // -1

	fmt.Println("------------ IndexRune ------------")
	// return the index of first instance of Unicode code point r or -1
	fmt.Println(strings.IndexRune("Chicken", 'k')) // 4
	fmt.Println(strings.IndexRune("Chicken", 'd')) // -1

	fmt.Println("------------ IndexFunc ------------")
	// return index in s of the first unicode code point satisfying f(c)
	f = func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))    // 7
	fmt.Println(strings.IndexFunc("Hello, world", f)) // -1

	fmt.Println("------------ Join ------------")
	// joins the string array with a seperator
	s := []string{"first", "bar", "baz"}
	fmt.Println(strings.Join(s, ","))

	fmt.Println("------------ Map ------------")
	// returns a copy of the string s with all its characters
	// modified according to the mapping function
	g := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			rem := r - 'A' - 13
			if rem < 0 {
				rem += 26
			}
			return 'A' + (rem)%26
		case r >= 'a' && r <= 'z':
			rem := r - 'a' - 13
			if rem < 0 {
				rem += 26
			}
			return 'a' + (rem)%26
		}
		return r
	}
	// You came a long way...
	fmt.Println(strings.Map(g, "Lbh pnzr n ybat jnl..."))

	fmt.Println("------------ Count ------------")
	// return a new string consisting of n copies of s
	fmt.Println(strings.Repeat("la", 2) + " Land") // lala Land

	fmt.Println("------------ Replace & ReplaceAll ------------")
	// return a copy of string s with first n non-overlapping instances
	// of old replaced by new
	fmt.Println(strings.Replace("old old old", "o", "go", 2))  // gold gold old
	fmt.Println(strings.Replace("old old old", "o", "go", -1)) // gold gold gold
	fmt.Println(strings.ReplaceAll("old old old", "o", "go"))  // gold gold gold

	fmt.Println("------------ Split ------------")
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))                 // ["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))                  // [" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal", "a ")) // ["" "man " "plan " "canal"]
	fmt.Printf("%q\n", strings.Split("", "Ronaldo"))                // [""]

	fmt.Println("------------ Lower Upper ------------")
	fmt.Println(strings.ToUpper("lowerCaSe")) // LOWERCASE
	fmt.Println(strings.ToLower("LOWercaSE")) // lowercase

	fmt.Println("------------ Trim ------------")
	// returns a slice of the string s with all leading and
	// trailing Unicode code points contained in cutset removed.
	fmt.Println(strings.Trim("!#$Hello Gopher^&", "!#$^&"))      // Hello Gopher
	fmt.Println(strings.TrimLeft("!#$Hello Gopher^&", "!#$^&"))  // Hello Gopher^&
	fmt.Println(strings.TrimRight("!#$Hello Gopher^&", "!#$^&")) // !#$Hello Gopher
	// use trim function
	fmt.Println(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})) // Hello, Gophers
	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n")) // Hello, Gophers

	// reverse a string
	runeArray := []rune("elppa")
	for i, j := 0, len(runeArray)-1; i < j; i, j = i+1, j-1 {
		runeArray[i], runeArray[j] = runeArray[j], runeArray[i]
	}
	fmt.Println(string(runeArray)) // apple
}
