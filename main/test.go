package main //声明文件所在包，每个go文件必须有归属包

import (
	"fmt"
	_ "fmt"
	"strconv"
)

func main() {
	/*var x int
	fmt.Println(x)*/

	/*var str string = "这是数组"
	var arr [10]int
	var sl []string
	var mp map[string]int

	fmt.Println(str)
	fmt.Println(arr)
	if sl == nil {
		fmt.Println(sl)
	}

	if mp == nil {
		fmt.Println(mp)
	}*/

	/*list := []int{1, 2}
	list1 := []int{
		1,
		2,
	}
	fmt.Println(list, list1)*/
	//fmt.Println()
	/*_, err := 1, 2
	fmt.Println(err)*/
	/*{
		a := 1
		fmt.Println(a)
		{
			a := 2
			fmt.Println(a)
		}
	}*/

	//x := nil
	//var x interface{} = nil
	//var x = nil
	//fmt.Println(x)

	/*var num1 int = 100
	fmt.Println(int64(num1))

	var num2 int64 = 100
	fmt.Println(int(num2))*/

	var num3 = 100
	var str1 = "100a"
	fmt.Println(strconv.Itoa(num3))
	fmt.Println(strconv.Atoi(str1))
}

/*func get() (res int, err error) {
	fmt.Println("fff")
	return 1, nil
}*/
