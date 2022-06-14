# GO and OS

```bash
$ go version
go version go1.18.3 linux/amd64
```

## Some benefits of GO
- Go has support for procedural, concurrent, and distributed programming.
- Go supports **Garbage Collection**, so you do not have to deal with memory allocation and deallocation.
- Go does not have a **preprocessor**. It does high-speed compilation. As a consequence, Go can also be used as a scripting language.
- Go uses static linking by default, which means that the binary files produced can be easily transferred to other machines with the same OS. As a consequence, once a Go program is compiled successfully and an executable file is generated, the developer does not need to worry about libraries, dependencies, and different library versions anymore.

## Disadvantages of GO
- Go does not have direct support for object-oriented programming (OOP), which can be a problem for programmers who are used to writing code in an object-oriented manner. Nevertheless, you can use composition in Go to mimic
inheritance.

## Compiling Go code
Go does not care about the name of the source file of an autonomous program as long as the package name is `main` and there is a single `main()` function in it, because the `main()` function is where the program execution begins. As a result, you cannot have multiple `main()` functions in the files of a single project.

```go
// aSourceFile.go
package main

import "fmt"

func main() {
	fmt.Println("Hi Go Language")
}
```
```bash
$ go run aSourceFile.go 
Hi Go Language
```
So, in order to compile aSourceFile.go and create a **statically linked executable file**, you will need to execute this command:
```
$ go build aSourceFile.go

$ ll -h
-rwxrwxr-x 1 himanshu himanshu 1.7M Jun 11 08:54 aSourceFile*
-rw-rw-r-- 1 himanshu himanshu   75 Jun 11 08:52 aSourceFile.go

$ ./aSourceFile 
Hi Go Language
```
The main reason that aSourceFile is that big is because it is **statically linked**, which means that it does not require any external libraries in order to run.

## Executing Go code
There is another way to execute your Go code that does not create any permanent executable files—it just generates some intermediate files that are automatically deleted afterwards. Which allows you to use Go as if it were a
scripting language like Python, Ruby, and Perl.
```
$ go run aSourceFile.go
Hi Go Language
```

## You either use a Go package or do not include it
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello there!")
}
```
```
$ go run packageNotUsed.go 
# command-line-arguments
./packageNotUsed.go:5:2: imported and not used: "os"
```
If you remove the `os` package from the import list of the program, `packageNotUsed.go` will compile just fine

### Way to bypass
Using an underscore character in front of a package name in the import list will not create an error message in the compilation process, even if that package is not used in the program:
```go
package main

import (
	"fmt"
	_ "os"
)

func main() {
	fmt.Println("Hello there!")
}
```
```
$ go run packageNotUsedUnderscore.go 
Hello there!
```

## There is only one way to format curly braces
```go
package main

import (
	"fmt"
)

func main()
{
fmt.Println("Go has strict rules for curly braces!")
}
```
```
$ go run curly.go 
# command-line-arguments
./curly.go:8:1: syntax error: unexpected semicolon or newline before {
```
The official explanation for this error message is that Go requires the use of semicolons as statement terminators in many contexts, and the compiler automatically inserts the required semicolons when it thinks that they are necessary. Therefore, putting the opening brace `{` in its own line will make the Go compiler insert a semicolon at the end of the previous line `func main()`, which produces the error message.

## Creating modules and packages and difference in them
A module has a number of Go code files implementing the functionality of a package, but it also has two additional and important files in the root: the `go.mod` file and the `go.sum`.
```
└── projects
    └── mymodule
        └── go.mod
```
To progress create a dir `go_project` and then `mymodule`, move in `mymodule` and create a `go.mod` file and `go.sum` file with this command.
```bash
$ go mod init mymodule
```
This create a `go.mod` file
```go
module mymodule

go 1.18
```
### Adding Go Code to Your Module
```
└── projects
    └── mymodule
        └── go.mod
        └── main.go
```
Create a `main.go` file adjacent to `go.mod`. This file need not to be names `main.go` it can be named anything (The `main.go` file is commonly used in Go programs to signal the starting point of a program) as the method `main()` present in this file is the starting point of execution but matching the two makes it easier to find. 

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Modules!")
}
```
```
$ go run main.go 
Hello, Modules!
```
### Adding a Package to Your Module
- A module may contain any number of packages and sub-packages, or it may contain none at all.
- Create a package named `mypackage` inside the `mymodule` directory.
```
└── projects
    └── mymodule
        └── mypackage
			└── mypackage.go
        └── main.go
        └── go.mod
```
```go
// mypackage.go
package mypackage

import "fmt"

func PrintHello() {
	fmt.Println("Hello, Modules! This is mypackage speaking!")
}
```
- Since you want the PrintHello function to be available from another package, the capital P in the function name is important. To understand more on [package visibility](https://www.digitalocean.com/community/tutorials/understanding-package-visibility-in-go).
- Now that you’ve created the `mypackage` package with an exported function, you will need to import it from the `mymodule` package to use it. 
```bash
package main

import (
	"fmt"

	"mymodule/mypackage"
)

func main() {
	fmt.Println("Hello, Modules!")

	mypackage.PrintHello()
}
```
```
$ go run main.go 
Hello, Modules!
Hello, Modules! This is mypackage speaking!
```
- New import begins with `mymodule`, which is the same module name you set in the `go.mod` file. This is followed by the path separator and the package you want to import, `mypackage` in this case.

### Declare GOPATH and add to PATH
```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

<!-- bin is the place where go install all it's packages and libraries -->
```
After doing this only we can install additional packages from internet.

### Adding a Remote Module as a Dependency
- You use the repository’s path as a way to reference the module you’d like to use
- Run go get and provide the module you’d like to add. In this case, you’ll get `github.com/spf13/cobra`.
```
$ go get github.com/spf13/cobra
```
- Installing this create a additional file `go.sum` and update the existing `go.mod`.
- Update your `main.go` file in the mymodule directory with the Cobra code below to use the new dependency
```go
package main

import (
	"fmt"
    
	"github.com/spf13/cobra"

	"mymodule/mypackage"
)

func main() {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Modules!")

			mypackage.PrintHello()
		},
	}

	fmt.Println("Calling cmd.Execute()!")
	cmd.Execute()
}
```
```
$ go run main.go 
Calling cmd.Execute()!
Hello, Modules!
Hello, Modules! This is mypackage speaking!
```

## Downloading Go packages
```go
// getPackage.go
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Opt())
}
```
```
$ go run getPackage.go 
getPackage.go:6:2: no required module provides package rsc.io/quote: go.mod file not found in current directory or any parent directory; see 'go help modules'
```
This is because we till now did not installed this module.
To install a published module in go use below command.
```
go get rsc.io/quote
```
To uninstall a module in go use below command.
```
go get rsc.io/quote@none
```
Now if we try to install the required module we see below error.
```
$ go get rsc.io/quote@none
go: go.mod file not found in current directory or any parent directory.
        'go get' is no longer supported outside a module.
```
Because `getPackage.go` is not part of any module. Let's create a module
```
$ go mod init main
go: creating new go.mod: module main
go: to add module requirements and sums:
        go mod tidy
```
```go
//  go.mod
module main

go 1.18
```
Now we are ready to install any published package.
```
$ go get rsc.io/quote
go: added golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: added rsc.io/quote v1.5.2
go: added rsc.io/sampler v1.3.0
```
This also create a `go.sum` file which contains hash of the modules installed and `go.mod` is updated to
```go
// go.mod
module main

go 1.18

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/quote v1.5.2 // indirect
	rsc.io/sampler v1.3.0 // indirect
)
```
Now we can execute the actual file.
```
$ go run getPackage.go 
If a program is too slow, it must have a loop.
```

## Unix stdin, stdout, and stderr
Every Unix operating system has three files open all the time for its processes. As you know, Unix considers everything a file, even a printer or your mouse.

By default, all Unix systems support three special and standard filenames: `/dev/stdin` , `/dev/stdout` , and `/dev/stderr` , which can also be accessed using file descriptors 0, 1, and 2, respectively. These three file descriptors are also called standard input, standard output, and standard error, respectively. Additionally, file descriptor 0 can be accessed as `/dev/fd/0` on a macOS machine and as both `/dev/fd/0` and `/dev/pts/0` on a Debian Linux machine.

Go uses `os.Stdin` for accessing standard input, `os.Stdout` for accessing standard output, and `os.Stderr` for accessing standard error. Although you can still use `/dev/stdin` , `/dev/stdout` , and `/dev/stderr` , or the related file descriptor values for accessing the same devices, it is better, safer, and more portable to stick with the `os.Stdin` , `os.Stdout` , and `os.Stderr` standard filenames that Go offers.

## About printing output
The simplest way to print something in Go is by using the `fmt.Println()` and `fmt.Printf()` functions. You can also use the `fmt.Print()` function instead of `fmt.Println()` . The main difference between `fmt.Print()` and `fmt.Println()` is that the `fmt.Println()` automatically adds a newline character each time you call it, whereas the biggest difference between `fmt.Println()` and `fmt.Printf()` is that the `fmt.Printf()` requires a format specifier for each thing that you want to print, just like the C `printf(3)` function, which means that you have better control over what you are doing, though you have to write
more code. 

Go calls these format specifiers **verbs**. You can find more information about verbs at https://golang.org/pkg/fmt/ . If you have to perform any formatting before printing something, or you have to arrange multiple variables, then using `fmt.Printf()` might be a better choice. However, if you only have to print a single variable, then you might need to choose either `fmt.Print()` or `fmt.Println()` , depending on whether you need the newline character or not.

Apart from `fmt.Println()` , `fmt.Print()` , and `fmt.Printf()` , which are the simplest functions that can be used for generating output on the screen, there is also the `S` family of functions that includes `fmt.Sprintln()` , `fmt.Sprint()` , and `fmt.Sprintf()` , which are used for creating strings based on the given format and the `F` family of functions. This includes `fmt.Fprintln()` , `fmt.Fprint()` and `fmt.Fprintf()` , which are used for writing to files using an `io.Writer`.

```go
// printing.go
package main

import "fmt"

func main() {
	v1 := "123"
	v2 := 123
	v3 := "Have a nice day\n"
	v4 := "abc"

	fmt.Println(v1, v2, v3, v4)                    // add a space among each variable and a line after
	fmt.Print(v1, v2, v3, v4)                      // no space and no line in the end
	fmt.Println()                                  // just to add a line after the previous print end
	fmt.Print(v1, " ", v2, " ", v3, " ", v4, "\n") // in order to make in similar to other print we need to add literals manually
	fmt.Printf("%s%d %s %s\n", v1, v2, v3, v4)     // in order to get even fine control
}
```
```
123 123 Have a nice day
 abc
123123Have a nice day
abc
123 123 Have a nice day
 abc
123123 Have a nice day
 abc
```

## Using standard output
Here `stdOUT.go` uses the `io` package instead of the `fmt` package. The `os` package is used for reading the command-line arguments of the program and for accessing `os.Stdout`.

```go
// stdOUT.go
package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
```
```
$ go run stdOUT.go first_argument
first_argument
```

## Getting user input
There are three main ways to acquire user input:
1. By reading the command-line arguments of a program
2. By asking the user for input
3. By reading external files

## About := and =
The official name for `:=` is the **short assignment statement**. The short assignment statement can be used in place of a `var` declaration with an implicit type (declaring variable without specifing datatype).

The `var` keyword is mostly used for declaring global variables in Go programs as well as for declaring variables without an initial value. The reason for the former is that every statement that exists outside of the code of a function must begin with a keyword, such as `func` or `var`. This means that the short assignment statement cannot be used outside of a function because it is not available there.

`:=` is for both declaration, assignment, and also for redeclaration; and it guesses (infers) the variable's type automatically whereas `=` is for assignment only.

For example, `foo := 32` is a short-hand form of:
```go
var foo int
foo = 32
// OR:
var foo int = 32
// OR:
var foo = 32
```

But there are some rules
1. You can't use `:=` outside of `funcs`. It's because, outside a `func`, a statement should start with a keyword.
```go
// no keywords below, illegal.
illegal := 42

// `var` keyword makes this statement legal.
var legal = 42

func foo() {
  alsoLegal := 42
  // reason: it's in a func scope.
}
```

2. You can't use them twice (in the same scope) because, `:=` introduces "a new variable", hence using it twice does not redeclare a second variable, so it's illegal.
```go
legal := 42
legal := 42 // <-- error
```

3. You can use them for multi-variable declarations and assignments
```go
foo, bar   := 42, 314
jazz, bazz := 22, 7
```

4. You can use them twice in "multi-variable" declarations, if one of the variables is new. This is legal, because, you're not declaring all the variables, you're just reassigning new values to the existing variables, and declaring new variables at the same time. This is called **redeclaration**.
```go
foo, bar  := someFunc()
foo, jazz := someFunc()  // <-- jazz is new
baz, foo  := someFunc()  // <-- baz is new
```

5. You can use the short declaration to declare a variable in a newer scope even if that variable is already declared with the same name before:
```go
var foo int = 34

func some() {
  // because foo here is scoped to some func
  foo := 42  // <-- legal	(declare + assign)
  foo = 314  // <-- legal	(only assign)
}
```

6. You can declare the same name in short statement blocks like: **if**, **for**, **switch**:
```go
foo := 42
if foo := someFunc(); foo == 314 {
  // foo is scoped to 314 here
  // ...
}
// foo is still 42 here
```

So, as a general rule: If you want to easily declare a variable you can use `:=`, or, if you only want to overwrite an existing variable, you can use `=`.

## Reading from standard input
```go
// stdIN.go

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}
```
```
Hello
> Hello
Hi
> Hi
```
Here there is a call to `bufio.NewScanner()` using standard input ( `os.Stdin` ) as its parameter. This call returns a `bufio.Scanner` variable, which is then used with the `Scan()` function for reading from `os.Stdin` line by line. Each line that is read is printed on the screen before getting the next one.

## Working with command-line arguments
```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}
	arguments := os.Args                           // os.Args is a Go slice with string values
	min, _ := strconv.ParseFloat(arguments[1], 64) // convert string to float 64 bit
	max, _ := strconv.ParseFloat(arguments[1], 64) // convert string to float 64 bit

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
```
```
$ go run cla.go 2.5 3.5 6.5
Min: 2.5
Max: 6.5
```
Note that `os.Args` is a Go slice with string values. The first element in the slice is the name of the executable program. Therefore, in order to initialize the `min` and `max` variables, you will need to use the second element of the `os.Args` string slice that has an index value of 1.

`cla.go` ignores the error value returned by the `strconv.ParseFloat()` function using the following statement:
```go
n, _ := strconv.ParseFloat(arguments[i], 64)
```
The underscore character, which is called a **blank identifier**, is the Go way of discarding a value. As you might expect, the program does not behave well when it receives erroneous input. Worst of all, the program does not generate any warnings to inform the user that there were one or more errors while processing the command-line arguments.
```
$ go run cla.go a b c 10
Min: 0
Max: 10
```

## About error output
```go
package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}
	io.WriteString(os.Stdout, "This is Standard output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}
```
The preceding output cannot help you differentiate between data written to standard output and data written to standard error, which could be very useful at times.
Thus, when using `bash(1)` , you can redirect the standard error output to a file as follows:
```
$ go run stdERR.go 2>/tmp/stdError
This is Standard output
$ cat /tmp/stdError
Please give me one argument!
```

## Writing to log files
The `log` package allows you to send log messages to the system logging service of your Unix machine, whereas the `syslog` Go package, which is part of the log package, allows you to define the logging level and the logging facility that your Go program will use.

Usually, most system log files on a Unix operating system can be found under the `/var/log` directory. However, the log files of many popular services, such as Apache and Nginx, can be found elsewhere, depending on their configuration.

The `log` package offers many functions for sending output to the syslog server of a Unix machine. The list of function includes `log.Printf()` , `log.Print()` , `log.Println()` , `log.Fatalf()` , `log.Fatalln()` , `log.Panic()` , `log.Panicln()` , and `log.Panicf()` .

## Logging levels
The logging level is a value that specifies the severity of the log entry. Various logging levels exist including `debug` , `info` , `notice` , `warning` , `err` , `crit` , `alert` , and `emerg` (in reverse order of severity).

## Log servers
All Unix machines have a separate server process that is responsible for receiving logging data and writing it to log files. Various log servers exist that work on Unix machines; however, only two of them are used on most Unix variants: `syslogd(8)` and `rsyslogd(8)`.

## A Go program that sends information to log files

```go
package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0]) // logFiles
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7,
		programName)
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")
	sysLog, err = syslog.New(syslog.LOG_MAIL, "Some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_MAIL: Logging in Go!")
	fmt.Println("Will you see this?")
}
```
```
$ go run logFiles.go 
Will you see this?
```
The logs are stored in some different location.
```
$ grep LOG_ /var/log/syslog
Jun 14 05:47:48 workstation logFiles[5784]: 2022/06/14 05:47:48 LOG_INFO + LOG_LOCAL7: Logging in Go!
Jun 14 05:47:48 workstation Some program![5784]: 2022/06/14 05:47:48 LOG_MAIL: Logging in Go!
```
The important thing to remember from this section is that if the logging server of a Unix machine is not configured to catch all logging facilities, some of the log entries that you send to it might get discarded without any warnings.

## About log.Fatal()
The `log.Fatal()` function is used when something really bad has happened, and you just want to exit your program as fast as possible after reporting the bad situation.
```go
package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL,
		"Some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Fatal(sysLog)
	fmt.Println("Will you see this?")
}
```
```
$ go run logFatal.go 
exit status 1
```
As you can easily understand, the use of `log.Fatal()` terminates a Go program at the point where `log.Fatal()` was called, which is the reason that you did not see the output from the `fmt.Println("Will you see this?")` statement.
However, because of the parameters of the `syslog.New()` call, a log entry has been added to the log file that is related to mail, which is `/var/log/mail.log`.
```
$ grep 'MAIL' /var/log/mail.log 
Jun 14 05:47:48 workstation Some program![5784]: 2022/06/14 05:47:48 LOG_MAIL: Logging in Go!
```

## About log.Panic()
There are situations where a program will fail for good, and you want to have as much information about the failure as possible. In such difficult circumstances, you might consider using `log.Panic()`.

```go
package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL,
		"Some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Panic(sysLog)
	fmt.Println("Will you see this?")
}
```
```
$ go run logPanic.go 
panic: &{17 Some program! workstation   {0 0} 0xc0000a6078}

goroutine 1 [running]:
log.Panic({0xc000048f60?, 0x515198?, 0xc0000ae120?})
        /usr/local/go/src/log/log.go:385 +0x65
main.main()
        /home/himanshu/HP/dev/Learn-Go/go_and_os/logPanic.go:17 +0xb7
exit status 2
```
Looking for logs in logs files
```
$ grep 'Some program' /var/log/mail.log 
Jun 14 06:27:59 workstation Some program![9850]: 2022/06/14 06:27:59 &{17 Some program! workstation   {0 0} 0xc0000a6078}
```

## Error handling in Go
There is a separate data type for errors, named `error`.
In order to create a new error variable, you will need to call the `New()` function of the errors standard Go package.

```go
package main

import (
	"errors"
	"fmt"
)

func returnError(a, b int) error {
	if a == b {
		err := errors.New("Error in returnError() function!")
		return err
	} else {
		return nil
	}
}

func main() {
	err := returnError(1, 2)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}
	err = returnError(10, 10)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}
	if err.Error() == "Error in returnError() function!" {
		fmt.Println("!!")
	}
}
```
```
$ go run newError.go 
returnError() ended normally!
Error in returnError() function!
!!
```

If you try to compare an `error` variable with a string variable without converting the `error` variable to a `string` first, the Go compiler will create the following error message:
```
# command-line-arguments
./newError.go:33:9: invalid operation: err == "Error in returnError()
function!" (mismatched types error and string)
```

## Error handling
```go
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// not extra cmd line arg given
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}
	// get cmd line args
	arguments := os.Args
	var err error = errors.New("An error")

	/*
		This is the trickiest part of the program because, if the first command-line argument is not a
		proper float, you will need to check the next one and keep checking until you find a suitable
		command-line argument. If none of the command-line arguments are in the correct format,
		errors.go will terminate and print a message on the screen. All this checking happens by
		examining the error value that is returned by strconv.ParseFloat() . All of this code is
		there just for the accurate initialization of the min and max variables.
	*/

	k := 1
	var n float64
	for err != nil {
		if k >= len(arguments) {
			fmt.Println("None of the arguments is a float!")
			return
		}
		n, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}
	min, max := n, n
	for i := 2; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
```
```
$ go run errors.go st 5 pp 1
Min: 1
Max: 5
```
