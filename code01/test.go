package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"windego.cn/code01/morestrings"
)

func main() {
	fmt.Println("hello world")
	s := morestrings.ReverseRunes("test")
	fmt.Printf("s: %v\n", s)

	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
