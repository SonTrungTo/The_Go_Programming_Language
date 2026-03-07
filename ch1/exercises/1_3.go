//go:build ignore

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for _, v := range os.Args {
		s += sep + v
		sep = " "
	}
	fmt.Printf("Elapsed time after inefficiency: %.6f seconds\n", time.Since(start).Seconds())

	start = time.Now()
	strings.Join(os.Args[0:], " ")
	fmt.Printf("Elapsed time after optimisation: %.6f seconds\n", time.Since(start).Seconds())

	// Some results when args big enough:
	// Elapsed time after inefficiency: 1.772906 seconds
	// Elapsed time after optimisation: 0.000227 seconds
}
