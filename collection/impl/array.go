package impl

/*
	数组与切片区别：
		1、切片是指针类型，数组是值类型；
		2、数组的长度是固定的，而切片不是（切片可以看成动态的数组）；
		3、切片比数组多一个容量（cap）属性；
		4、切片的底层是数组。
		5、a[:]、b[:]是切片，切片之间不能进行等值判断，只能和nil判断。即a[:]==b[:]编译时会抛出异常
*/
import (
	"fmt"
	"strconv"
)

func Arraydemo() {
	var a [3]int
	a[0] = 0
	a[1] = 1
	a[2] = 2

	for i, v := range a {
		fmt.Println(strconv.Itoa(i) + "-" + strconv.Itoa(v))
	}

}

/*
slice的底层结构：
type slice struct {
    array unsafe.Pointer // 指向数组的指针
    len   int // 切片中元素的数量
    cap   int // array 数组的总容量
}
*/
func SliceDemo() {
	var a [4]int
	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3

	slice := a[1:]
	// 以下打印为 [1 2 3]
	fmt.Println(slice)
	// 下面注释放开，编译会报错
	// slice1 := a[2:]
	// fmt.Print(slice == slice1)
	for i, v := range slice {
		fmt.Printf("第%d个元素的值为%d \n", i, v)
	}

	// 深度拷贝: copy(sliceA, sliceB)
	// 浅拷贝: sliceA = sliceB
	var slice1 = slice

	slice1[0] = 10
	for i, v := range slice {
		fmt.Printf("第%d个元素的值为%d \n", i, v)
	}

	for i, v := range slice1 {
		fmt.Printf("第%d个元素的值为%d \n", i, v)
	}
}
