//go:build ignore

/**
 * This program prints a duplicated line preceded by its count and the names of all files in which it appears.
 * Usage: Use with a list of named files.
 */

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	where := make(map[string]string)

	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "dup: no files\n")
		os.Exit(1)
	} else {
		for _, file := range files {
			b, err := os.ReadFile(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup err: %v\n", err)
				continue
			}
			input := bufio.NewScanner(bytes.NewReader(b))
			for input.Scan() {
				line := input.Text()
				counts[line]++
				if where[line] == "" {
					where[line] = file
					// This is problematic if execute: input1.bak.txt input1.txt and both has "hello"
					// } else if !strings.Contains(where[input.Text()], file) {
				} else {
					exist := false
					files := strings.Split(where[line], ", ")
					for _, f := range files {
						if f == file {
							exist = true
							break
						}
					}
					if !exist {
						where[line] += ", " + file
					}
				}
			}
			// Handling input error
			inputErr := input.Err()
			if inputErr != nil {
				fmt.Fprintf(os.Stderr, "dup err: scanning %s: %v\n", file, inputErr)
				continue
			}
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\t%s\n", count, line, where[line])
		}
	}
}
