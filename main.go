package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	src := readInput()
	count := wordCount(src)
	fmt.Println(count)
}

// match returns true if src matches pattern,
// false otherwise.
func wordCount(src string) int {
	return len(strings.Fields(src))
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string) {
	flag.Parse()
	src = strings.Join(flag.Args(), "")
	return src
}
