package routes

import (
	"github.com/gin-gonic/gin"
	"turing-resume/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/fortec", controllers.CreateOrUpdateCx) //创新组
	r.POST("/forpla", controllers.CreateOrUpdateCy) //创业组
}
