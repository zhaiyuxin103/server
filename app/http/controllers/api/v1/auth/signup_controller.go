// Package auth 处理用户身份认证相关逻辑
package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/spf13/cast"
	"server/app/http/controllers/api"
	"server/app/models"
	"server/app/models/user"
	"server/app/requests"
	"server/pkg/jwt"
	"server/pkg/response"
)

// SignupController 注册控制器
type SignupController struct {
	api.BaseController
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

// SignupUsingPhone 使用手机和验证码进行注册
func (sc *SignupController) SignupUsingPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.SignupRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}

	// 2. 验证成功，创建数据
	userModel := user.User{
		LastName:     request.LastName,
		FirstName:    request.FirstName,
		LastKana:     request.LastKana,
		FirstKana:    request.FirstKana,
		Birthday:     carbon.Date{Carbon: carbon.Time2Carbon(request.Birthday).SetTimezone(carbon.PRC)},
		AvatarID:     cast.ToUint64(request.AvatarID),
		Gender:       request.Gender,
		Email:        request.Email,
		Phone:        request.Phone,
		Password:     request.Password,
		Introduction: request.Introduction,
		CommonTimestampsField: models.CommonTimestampsField{
			State: request.State,
			Order: request.Order,
		},
	}
	userModel.Create()

	if userModel.ID > 0 {
		accessToken := jwt.NewJWT().IssueToken(userModel)
		response.CreatedJSON(c, gin.H{
			"user":         userModel,
			"token_type":   "Bearer",
			"expires_in":   jwt.NewJWT().ExpireAtTime().Unix(),
			"access_token": accessToken,
		}, "注册成功！")
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}

// SignupUsingEmail 使用 Email + 验证码进行注册
func (sc *SignupController) SignupUsingEmail(c *gin.Context) {

	// 1. 验证表单
	request := requests.SignupRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingEmail); !ok {
		return
	}

	// 2. 验证成功，创建数据
	userModel := user.User{
		LastName:     request.LastName,
		FirstName:    request.FirstName,
		LastKana:     request.LastKana,
		FirstKana:    request.FirstKana,
		Birthday:     carbon.Date{Carbon: carbon.Time2Carbon(request.Birthday).SetTimezone(carbon.PRC)},
		AvatarID:     cast.ToUint64(request.AvatarID),
		Gender:       request.Gender,
		Email:        request.Email,
		Phone:        request.Phone,
		Password:     request.Password,
		Introduction: request.Introduction,
		CommonTimestampsField: models.CommonTimestampsField{
			State: request.State,
			Order: request.Order,
		},
	}
	userModel.Create()

	if userModel.ID > 0 {
		accessToken := jwt.NewJWT().IssueToken(userModel)
		response.CreatedJSON(c, gin.H{
			"user":         userModel,
			"token_type":   "Bearer",
			"expires_in":   jwt.NewJWT().ExpireAtTime().Unix(),
			"access_token": accessToken,
		}, "注册成功！")
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}
