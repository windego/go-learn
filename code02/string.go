package main

import (
	"bytes"
	"fmt"
	"strings"
)

func strDemo() {

	s := "hello, world"
	a := 2
	b := 5
	sp := strings.Split(s, ",")

	fmt.Printf("%v\n", sp)
	bool := strings.Contains(s, "hello")
	fmt.Printf("bool: %v\n", bool)
	bool1 := strings.HasPrefix(s, "hell")
	fmt.Printf("bool1: %v\n", bool1)

	fmt.Printf("s[a]: %c\n", s[a])
	fmt.Printf("s[a:b]: %v\n", s[a:b])
	fmt.Printf("s[:b]: %v\n", s[:b])
	fmt.Printf("s[a:]: %v\n", s[a:])

	var html string = `<html>
	<head><title>hello golang</title>
	</html>`
	fmt.Printf("html: %v\n", html)

	name := "golang"
	ide := "vscode"
	msg := fmt.Sprintf("hello %s, %s", name, ide)
	fmt.Printf("msg: %v\n", msg)

	// s := strings.Join([]string{name, ide}, ",")
	// fmt.Printf("s: %v\n", s)

	var buffer bytes.Buffer
	buffer.WriteString("zhang")
	buffer.WriteString(",")
	buffer.WriteString("san")
	fmt.Printf("buffer: %v\n", buffer.String())
	print(name + "\n")
	print(name)

}
