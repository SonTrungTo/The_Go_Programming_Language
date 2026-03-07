//go:build ignore

// Echo1 prints its command-line arguments.

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	sep = "" // This is done implicitly by the zero value of a string, but it's explicitly set here for clarity.
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Println(s)
	// Full, original version
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	// Auto-generated loop to print command-line arguments, similar to the one above but using range for cleaner syntax.
	for i, arg := range os.Args {
		fmt.Printf("args[%d] = %q\n", i, arg)
	}
}
