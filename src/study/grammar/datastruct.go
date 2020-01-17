package main

/**
 * 数据结构相关语句
 */

import (
	"container/list"
	"fmt"
	"reflect"
	"sync"
)

func main() {
	//structDemo()
	//array()
	//slice()
	//mapDemo()
	listDemo()
}

/**
====================struct====================
*/
// 定义结构体
type Person struct {
	Name string
	Age  int
}

//  (p Person) ：此方法所属结构体（go中的名词叫“接收者”）
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
func array() {
	//声明array
	var intarr [3]int
	//array赋初值
	intarr[0] = 0
	intarr[1] = 1
	intarr[2] = 2
	//遍列array
	for index, val := range intarr {
		fmt.Println("索引值:", index)
		fmt.Println("结果值:", val)
	}

}

/**
数据切片：是对数组的抽象，可以进行动态扩容
====================slice====================
*/
func slice() {
	//声明array并初始化:200当前切片的大小，300容量，预分配
	var intarr = [3]int{100, 200, 300}
	//0:2 表示从第0个元素开始，切2个元素
	var subarr = intarr[0:2]
	//遍列array
	for index, val := range subarr {
		fmt.Println("索引值:", index)
		fmt.Println("结果值:", val)
	}

	//使用make初始化,通过append自动扩容
	var tmp = make([]int, 6, 10)
	tmp[0] = 20
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)
	tmp = append(tmp, 32)

	for index, val := range tmp {
		fmt.Println("索引值:", index)
		fmt.Println("结果值:", val)
	}
}

/**
====================map====================
*/

func mapDemo() {
	//初始化Map
	var valMap = map[string]string{"1": "1", "2": "2 value"}
	//make 初始化map
	var valMap1 = make(map[string]string)
	//赋值
	var key1 = "key1"
	valMap1[key1] = "val1"

	fmt.Println(valMap["2"])
	fmt.Println(valMap1[key1])
	for key, value := range valMap {
		fmt.Println("key: ", key, "value:", value)
	}

	delete(valMap, "1")
	//输出为空（在控制台没法打印）
	fmt.Println(valMap["1"])

	//sync.map
	var syncmap = sync.Map{}
	syncmap.Store("sync_key1", "sync_v1")
	syncmap.Store("sync_key2", "sync_v2")
	syncmap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}

/**
====================list====================
*/

func listDemo() {
	var list1 = list.New()
	var list2 = list.List{}
	list3 := list.New()
	list1.PushBack("1")
	list1.PushBack("2")
	list1.PushBack("2")

	list2.PushFront("2_1")
	list3.PushFront("3_1")

	for val := list1.Front(); val != nil; val = val.Next() {
		fmt.Println(reflect.TypeOf(val))
		fmt.Println(val.Value)
	}
}
