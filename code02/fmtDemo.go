package main

import "fmt"

type WebServer struct {
	name string
}

func fmtDemo() {

	var site = WebServer{name: "golang"}
	fmt.Printf("site: %v\n", site)
	fmt.Printf("site: %#v\n", site)
	fmt.Printf("site: %T\n", site)

	bool := true
	fmt.Printf("bool: %t\n", bool)
	fmt.Printf("bool: %v\n", bool)

	x := 100
	p := &x
	fmt.Printf("p: %v\n", p)
	fmt.Printf("p: %p\n", p)
	fmt.Printf("x: %p\n", &x)

}
