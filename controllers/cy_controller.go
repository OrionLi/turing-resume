package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"turing-resume/config"
	"turing-resume/models"
)

func CreateOrUpdateCy(c *gin.Context) {
	var cy models.Cy
	if err := c.ShouldBindJSON(&cy); err != nil {
		SendResponse(c, http.StatusBadRequest, "参数解析错误", nil)
		return
	}

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(cy); err != nil {
		SendResponse(c, http.StatusBadRequest, "参数校验失败", err.Error())
		return
	}

	// 检查是否存在该学生ID
	var existingCy models.Cy
	result := config.DB.First(&existingCy, "student_id = ?", cy.StudentID)
	if result.Error != nil {
		// 如果不存在，则创建新记录
		config.DB.Create(&cy)
	} else {
		// 如果存在，则更新记录
		config.DB.Model(&existingCy).Updates(cy)
	}

	SendResponse(c, http.StatusOK, "数据保存成功", nil)
}
