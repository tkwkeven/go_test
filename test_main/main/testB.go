package main

import (
	"fmt"
	"moduleName/method"
)

/*
	func AddFtt(name string, age int) int {
		fmt.Printf("%s %d", name, age)
		return age
	}
*/
var t1 method.JieGouTi1

func main() {

	//泛型的使用
	fmt.Println(method.GetMax_[float64](3, 4))
	fmt.Println(method.GetF(true))

	//method.Newss("aaa", 2)
	/*fmt.Println("------------")
	as := &method.Ssa{Name: "fff"}
	fmt.Println(as.Getname())*/

	/*var ss *method.Ssa
	fmt.Println(ss.Getname())*/

	//fmt.Println(method.Newss("aa", 4))
	//fmt.Println(method.Newss("aa", 4).Getname())

	/*a := 10
	b := 20*/
	//fmt.Println(count.Xxgo_2(a, b))
	//met.Chd(a, &b)
	/*fmt.Println(method.G)

	method.Quanju(a, b)

	fmt.Println(method.G)*/

	//第一种格式化结构体方法
	/*t1.Name = "fff"
	t1.Age = 3
	t1.Sex = "jfaldj"
	*/

	//第二种格式化结构体方法
	/*t2 := new(method.JieGouTi1)
	t2.Name = "aaa"
	t2.Age = 2
	t2.Sex = "bbbbb"*/

	//第三种格式化结构体方法
	/*t1 = method.JieGouTi1{"bcbc", 3223, "lllal"}
	t1 = method.JieGouTi1{Name: "assss", Age: 3}*/

	//第4种格式化结构体方法
	//t4 := &method.JieGouTi1{Name: "kkkkk"}

	//fmt.Println(*t4)

	//fmt.Println(b)

	/*var dd int
	dd = AddFtt("jjj", 2)
	fmt.Println(dd)*/
	//aa = AddFtt("jjj", 2)
	/*roo:
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			fmt.Printf("%d", j)
			if i == 3 {
				break roo
			}
		}
	}
	fmt.Printf("---------")

	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			fmt.Printf("%d", j)
			if i == 3 {
				break
			}
		}
	}
	*/
	/*	var age int
		fmt.Println("输入学生的名字")
		fmt.Scanln(&age)
		fmt.Printf("%d", age)

		fmt.Println("再输入一次")
		fmt.Scanf("%d", &age)
		fmt.Printf("%d", age)*/
	//var counta int = 20
	/*if cob := 22; cob > 30 {
		fmt.Printf("小于")
	} else if cob < 30 {
		fmt.Println("zheshi xiaouoj30")
	}*/

	/*var coa = 20
	switch coa {
	case 10:
		fmt.Printf("10")
	case 21:
		fmt.Printf("21")
	case 20:
		fmt.Printf("20")
	default:
		fmt.Printf("啥也没有")

	}*/
}
