//go:build ignore

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for index, arg := range os.Args {
		fmt.Println(strings.Join([]string{"Index:", fmt.Sprint(index), "Value:", arg}, " "))
	}
}
