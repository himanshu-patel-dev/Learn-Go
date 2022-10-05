# Basic Go data types

## `for` loop

```go
package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println() // i not accessible here

	i := 1
	for ; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println(i)
}
```
```bash
1 2 3 4 5 6 7 8 9 
1 2 3 4 5 6 7 8 9 10
```
In this case, `i` is a local and temporary variable, which means that after the termination of the `for` loop, `i` will be garbage collected at some point and disappear. However, if `i` was defined outside the `for` loop, it would have kept its value after the termination of the for loop. In that case, the value of `i` after the termination of the `for` loop would have been `10` .

```go
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		}
		if i == 5 {
			fmt.Println("Breaking out of loop")
			break // break here
		}
		fmt.Println("The value of i is", i)
	}
	fmt.Println("Exiting program")
}
```
```
The value of i is 0
The value of i is 1
The value of i is 3
The value of i is 4
Breaking out of loop
Exiting program
```

## `while` loop
Go does not offer the `while` keyword for writing `while` loops, but it allows you to use a `for` loop instead of a `while` loop.
```go
package main

import "fmt"

func main() {
	// infinite loop (if no break condition)

	// for {
	// 	fmt.Println(1)
	// }

	// while loop
	i := 1
	for {
		fmt.Print(i, " ")
		if i == 5 {
			break
		}
		i++
	}
}
```
```
1 2 3 4 5
```

## `range` keyword
```go
package main

import "fmt"

func main() {
	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Println("index:", i, "value: ", value)
	}

	fmt.Println("--------------")
	
	for _, value := range anArray {
		fmt.Println("value: ", value)
	}
	
	fmt.Println("--------------")
	
	for index := range anArray {
		fmt.Println("index: ", index)
	}
}
```
```
index: 0 value:  0
index: 1 value:  1
index: 2 value:  -1
index: 3 value:  2
index: 4 value:  -2
--------------
value:  0
value:  1
value:  -1
value:  2
value:  -2
--------------
index:  0
index:  1
index:  2
index:  3
index:  4
```

## `arrays`
```go
package main

import (
	"fmt"
)

func main() {
	anArray := [5]int{0, 1, -1, 2, -2}
	fmt.Println(anArray)      // [0 1 -1 2 -2]
	fmt.Println(len(anArray)) // 5

	twoD := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(twoD) // [[1 2 3 4] [5 6 7 8] [9 10 11 12]]

	threeD := [3][2][1]int{ // [[[1] [2]] [[3] [4]] [[5] [6]]]
		{
			{1}, {2},
		}, {
			{3}, {4},
		}, {
			{5}, {6},
		},
	}
	fmt.Println(threeD)

	// printing 3D array
	fmt.Println("----------------")
	for i := 0; i < len(threeD); i++ {
		for j := 0; j < len(threeD[i]); j++ {
			for k := 0; k < len(threeD[i][j]); k++ {
				fmt.Print(threeD[i][j][k], " ")
			}
		}
	}
	fmt.Println("\n----------------")
	for _, v := range threeD {
		for _, m := range v {
			for _, s := range m {
				fmt.Print(s, " ")
			}
		}
	}
	fmt.Println("\n----------------")
}
```
```
[0 1 -1 2 -2]
5
[[1 2 3 4] [5 6 7 8] [9 10 11 12]]
[[[1] [2]] [[3] [4]] [[5] [6]]]
----------------
1 2 3 4 5 6 
----------------
1 2 3 4 5 6 
----------------
```

### The shortcomings of Go arrays
- Once you define an array, you cannot change its size, which means that Go arrays are not dynamic. Putting it simply, if you need to add an element to an existing array that has no space left, you will need to create a bigger array and copy all of the elements of the old array to the new one. 
- Second, when you pass an array to a function as a parameter, you actually pass a copy of the array, which means that any changes you make to an array inside a function will be lost after the function exits. 
- Last, passing a large array to a function can be pretty slow, mostly because Go has to create a copy of the array. The solution to all of these problems is to use Go `slices`.

## `slices`

As `slices` are passed by reference to functions, which means that what is actually passed is the memory address of the slice variable, any modifications that you make to a slice inside a function will not get lost after the function exits. Additionally, passing a big slice to a function is significantly faster than passing an array with the same number of elements because Go will not have to make a copy of the slice—it will just pass the memory address of the slice variable.

`slice` literals are defined just like arrays but without the element count
```go
slice := []int{1, 2, 3}
fmt.Println(slice)	// [1 2 3]
```

Also the `make()` function that allows you to create empty slices with the desired `length` and `capacity` as the parameters passed to `make()`. The `capacity` parameter can be omitted. In that case, the `capacity` of the slice will be the same as its `length`.

Define a new empty slice with 20 places that can be automatically expanded when needed as follows:
```go
integer := make([]int, 20)
```
Please note that Go automatically initializes the elements of an empty slice to the zero value of its type, which means that the value of the initialization depends on the type of the object stored in the slice.

```go
package main

import "fmt"

func printSlice(slice []int) {
	fmt.Println("----------------")
	fmt.Println(slice)
	fmt.Printf("%T\n", slice)
	fmt.Println(len(slice), cap(slice))
	fmt.Println("----------------")
}

func main() {
	// slice literals are defined just like arrays
	// but without the element count
	slice := []int{1, 2, 3}
	printSlice(slice)

	integer := make([]int, 5, 10)
	printSlice(integer)

	for _, v := range integer {
		fmt.Print(v)
	}
	fmt.Println()

	// append to slice, len will inc but cap will remain same until
	// len become more than capacity
	integer = append(integer, 1, 2, 3, 4, 5)
	printSlice(integer) // 10 10

	// if len further inc then capacity will become double of prev
	integer = append(integer, 11)
	printSlice(integer) // 11 20
}
```
```
----------------
[1 2 3]
[]int
3 3
----------------
----------------
[0 0 0 0 0]
[]int
5 10
----------------
00000
----------------
[0 0 0 0 0 1 2 3 4 5]
[]int
10 10
----------------
----------------
[0 0 0 0 0 1 2 3 4 5 11]
[]int
11 20
----------------
```

- `Pointer`: The pointer is used to points to the first element of the array that is accessible through the slice. Here, it is not necessary that the pointed element is the first element of the array.
- `Length`: The length is the total number of elements present in the array.
- `Capacity`: The capacity represents the maximum size upto which it can expand.

### empty a slice
```go
package main

import "fmt"

func printSlice(slice []int) {
	fmt.Println("----------------")
	fmt.Println(slice)
	fmt.Printf("%T\n", slice)
	fmt.Println(len(slice), cap(slice))
	fmt.Println("----------------")
}

func main() {
	newSlice := []int{1, 2, 3, 4, 5}
	printSlice(newSlice)

	// keep allocated memory - slice the array to zero length
	// but the capacity and type remians the same
	newSlice = newSlice[:0]
	printSlice(newSlice)

	// if slice is extended the original data reappears
	newSlice = newSlice[:2]
	printSlice(newSlice)

	// To remove all elements, simply set the slice to nil.
	// This will release the underlying array to the garbage collector
	// (assuming there are no other references). This is also the recommended way in Go Lang
	newSlice = nil
	printSlice(newSlice)
}
```
```
----------------
[1 2 3 4 5]
[]int
5 5
----------------
----------------
[]
[]int
0 5
----------------
----------------
[1 2]
[]int
2 5
----------------
----------------
[]
[]int
0 0
----------------
```

Additionally, you can use `[:]` notation for creating a new slice from an existing slice or array. Unlike python here the second index (3 here) can't be more than length of array else we get an error.
```go
s2 := integer[1:3]	// include all element starting from index 1 till index 2
```
Please note that this process is called re-slicing, and it can cause problems in some cases. The changes done to new slice are also applied to parent slice.
```go
package main

import "fmt"

func main() {
	s1 := make([]int, 5)
	reSlice := s1[1:3]
	fmt.Println(s1)
	fmt.Println(reSlice)
	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println(s1)
	fmt.Println(reSlice)
}
```
```
[0 0 0 0 0]
[0 0]
[0 -100 123456 0 0]
[-100 123456]
```
Because both slices point to the same memory address! Put simply, the re-slice process does not make a copy of the original slice.

The second problem of re-slicing is that, even if you re-slice a slice in order to use a small part of the original slice, the underlying array from the original slice will be kept in memory for as long as the smaller re-slice exists because the original slice is being referenced by the smaller re-slice. Although this is not truly important for small slices, it can cause problems when you are reading big files into slices and you only want to use a small part of them.

## `copy()` function

You should be very careful when using the `copy()` function on slices because the built-in `copy(dst, src)` copies the minimum of the `len(dst)` and `len(src)` elements.

```go
package main

import (
	"fmt"
)

func main() {

	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)

	copy(a6, a4) // dest src
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	fmt.Println()

	b6 := []int{-10, 1, 2, 3, 4, 5}
	b4 := []int{-1, -2, -3, -4}
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)
	copy(b4, b6)
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)

	fmt.Println()
	array4 := [4]int{4, -4, 4, -4}
	s6 := []int{1, 1, -1, -1, 5, -5}
	copy(s6, array4[1:])
	fmt.Println("array4:", array4[:])
	fmt.Println("s6:", s6)

	fmt.Println()

	array5 := [5]int{5, -5, 5, -5, 5}
	s7 := []int{7, 7, -7, -7, 7, -7, 7}
	copy(array5[1:], s7)
	fmt.Println("array5:", array5)
	fmt.Println("s7:", s7)
}
```
```
a6: [-10 1 2 3 4 5]
a4: [-1 -2 -3 -4]
a6: [-1 -2 -3 -4 4 5]
a4: [-1 -2 -3 -4]

b6: [-10 1 2 3 4 5]
b4: [-1 -2 -3 -4]
b6: [-10 1 2 3 4 5]
b4: [-10 1 2 3]

array4: [4 -4 4 -4]
s6: [-4 4 -4 -1 5 -5]

array5: [5 7 7 -7 -7]
s7: [7 7 -7 -7 7 -7 7]
```

As `copy()` only accepts slice arguments, you should also use the `[:]` notation to convert the array into a slice. If you try to copy an array into a slice or vice versa without using the `[:]` notation, the program will fail to compile and will display one of the following error messages:
```
# command-line-arguments
./a.go:42:6: first argument to copy should be slice; have [5]int
./a.go:43:6: second argument to copy should be slice or string; have
[5]int
./a.go:44:6: arguments to copy must be slices; have [5]int, [5]int
```

## `sort.slice()`
```go
package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	// sorting string array
	strs := []string{"c", "b", "a"}
	sort.Strings(strs)
	fmt.Println("String: ", strs)

	// sorting int array
	ints := []int{4, 2, 1}
	sort.Ints(ints)
	fmt.Println(ints)

	// sorting custom array
	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})
	mySlice = append(mySlice, aStructure{"Bill", 134, 45})
	mySlice = append(mySlice, aStructure{"Marietta", 155, 45})
	mySlice = append(mySlice, aStructure{"Epifanios", 144, 50})
	mySlice = append(mySlice, aStructure{"Athina", 134, 40})
	fmt.Println("0:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("<:", mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Println(">:", mySlice)
}
```
```
String:  [a b c]
[1 2 4]
0: [{Mihalis 180 90} {Bill 134 45} {Marietta 155 45} {Epifanios 144 50} {Athina 134 40}]
<: [{Bill 134 45} {Athina 134 40} {Epifanios 144 50} {Marietta 155 45} {Mihalis 180 90}]
>: [{Mihalis 180 90} {Marietta 155 45} {Epifanios 144 50} {Bill 134 45} {Athina 134 40}]
```

## maps
```go
package main

import "fmt"

func main() {
	// define map using make
	firstMap := make(map[string]int)
	firstMap["first"] = 1
	firstMap["second"] = 2
	fmt.Println("Map: ", firstMap) // Map:  map[first:1 second:2]

	// define map using map literal
	secondMap := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(secondMap) // map[a:1 b:2]

	// try deleting one element
	delete(secondMap, "a")
	fmt.Println(secondMap) // map[b:2]
	// try deleting again
	delete(secondMap, "a")
	fmt.Println(secondMap) // map[b:2]

	// iterate over all elements of map
	for key, value := range firstMap {
		fmt.Println(key, value)
	}

	// checking if a element exists in a map
	// event if the key don't exists we get the zero value of
	// data type define for values in map
	ele, ok := firstMap["third"]
	fmt.Println("Element exists: ", ele, ok) // 0 false
}
```
```
Map:  map[first:1 second:2]
map[a:1 b:2]
map[b:2]
map[b:2]
first 1
second 2
Element exists:  0 false
```

### Storing to a nil map
```go
package main

import "fmt"

func main() {
	// works fine
	aMap := map[string]int{}
	aMap["test"] = 1

	// making a nil map
	bMap := map[string]int{}
	bMap = nil
	fmt.Println(bMap)

	// lookup do not give any error, similarly range and len
	ele, ok := bMap["key"]
	fmt.Println("Lookup: ", ele, ok)

	// inserting in nil map gives error
	bMap["test"] = 1
}
```
```
map[]
Lookup:  0 false
panic: assignment to entry in nil map

goroutine 1 [running]:
main.main()
        /home/himanshu/HP/dev/Learn-Go/go_data_types/failMap.go:20 +0x1b4
exit status 2
```

```go
	// how a nil map get into picture
	var cMap map[string]int
	fmt.Println(cMap) // map[]
	if cMap == nil {  // nil Map
		fmt.Println("nil Map")
	} else {
		fmt.Println("Not nil Map")
	}
```

## `constants`
constants are usually global variables. Thus, you might rethink your approach if you find yourself defining too many
constant variables with a local scope!

The main benefit that you get from using constants in your programs is the guarantee that their value will not change during program execution!

Strictly speaking, the value of a constant variable is defined at compile time-not at run time. Behind the scenes, Go uses Boolean, string, or number as the type for storing a constant variable.

