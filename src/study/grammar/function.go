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
