package main

import "fmt"

func arrayDemo() {
	// var a [3]int    // 定义一个int类型的数组a，长度是3
	// var s [2]string // 定义一个字符串类型的数组s，长度是2
	// var b [2]bool

	// fmt.Printf("a: %T   val:%v\n", a, a)
	// fmt.Printf("s: %T   val:%v\n", s, s)
	// fmt.Printf("b: %v\n", b)

	// var a = [3]int{1, 2, 3}
	// var s = [2]string{"tom", "kite"}
	// var b = [2]bool{true, false}

	// a1 := [2]int{1, 2} // 类型推断

	// fmt.Printf("a: %v\n", a)
	// fmt.Printf("s: %v\n", s)
	// fmt.Printf("b: %v\n", b)
	// fmt.Printf("a1: %v\n", a1)
	// var a = [...]int{0: 1, 2: 2}
	// var s = [...]string{1: "tom", 2: "kite"}
	// var b = [...]bool{2: true, 5: false}

	// a1 := [...]int{1, 2} // 类型推断

	// fmt.Printf("a: %T   %v\n", a, a)
	// fmt.Printf("s: %T   %v\n", s, s)
	// fmt.Printf("b: %T   %v\n", b, b)
	// fmt.Printf("a1: %v\n", a1)
	// s[10] = "hello"

	//
	// a := [...]int{1, 2, 3, 4, 5, 6}

	// for i := 0; i < len(a); i++ {
	// 	fmt.Printf("a[%d]: %v\n", i, a[i])
	// }\\\\\

	// a := [...]int{1, 2, 3, 4, 5, 6}

	// for i, v := range a {
	// 	fmt.Printf("a[%d]: %v\n", i, v)
	// }

	// var names []string
	// var name [3]string
	// var numbers []int
	// fmt.Printf("names: %T     %v\n", names, names)
	// fmt.Printf("name: %T      %v\n", name, name)
	// fmt.Printf("numbers: %v\n", numbers)
	// fmt.Println(names == nil)
	// fmt.Println(numbers == nil)

	// var s1 = make([]string, 3, 10)
	// fmt.Printf("len: %d cap: %d\n", len(s1), cap(s1))
	// fmt.Printf("s1: %T   %v\n", s1, s1)

	// var s1 = []int{1, 2, 3, 4, 5, 6}
	// s2 := s1[0:3] // [)
	// fmt.Printf("s2: %v\n", s2)
	// s3 := s1[3:]
	// fmt.Printf("s3: %v\n", s3)
	// s4 := s1[2:5]
	// fmt.Printf("s4: %v\n", s4)
	// s5 := s1[:]
	// fmt.Printf("s5: %v\n", s5)

	// var s1 []int
	// fmt.Println(s1 == nil)
	// fmt.Printf("len: %d, cap: %d\n", len(s1), cap(s1))

	// s1 := []int{}
	// s1 = append(s1, 1)
	// s1 = append(s1, 2)
	// s1 = append(s1, 3, 4, 5) // 添加多个元素
	// s1 = append(s1, 6)
	// fmt.Printf("s1: %v\n", s1)

	// s3 := []int{3, 4, 5}
	// s4 := []int{1, 2}
	// s4 = append(s4, s3...) // 添加另外一个切片
	// fmt.Printf("s4: %v\n", s4)

	// s1 := []int{1, 2, 3, 4, 5}
	// // 删除索引为2的元素
	// s1 = append(s1[:2], s1[3:]...)
	// fmt.Printf("s1: %v\n", s1)

	s1 := []int{1, 2, 3}
	s2 := s1
	s1[0] = 100
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Println("----------")

	s3 := make([]int, 3)

	copy(s3, s1)

	s1[0] = 1

	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s3: %v\n", s3)

}
