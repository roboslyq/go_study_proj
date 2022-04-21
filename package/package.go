package main

import (
	"fmt"
	//pkg 为导入包的别名，有了别名后就不能使用原包名packagedemo
	pkg "go_study_proj/package/packagedemo"
)

func main() {
	demo := new(pkg.Packagedemo)
	demo.Name = "roboslyq"
	fmt.Printf(demo.Name)
	demo.ShowName()
	//不能访问
	//demo.playName()
}
