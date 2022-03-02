package main

import "fmt"

func indentDemo() {
	var name string
	var age int
	var _sys int
	name = "name"
	age = 11
	_sys = 1

	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("_sys: %v\n", _sys)

	var (
		title string
		total int
		b     bool
	)
	fmt.Printf("title: %v\n", title)
	fmt.Printf("total: %v\n", total)
	fmt.Printf("b: %v\n", b)

	// 类型推断
	var name1 = "tom"
	var age1 = 11

	fmt.Printf("name1: %v\n", name1)
	fmt.Printf("age1: %v\n", age1)
	//批量初始化
	var name2, age2, b2 = "tom", 11, true
	fmt.Printf("name2: %v\n", name2)
	fmt.Printf("age2: %v\n", age2)
	fmt.Printf("b2: %v\n", b2)

	//短变量声明 :=   只能在函数内部
	name3 := "name3"
	age3 := 33
	fmt.Printf("name3: %v\n", name3)
	fmt.Printf("age3: %v\n", age3)

	name4, age4 := getNameAndAge()
	fmt.Printf("name4: %v\n", name4)
	fmt.Printf("age4: %v\n", age4)

	//匿名变量
	_, age5 := getNameAndAge2()
	fmt.Printf("age5: %v\n", age5)

}

func getNameAndAge() (name string, age int) {
	name = "tom444"
	age = 4444
	return name, age
}

//匿名变量
func getNameAndAge2() (string, int) {
	return "tom555", 555
}
