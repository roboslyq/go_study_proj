package modelfunc

import (
	"fmt"
)

func ImportSamePack() (string, int) {
	return mutilReturnVal()
}

/**
 * 包的初始化函数，在包执行之前执行相关的初始化函数入品
 */
func init() {
	fmt.Printf("function 02 init ... ...\n")
}
