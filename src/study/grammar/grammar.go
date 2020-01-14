package main

import (
	"fmt"
	"math/rand"
)

/**
 * 主函数入口
 */
func main() {
	condController()
	forController()
}

/**
 *if判断语句
 */
func condController()  {
	var a = rand.Int31n(10)
	if a >= 5 {
		fmt.Printf("a大于等于5，a == %d \n" , a)
	}else{
		fmt.Printf("a小于5，a == %d \n",a)
	}
}
/**
 *for判断语句
 */
func forController()  {
	//传统的for init; condition; post {
	//}

	fmt.Printf("开始传统循环... \n")
	for  i  := 0;i<10;i++ {
		condController()
	}

	//关系表达式或逻辑表达式控制循环
	i := 0
	fmt.Printf("开始关系表达式或逻辑表达式控制循环... \n")
	for i < 5 {
		condController()
		i++
	}

	//无限循环控制
	fmt.Printf("开始无限循环控制(演示输出5次)... \n")
	count := 5
	for  {
		condController()
		count--
		if count == 0 {
			break
		}
	}
	//开始循环嵌套
	fmt.Printf("开始循环嵌套... \n")
	var k, j int
	for k = 2; k < 100; k++ {
		for j = 2; j<=(k / j);j++{
			if k%j == 0 {
				break
			}
			if j == k/j {
				fmt.Printf("%d 是素数 \n",j)
			}
		}
	}
}

