// go build and go get are two different
// subcommands of the go tool

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// We declare a subcommand using the NewFlagSet
	// function, and proceed to define new flags
	// specific for this subcommand.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// For a different subcommand we can define
	// different supported flags
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// The subcommand is expected as the first argument
	// to the program
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// flag.Parse() // this dosent work while using sub cmd
	// look at all the flags passed
	fmt.Println(os.Args[1:])

	switch os.Args[1] {
	case "foo":
		// only foo cmd is used along with it's sub cmd
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		// additional cmd line param apart from subcmd
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		// additional cmd line param apart from subcmd
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

// $ go run sub_cmd/sub_cmd.go foo -enable -name=joe
// [foo -enable -name=joe]
// subcommand 'foo'
//   enable: true
//   name: joe
//   tail: []

// $ go run sub_cmd/sub_cmd.go bar -level 3
// [bar -level 3]
// subcommand 'bar'
//   level: 3
//   tail: []
