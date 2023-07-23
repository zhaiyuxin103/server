// Package auth 处理用户身份认证相关逻辑
package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "server/app/http/controllers/api/v1"
	"server/app/models/user"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	// 请求对象
	type PhoneExistRequest struct {
		Phone string `json:"phone" form:"phone" binding:"required"`
	}
	request := PhoneExistRequest{}

	// 解析 JSON 请求
	if err := c.ShouldBind(&request); err != nil {
		// 解析失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "fail",
			"code":    http.StatusUnprocessableEntity,
			"message": "操作失败！",
			"data":    gin.H{},
			"error":   err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}

	//  检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "操作成功！",
		"data": gin.H{
			"exist": user.IsPhoneExist(request.Phone),
		},
		"error": gin.H{},
	})
}
