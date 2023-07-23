// Package auth 处理用户身份认证相关逻辑
package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "server/app/http/controllers/api/v1"
	"server/app/models/user"
	"server/app/requests"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	// 初始化请求对象
	request := requests.SignupPhoneExistRequest{}

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

	// 表单验证
	errs := requests.ValidateSignupPhoneExist(&request, c)
	// errs 返回长度等于零即通过，大于 0 即有错误发生
	if len(errs) > 0 {
		// 验证失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "fail",
			"code":    http.StatusUnprocessableEntity,
			"message": "操作失败！",
			"data":    gin.H{},
			"error":   errs,
		})
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

// IsEmailExist 检测邮箱是否已注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {

	// 初始化请求对象
	request := requests.SignupEmailExistRequest{}

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

	// 表单验证
	errs := requests.ValidateSignupEmailExist(&request, c)
	// errs 返回长度等于零即通过，大于 0 即有错误发生
	if len(errs) > 0 {
		// 验证失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "fail",
			"code":    http.StatusUnprocessableEntity,
			"message": "操作失败！",
			"data":    gin.H{},
			"error":   errs,
		})
		return
	}

	//  检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "操作成功！",
		"data": gin.H{
			"exist": user.IsEmailExist(request.Email),
		},
		"error": gin.H{},
	})
}
