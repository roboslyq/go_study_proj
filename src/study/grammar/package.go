package main

import (
	"fmt"
	"study/grammar/packagedemo"
)

func main() {
	fmt.Printf("欢迎使用go语言 \n")
	demo := new(packagedemo.Packagedemo)
	demo.Name = "roboslyq"
	fmt.Printf(demo.Name)
	demo.ShowName()
}
