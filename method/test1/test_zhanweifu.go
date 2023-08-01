package main

import (
	"fmt"
)

type simple struct {
	value int
}

func main() {
	a := simple{
		value: 10,
	}

	//ss := "-"

	b := 10
	b2 := 20170

	c := 123.789
	c2 := 123121234.45243524

	//通用占位符
	fmt.Printf("默认格式：%v \n", a)
	fmt.Printf("打印包含字段名的值：%+v \n", a)
	fmt.Printf("打印go语法表示的值：%#v \n", a)
	fmt.Printf("打印go语法表示的类型：%T \n", a)
	fmt.Printf("输出字面上的百分号：%%10 \n")
	//fmt.Printf("%v", ss*100)

	//整数占位符
	fmt.Printf("这是二进制：%b \n", b)
	fmt.Printf("Unicode码点转字符：%c \n", b2)
	fmt.Printf("输出10进制：%d \n", b)
	fmt.Printf("输出8进制：%o \n", b2)
	fmt.Printf("输出0o为前缀的8进制：%O \n", b)
	fmt.Printf("用单引号将字符值包起来：%q \n", b2)
	fmt.Printf("输出16进制：%x \n", b)
	fmt.Printf("输出16进制大写：%X \n", b)
	fmt.Printf("Unicode格式：%U \n", b2)

	fmt.Printf("%v 的二进制为：%b;go语法表示二进制为：%#b ; 指定二进制宽度为8位，不足补0：%08b \n", b, b, b, b)
	fmt.Printf("%v 的十六进制为：%x ；使用go语法表示为：%#x ; 定宽度为8位，不足补0：%#08x \n", b2, b2, b2, b2)
	fmt.Printf("%v 的字符为：%c ；指定宽度为5位，不足补空格：%5c \n", b2, b2, b2)

	//浮点类型占位符
	fmt.Printf("指数为2的幂的无小数科学计数法：%b \n", c)
	fmt.Printf("科学计数法：%e \n", c)
	fmt.Printf("科学计数法，大写：%E \n", c)
	fmt.Printf("有小数点而无指数，既常规的浮点数格式，默认宽度和精度：%f \n", c)
	fmt.Printf("宽度9,精度默认：%9f \n", c)
	fmt.Printf("默认宽度，精度保留两位小数：%.2f \n", c)
	fmt.Printf("宽度为9，精度保留2位小数：%9.2f \n", c)
	fmt.Printf("宽度为9，精度保留0为小数：%9.f \n", c)
	fmt.Printf("根据情况自动选择 %%e 或 %%f 来输出，以产生更紧凑的输出（末位无0）：%g %g \n", c, c2)
	fmt.Printf("根据情况自动选择 %%E 或 %%f 来输出，以产生更紧凑的输出（末位无0）：%G %G \n", c, c2)
	fmt.Printf("16进制方式表示：%x \n", c)
	fmt.Printf("16进制方式表示，大写：%X \n", c)

	//字符串类型占位符
	var str = "今天是个好日子"
	fmt.Printf("描述一下今天： %s \n", str)
	fmt.Printf("描述一下今天：%q \n", str)
	fmt.Printf("%x \n", str)
	fmt.Printf("% X \n", str)
}
