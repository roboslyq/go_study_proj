package main

/**
 * 流程控制相关语法
 */
import (
	"fmt"
	f1 "go_study_proj/grammar/chaptors"
	"strings"
)

/**

 * 主函数入口
 */
func main() {
	f1.CondController()
	f1.ForController()
	f1.SwitchDemo()
	// ASCII码遍历
	f1.IteralStrAscii("中华人民共和国")
	// Unicode遍历
	f1.IteralStr("中华人民共和国")

	//获取一段字符:字节为单位，注意位置，否则容易乱码
	s := "中华人民共和国"
	s1 := s[5:]
	fmt.Println(s1)                                 //结果是：�人民共和国
	fmt.Println(s[strings.Index(s, "华")+len("华"):]) // 结果是：人民共和国
}
