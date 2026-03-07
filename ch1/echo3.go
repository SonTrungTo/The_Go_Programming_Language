//go:build ignore

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	// Don't care about format version
	fmt.Println(os.Args[1:])
}
