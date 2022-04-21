package demo

import (
	"fmt"
	"strconv"
)

func Arraydemo() {
	var a [3]int
	a[0] = 0
	a[1] = 1
	a[2] = 2

	for i, v := range a {
		fmt.Println(strconv.Itoa(i) + "-" + strconv.Itoa(v))
	}
}
