package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/haukened/gourban"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprint(os.Stderr, "Missing required query string.\n")
		os.Exit(1)
	}
	queryString := strings.Join(os.Args[1:], " ")
	test, err := gourban.Query(queryString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("[%s]: %s\n", test[0].Word, test[0].Definition)
}
