package main

import (
	"fmt"
	err "go_study_proj/error/impl"
	"time"
)

func main() {
	// 错误测试
	// errCall()
	// 异常测试：如果没有defer recover则主程序直接结束。如果有defer recover,程序可以一直执行
	for {
		panicCall()
		time.Sleep(time.Duration(2) * time.Second)
	}
}

func errCall() {
	fmt.Printf("请输入用户名：")
	var name string
	_, _ = fmt.Scanln(&name)
	res, error1 := err.Login(name)
	if error1 != nil {
		fmt.Println(error1.Error())
		fmt.Println(error1)
	} else {
		fmt.Println(res)
	}
}

func panicCall() {
	defer err.Pdefer()
	err.PanicDemo()
	fmt.Println("看看是否会打印！！！")
}
