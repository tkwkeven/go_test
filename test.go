package main //声明文件所在包，每个go文件必须有归属包
import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	done := make(chan struct{})
	ch := make(chan interface{}, 10)
	work := 10

	for i := 0; i < work; i++ {
		wg.Add(1)
		go doo(i, ch, done, &wg) // wg 传指针，doo()内部会改变 wg的值
	}
	// 主进程充当生产者，向通道里发送消息
	for i := 0; i < work; i++ {
		ch <- i
	}
	// 关闭通道，当生产者完成所有消息的生产之后，主动关闭通道
	close(ch)
	// 关闭done通道，利用广播消息通知所有 goroutine 关闭信号
	close(done)
	wg.Wait()
	fmt.Println("all done!")

	/*work := 2
	for i := 0; i < work; i++ {
		go doo(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("all done!")*/

	/*var mp = make(map[int]string)
	mp[1] = "a"
	mp[2] = "b"
	mp[3] = "c"

	fmt.Println(mp)

	for k, v := range mp {
		fmt.Println(k, v)
	}*/

	/*tmp := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	mp := map[int]map[string]int{
		1: tmp,
		2: tmp,
		3: tmp,
	}

	fmt.Println(mp)
	m := mp[1]
	m["a"] = 10
	fmt.Println(mp)*/

	/*var arr [5]int
	setaaa(arr)
	fmt.Println(arr)*/

	/*var list []int = make([]int, 5)
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

	}*/
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

func doo(work int, ch <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", work)
	defer wg.Done()
	for {
		select {
		case m, ok := <-ch:
			if ok { // 判断通道是否关闭，false表示关闭，true表示正常
				fmt.Printf("[%v] m => %v\n", work, m)
			}
		case <-done: // 只要接受到该通道发来的任何消息，都判定为收到关闭消息，改操作利用了通道的关闭时的广播消息
			fmt.Printf("[%v] IS done\n", work)
			return
		}
	}
}

/*func setaaa(arr [5]int) {
	arr[2] = 1
	fmt.Println(arr)
}*/

/*func get() (res int, err error) {
	fmt.Println("fff")
	return 1, nil
}*/
