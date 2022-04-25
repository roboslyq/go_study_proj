package impl

import (
	"errors" // 引入库errors包，定义错误相关
)

func Login(userName string) (string, error) {
	if userName != "roboslyq" {
		return "", errors.New("用户不存在")
	}
	return "登录成功", nil
}
