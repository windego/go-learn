package main

import "fmt"

func rangeDemo() {
	//数组
	// var a = [5]int{1, 2, 3, 4, 5}
	// for i, v := range a {
	// 	fmt.Printf("i: %d, v: %v\n", i, v)
	// }

	// var s = "多课网，go教程"
	// for i, v := range s {
	// 	fmt.Printf("i: %v, v: %c\n", i, v)
	// }

	//切片
	// var s = []int{1, 2, 3, 4, 5}
	// for i, v := range s {
	// 	fmt.Printf("i, %d, v: %v\n", i, v)
	// }

	//map
	// m := make(map[string]string)
	// m["name"] = "tom"
	// m["age"] = "20"
	// m["email"] = "tom@gmail.com"
	// for k, v := range m {
	// 	fmt.Printf("k: %v, v: %v\n", k, v)
	// }

	// maps := map[string]string{"name": "ddd"}
	// for k, v := range maps {
	// 	fmt.Printf("k: %v, v: %v\n", k, v)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break // 退出循环
	// 	}
	// 	fmt.Printf("i: %v\n", i)
	// }

	// MY_LABEL:
	// 	for i := 0; i < 10; i++ {
	// 		if i == 5 {
	// 			break MY_LABEL
	// 		}
	// 		fmt.Printf("%v\n", i)
	// 	}
	// 	fmt.Println("end...")
	// MY_LABEL:
	// for i := 0; i < 5; i++ {
	// MY_LABEL:
	// 	for j := 0; j < 5; j++ {
	// 		if i == 2 && j == 2 {
	// 			continue MY_LABEL
	// 		}
	// 		fmt.Printf("i=%d,j=%d\n", i, j)
	// 	}
	// }

	//
	// a := 0
	// if a == 1 {
	// 	goto LABEL1
	// } else {
	// 	fmt.Println("other")
	// }
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto LABEL1
		}
		fmt.Printf("%v\n", i)
	}
	fmt.Printf("next11111...")

LABEL1:
	fmt.Printf("next...")

}
