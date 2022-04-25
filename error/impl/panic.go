package impl

import (
	"fmt"
)

func Pdefer() {
	fmt.Println("Enter defer function.")
	if p := recover(); p != nil {
		fmt.Printf("panic: %s\n", p)
	}
	fmt.Println("Exit defer function.")
}

// 引发 panic。
func PanicDemo() {
	// panic(errors.New("something wrong"))
	aint := []int{1, 2, 3}
	fmt.Println(aint[4])
}
