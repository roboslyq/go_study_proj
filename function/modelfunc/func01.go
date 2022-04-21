package modelfunc

import (
	"fmt"
)

/*
* 函数的定义
* func : 关键字，表示这是一个函数
* Funce01: 函数名称，首字母大写表示公有函数，外面可以直接引用。如果首字母是小写字母，则只能在当前包使用
* (s string):表示一个string类的参数，形参的名称就s
* (res string):返回返回一个string 类型的字符串
 */
func Func01(s string) (res string) {
	fmt.Println(s)
	return "hello," + s
}

/*
* 返回多个参数
 */
func MutilReturnVal() (string, int) {
	return "1", 1
}

/**
* 私有函数，外面的包不能引用
 */
func mutilReturnVal() (string, int) {
	return "1", 1
}

/**
 * 包的初始化函数，在包执行之前执行相关的初始化函数入品
 */
func init() {
	fmt.Printf("package init ... ...\n")
}
