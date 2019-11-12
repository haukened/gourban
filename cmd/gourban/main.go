package main

import (
	"fmt"
	"github.com/haukened/gourban"
)

func main() {
	test, err := gourban.Query("yeet")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s: %s\n", test[0].Word, test[0].Definition)
}
