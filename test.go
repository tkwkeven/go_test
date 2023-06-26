package main //声明文件所在包，每个go文件必须有归属包
import "fmt"

func main() {

	var list []int = make([]int, 5)
	var mp map[string]int = make(map[string]int, 5)

	var ch chan int64 = make(chan int64, 2)

	fmt.Println(list)
	fmt.Println(mp)
	fmt.Println(ch)

	if list == nil {
		fmt.Println("list为空")
	} else if mp == nil {
		fmt.Println("mp为空")
	} else if ch == nil {
		fmt.Println("ch为空")

	}
	/*var i *int = new(int)
	fmt.Println(*i)

	var list []int = make([]int, 5, 10)
	fmt.Println(len(list), cap(list))

	list[4] = 11
	fmt.Println(list)

	list1 := list[:10]
	fmt.Println(list1)*/

	/*str1 := "abc"
	fmt.Println(utf8.ValidString(str1))*/

	/*fmt.Println("ff")
	f1.F1()
	f2.F2()*/
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

	/*var num3 = 100
	var str1 = "100a"
	fmt.Println(strconv.Itoa(num3))
	fmt.Println(strconv.Atoi(str1))*/

	//var str2 string = "今天天气真好"
	//var byte1 = []byte{228, 187, 138, 229, 164, 169, 229, 164, 169, 230, 176, 148, 231, 156, 159, 229, 165, 189}

	/*fmt.Println([]byte(str2))
	fmt.Println(string(byte1))*/

	/*fmt.Println([]rune(str2))

	var runeList = []rune{20170, 22825, 22825, 27668, 24456, 22909}
	fmt.Println(string(runeList[3]))
	fmt.Println(string(runeList))*/

	/*var inf interface{} = "100"
	i, err := inf.(int)
	fmt.Println(i, err)*/

}

/*func get() (res int, err error) {
	fmt.Println("fff")
	return 1, nil
}*/
