package chaptors

import "fmt"

// 遍列字符串中每一个字条
func IteralStrAscii(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

// 遍列字符串中每一个字条
func IteralStr(s string) {
	for index, value := range s {
		fmt.Printf("序号:%d,值：%c \n", index, value)
	}
}
