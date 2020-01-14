package main

/**
 * 数据结构相关语句
 */

import (
	"fmt"
)

func main() {
	structDemo()
}

/**
====================struct====================
*/
// 定义结构体
type Person struct {
	Name string
	Age  int
}

//  (p Person) ：此方法所属结构体
// getName()：方法名称
// string : 方法返回参数类型
func (p Person) getName() string {
	return p.Name
}

// 结构体指针接收者方法(更新必须使用指针)
func (p *Person) setName(name string) string {
	p.Name = name
	return name
}

func structDemo() {
	var person Person
	person = Person{
		Name: "roboslyq",
		Age:  10,
	}

	fmt.Printf("结构体创建: " + person.getName() + "\n")
	person.setName("roboslyq updated")
	fmt.Printf("结构体创建: " + person.getName() + "\n")

	//使用指针访问
	var person1 *Person
	person1 = &person
	fmt.Printf("结构体创建: " + person1.getName() + "\n")

}

/**
====================array====================
*/

/**
====================slice====================
*/

/**
====================map====================
*/
