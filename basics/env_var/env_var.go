package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO: ", os.Getenv("FOO")) // 1
	// return empty string if env var is not present
	fmt.Println("BAR: ", os.Getenv("BAR"))

	// os.Environ to list all key/value pairs in the environment
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0]) // printing only env var name, not values
	}
}

// $ go run env_var/env_var.go
// FOO:  1
// BAR:

// SHELL
// SESSION_MANAGER
// QT_ACCESSIBILITY
// COLORTERM
// XDG_CONFIG_DIRS
// SSH_AGENT_LAUNCHER
// XDG_MENU_PREFIX
// ....

// $ BAR=5 go run env_var/env_var.go
// FOO:  1
// BAR: 5

// SHELL
// SESSION_MANAGER
// QT_ACCESSIBILITY
// COLORTERM
// XDG_CONFIG_DIRS
// SSH_AGENT_LAUNCHER
// XDG_MENU_PREFIX
// ....
