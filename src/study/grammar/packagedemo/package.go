package packagedemo

import "fmt"

/**
1、使用首字母大小写来简单方便的判断权限。大写其它包有权限访问，小写其它包没有权限访问。
*/
type Packagedemo struct {
	Name string
}

/**
ShowName ： 首字母大写，其它包可以访问
*/
func (p *Packagedemo) ShowName() {
	fmt.Printf("show结构体的名称： %s \n", p.Name)
}

/**
playName ： 首字母小写，其它包不可以访问
*/
func (p *Packagedemo) playName() {
	fmt.Printf("play结构体的名称： %s \n", p.Name)
}

func init() {
	fmt.Println("程序启动前，包初始化...")
}
