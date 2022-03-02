package main

import (
	"fmt"
)

func constDemo() {

	const PI float32 = 3.1415926
	fmt.Printf("PI: %v\n", PI)

	const PI2 = 3.1415926
	fmt.Printf("PI2: %v\n", PI2)

	const (
		a = 100
		b = 200
		c = 300
	)

	const w, h = 200, 300

	const (
		a1 = iota
		b1 = iota
		c1 = iota
	)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("c1: %v\n", c1)

	const a2 = iota
	fmt.Printf("a2: %v\n", a2)

	const (
		d1 = iota
		_  //跳过
		d2 = iota
	)
	fmt.Printf("d1: %v\n", d1)
	fmt.Printf("d2: %v\n", d2)

	const (
		f1 = iota
		f2 = 100
		f3 = iota
	)
	fmt.Printf("f1: %v\n", f1)
	fmt.Printf("f2: %v\n", f2)
	fmt.Printf("f3: %v\n", f3)
}
