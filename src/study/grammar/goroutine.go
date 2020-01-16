package main

import "fmt"
import "time"

func main() {
	//用户可能通过 Enter键停止程序
	go running()
	//用户不能停止程序
	//running()
	fmt.Println("")
	var input string
	fmt.Scanln(&input)
}
func running() {

	var times int
	for {
		times++
		fmt.Println("ticks", times)
		time.Sleep(time.Second)
	}
}
