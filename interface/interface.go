package main

import (
	"fmt"
	"reflect"
)

/**
Interface学习：
1、Interface本质是一种规约（protocol）
2、实现interface:结构体
*/
func main() {
	var t int
	typeOfA := reflect.TypeOf(t)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
	//创建结构体实例
	c := new(cat)
	fmt.Println("test" + reflect.TypeOf(c).Name())
	c1 := &cat{}
	c3 := &cat{
		name: "z",
		age:  0,
	}
	//结构体初始化赋值
	c.name = "roboslyq"
	c.age = 18
	c.name = "roboslyq2"
	c.age = 20

	c1.name = "roboslyq1"
	c1.age = 19

	//定义接口类型
	var a animal
	//将结构体赋值给接口类型
	a = c
	//接口类型调用方法
	a.eat("西瓜")
	c.show()
	c1.show()
	c3.show()
}

type animal interface {
	eat(food string) string
}

type cat struct {
	name string
	age  int32
}

func (c *cat) eat(food string) string {
	fmt.Printf("eat food: %s \n", food)
	return "finish"
}

func (c cat) show() string {
	fmt.Printf("cat name %s \n", c.name)
	fmt.Printf("cat age %d \n", c.age)
	return "finish"
}
