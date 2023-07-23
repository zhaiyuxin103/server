// Package auth 处理用户身份认证相关逻辑
package auth

import (
	"github.com/gin-gonic/gin"
	v1 "server/app/http/controllers/api/v1"
	"server/app/models/user"
	"server/app/requests"
	"server/pkg/response"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 获取请求参数，并做表单验证
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	//  检查数据库并返回响应
	response.Data(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// IsEmailExist 检测邮箱是否已注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}

	//  检查数据库并返回响应
	response.Data(c, gin.H{
		"exists": user.IsEmailExist(request.Email),
	})
}
