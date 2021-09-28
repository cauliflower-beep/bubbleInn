package main

import (
	"BubbleInn/web/controller"
	"github.com/gin-gonic/gin"
)


//添加gin框架开发3步骤：
func main(){

	//初始化Redis连接池

	//1.初始化路由
	router := gin.Default()


	//设置静态资源路径，会进入view中找到index.html并展示
	router.Static("/home","view")
	router.GET("/api/v1.0/session",controller.GetSession)
	router.GET("api/v1.0/imagecode/:uuid",controller.GetTmageCd)

	//3.启动运行
	router.Run("192.168.1.8:8080")  //当然其他端口也完全可以
}

