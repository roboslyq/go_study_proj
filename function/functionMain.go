package main

import (
	"fmt"
	f1 "go_study_proj/function/modelfunc"
)

/*
*
 */
func main() {
	// 函数作为变量
	var f = demo
	f()

	// 普通函数调用
	name, age := f1.MutilReturnVal()
	fmt.Printf("返回名称（+）:" + name + "\n")
	fmt.Printf("返回名称（）:%s \n", name)
	fmt.Printf("返回年龄：%d \n", age)
	f1.Func01("test")

	name1, age1 := f1.ImportSamePack()
	fmt.Printf("同包引用返回名称:%s \n", name1)
	fmt.Printf("同包引用返回年龄：%d \n", age1)

	// 直接执行匿名函数，与Js中的Jquery类似
	func() {
		fmt.Println("say 匿名函数")
	}()

	// 匿名函数当参数,作为回调函数
	demo1(func() {
		fmt.Println("匿名函数当参数，回调！！！")
	})
}

func demo() {
	fmt.Println("say hello")
}

func demo1(f func()) {
	fmt.Println("do some 其它事情...")
	f()
}
