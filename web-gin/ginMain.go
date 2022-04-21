package main

import "github.com/gin-gonic/gin"

func main() {
	// 实例化一个默认的gin示例
	r := gin.Default()
	// 设置controller控制器，接收前端GET请求
	r.GET("/", func(c *gin.Context) {
		// 响应数据
		c.JSON(200, gin.H{
			"Blog": "https://roboslyq.github.io/",
		})
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "The user id is  %s", id)
	})
	// 启动监听端口
	r.Run(":8080")
}
