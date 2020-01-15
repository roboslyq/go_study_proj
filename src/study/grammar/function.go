package main

import (
	"fmt"
)

func main() {
	name, age := mutilReturnVal()
	fmt.Printf("返回名称（+）:" + name + "\n")
	fmt.Printf("返回名称（）:%s \n", name)
	fmt.Printf("返回年龄：%d \n", age)
}

func mutilReturnVal() (string, int) {
	return "1", 1
}

/**
 * 包的初始化函数，在包执行之前执行相关的初始化函数入品
 */
func init() {
	fmt.Printf("package init ... ...\n")
}
