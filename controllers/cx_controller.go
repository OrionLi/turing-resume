package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"turing-resume/config"
	"turing-resume/models"
)

func CreateOrUpdateCx(c *gin.Context) {
	var cx models.Cx
	if err := c.ShouldBindJSON(&cx); err != nil {
		SendResponse(c, http.StatusBadRequest, "参数解析错误", nil)
		return
	}

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(cx); err != nil {
		SendResponse(c, http.StatusBadRequest, "参数校验失败", err.Error())
		return
	}

	// 检查是否存在该学生ID
	var existingCx models.Cx
	result := config.DB.First(&existingCx, "student_id = ?", cx.StudentID)
	if result.Error != nil {
		// 如果不存在，则创建新记录
		config.DB.Create(&cx)
	} else {
		// 如果存在，则更新记录
		config.DB.Model(&existingCx).Updates(cx)
	}

	SendResponse(c, http.StatusOK, "数据保存成功", nil)
}
