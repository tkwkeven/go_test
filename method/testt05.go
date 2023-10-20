package method

import (
	"errors"
)

func Xxgo_2(a, b int) (c int, err error) {
	if a <= 0 && b <= 0 {
		err = errors.New("数字a 和 数字b 不能同时小于0")
		return
	}

	c = a + b
	//a += 1
	//*b += 1

	return
}
