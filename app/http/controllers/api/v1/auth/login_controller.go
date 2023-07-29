package auth

import (
	"server/app/http/controllers/api"
	"server/app/requests"
	"server/pkg/auth"
	"server/pkg/jwt"
	"server/pkg/response"

	"github.com/gin-gonic/gin"
)

// LoginController 用户控制器
type LoginController struct {
	api.BaseController
}

// LoginByPhone 手机登录
func (lc *LoginController) LoginByPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	// 2. 尝试登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "账号不存在")
	} else {
		// 登录成功
		accessToken := jwt.NewJWT().IssueToken(user)
		response.CreatedJSON(c, gin.H{
			"token_type":   "Bearer",
			"expires_in":   jwt.NewJWT().ExpireAtTime().Unix(),
			"access_token": accessToken,
		}, "登录成功！")
	}
}

// LoginByPassword 多种方法登录，支持手机号、email 和用户名
func (lc *LoginController) LoginByPassword(c *gin.Context) {
	// 1. 验证表单
	request := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return
	}

	// 2. 尝试登录
	user, err := auth.Attempt(request.UserName, request.Password)
	if err != nil {
		// 失败，显示错误提示
		response.Unauthorized(c, "账号不存在或密码错误")

	} else {
		// 登录成功
		accessToken := jwt.NewJWT().IssueToken(user)
		response.CreatedJSON(c, gin.H{
			"token_type":   "Bearer",
			"expires_in":   jwt.NewJWT().ExpireAtTime().Unix(),
			"access_token": accessToken,
		}, "登录成功！")
	}
}

// RefreshToken 刷新 Access Token
func (lc *LoginController) RefreshToken(c *gin.Context) {

	accessToken, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		response.Error(c, err, "令牌刷新失败")
	} else {
		response.Data(c, gin.H{
			"token_type":   "Bearer",
			"expires_in":   jwt.NewJWT().ExpireAtTime().Unix(),
			"access_token": accessToken,
		})
	}
}
