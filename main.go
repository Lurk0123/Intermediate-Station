package main

import (
	"Intermediate-Station/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	commandQueue := &utils.Queue{}
	// 创建Gin路由器s
	r := gin.Default()

	// 定义GET请求处理程序
	r.GET("/command", func(c *gin.Context) {
		command := commandQueue.Dequeue()
		if command != "" {
			c.JSON(http.StatusOK, gin.H{
				"command": commandQueue.Dequeue(),
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"result": "failed",
			})
		}
	})
	r.POST("/sendCommand", func(c *gin.Context) {
		commandQueue.Enqueue(c.PostForm("command"))
		c.JSON(http.StatusOK, gin.H{
			"result": "success",
		})
	})
	r.POST("/clearCommand", func(c *gin.Context) {
		commandQueue.Clear()
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	})

	// 启动HTTP服务器，监听端口8080
	r.Run(":8080")
}
