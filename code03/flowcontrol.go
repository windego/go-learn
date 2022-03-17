package main

import "fmt"

func flowcontrol() {
	// var flag = true
	// if flag {
	// 	fmt.Println("flag is true")
	// }
	// fmt.Printf("程序运行结束")
	//test2
	// var age = 20
	// if age > 18 {
	// 	fmt.Println("你是成年人")
	// }
	// fmt.Printf("程序运行结束")

	//test3
	// if age := 20; age > 18 {
	// 	fmt.Println("你是成年人")
	// }
	// fmt.Printf("程序运行结束")

	//test4: 不能使用0或非0表示真假 non-boolean condition in if statement
	// var i = 1
	// if i { // 编译失败
	// 	fmt.Println("here")
	// }
	// fmt.Printf("程序运行结束")
	// if age := 30; age > 18 {
	// 	fmt.Println("你是成年人")
	// } else {
	// 	fmt.Println("你是未成年人")
	// }

	//test4
	// if score := 80; score >= 60 && score <= 70 {
	// 	fmt.Println("C")
	// } else if score > 70 && score <= 90 {
	// 	fmt.Println("B")
	// } else {
	// 	fmt.Println("A")
	// }
	//test5

	// var c string
	// fmt.Println("输入一个字符：")
	// fmt.Scan(&c)

	// if c == "S" {
	// 	fmt.Println("输入第二个字符：")
	// 	fmt.Scan(&c)
	// 	if c == "a" {
	// 		fmt.Println("Saturday")
	// 	} else if c == "u" {
	// 		fmt.Println("Sunday")
	// 	} else {
	// 		fmt.Println("输入错误")
	// 	}
	// } else if c == "F" {
	// 	fmt.Println("Friday")
	// } else if c == "M" {
	// 	fmt.Println("Monday")
	// } else if c == "T" {
	// 	fmt.Println("输入第二个字符：")
	// 	fmt.Scan(&c)
	// 	if c == "u" {
	// 		fmt.Println("Tuesday")
	// 	} else if c == "h" {
	// 		fmt.Println("Thursday")
	// 	} else {
	// 		fmt.Println("输入错误")
	// 	}
	// } else if c == "W" {
	// 	fmt.Println("Wednesday")
	// } else {
	// 	fmt.Println("输入错误")
	// }

	// test6
	// grade := "A"
	// switch grade {
	// case "A":
	// 	fmt.Println("优秀")
	// case "B":
	// 	fmt.Println("良好")
	// default:
	// 	fmt.Println("一般")
	// }

	// day := 3
	// switch day {
	// case 1, 2, 3, 4, 5:
	// 	fmt.Println("工作日")
	// case 6, 7:
	// 	fmt.Println("休息日")
	// }

	//fallthrough
	// a := 300
	// switch a {
	// case 100:
	// 	fmt.Println("100")
	// 	fallthrough
	// case 200:
	// 	fmt.Println("200")
	// 	fallthrough
	// case 300:
	// 	fmt.Println("300")
	// 	fallthrough
	// default:
	// 	fmt.Println("other")
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("i: %v\n", i)
	// }

	// for i := 10; i < 40; i++ {
	// 	fmt.Print(i, " ")
	// }
	// i := 1
	// for ; i <= 10; i++ {
	// 	fmt.Printf("i: %v\n", i)
	// }
	// i := 1 // 初始条件
	// for i <= 10 {
	// 	fmt.Printf("i: %v\n", i)
	// 	i++ // 结束条件
	// }
	var a = [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Printf("i: %d, v: %v\n", i, v)
	}

}
