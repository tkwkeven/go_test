package method

import (
	"errors"
	"fmt"
)

type Users struct {
	Nane string
	Uidd int
}

func Xxgo(a, b int) (uu *Users, c int, err error) {
	if a <= 0 && b <= 0 {
		err = errors.New("数字a 和 数字b 不能同时小于0")
		return
	}
	uu = new(Users)
	uu.Nane = "mzji"
	uu.Uidd = 33
	fmt.Printf("11111%p \n", &uu)
	c = a + b

	return
}
