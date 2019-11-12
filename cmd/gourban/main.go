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
	fmt.Println(test[0].Definition)
}
