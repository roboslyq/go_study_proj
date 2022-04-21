package chaptors

import (
	"fmt"
	"math/rand"
)

/**
 *if判断语句
 */
func CondController() {
	var a = rand.Int31n(1)
	if a >= 5 {
		fmt.Printf("a大于等于5，a == %d \n", a)
	} else {
		fmt.Printf("a小于5，a == %d \n", a)
	}
}

/**
 *for判断语句
 */
func ForController() {
	//传统的for init; condition; post {
	//}

	fmt.Printf("开始传统循环... \n")
	for i := 0; i < 10; i++ {
		CondController()
	}

	//关系表达式或逻辑表达式控制循环
	i := 0
	fmt.Printf("开始关系表达式或逻辑表达式控制循环... \n")
	for i < 5 {
		CondController()
		i++
	}

	//无限循环控制
	fmt.Printf("开始无限循环控制(演示输出5次)... \n")
	count := 5
	for {
		CondController()
		count--
		if count == 0 {
			break
		}
	}
	//开始循环嵌套
	fmt.Printf("开始循环嵌套... \n")
	var k, j int
	for k = 2; k < 100; k++ {
		for j = 2; j <= (k / j); j++ {
			if k%j == 0 {
				break
			}
			if j == k/j {
				fmt.Printf("%d 是素数 \n", j)
			}
		}
	}
}

/*
switch语法
*/

func SwitchDemo() {
	val := 1
	switch val {
	case 1:
		fmt.Println("switch val:", val)
	case 2:
		fmt.Println("switch val:", val)
	case 3:
		fmt.Println("switch val:", val)
	default:
		fmt.Println("switch default:", val)
	}
}
