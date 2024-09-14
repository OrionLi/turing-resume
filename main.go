package main

import (
	"github.com/gin-gonic/gin"
	"turing-resume/config"
	"turing-resume/routes"
)

func main() {
	// 初始化数据库连接
	config.InitDB()

	// 创建Gin引擎
	router := gin.Default()

	// 加载路由
	routes.SetupRoutes(router)

	// 启动服务器
	router.Run(":8080")
}
