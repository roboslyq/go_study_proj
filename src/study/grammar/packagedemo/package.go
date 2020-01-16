package packagedemo

import "fmt"

type Packagedemo struct {
	Name string
}

func (p *Packagedemo) ShowName() {
	fmt.Printf("结构体的名称： %s \n", p.Name)
}
