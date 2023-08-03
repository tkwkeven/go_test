package main

import "fmt"

func main() {
	/*	var age int
		fmt.Println("输入学生的名字")
		fmt.Scanln(&age)
		fmt.Printf("%d", age)

		fmt.Println("再输入一次")
		fmt.Scanf("%d", &age)
		fmt.Printf("%d", age)*/
	//var counta int = 20
	if cob := 22; cob > 30 {
		fmt.Printf("小于")
	} else if cob < 30 {
		fmt.Println("zheshi xiaouoj30")
	}

	var coa = 20
	switch coa {
	case 10:
		fmt.Printf("10")
	case 21:
		fmt.Printf("21")
	case 20:
		fmt.Printf("20")
	default:
		fmt.Printf("啥也没有")

	}
}
