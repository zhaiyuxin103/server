package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"server/app/requests/validators"
	"server/pkg/auth"
	"time"
)

type UserRequest struct {
	LastName  string    `json:"last_name,omitempty" valid:"last_name" form:"last_name"`
	FirstName string    `json:"first_name,omitempty" valid:"first_name" form:"first_name"`
	LastKana  string    `json:"last_kana,omitempty" valid:"last_kana" form:"last_kana"`
	FirstKana string    `json:"first_kana,omitempty" valid:"first_kana" form:"first_kana"`
	Birthday  time.Time `json:"birthday,omitempty" valid:"birthday" form:"birthday" time_format:"2006-01-02"`
	// TODO: exists:avatars,id
	AvatarID     uint64 `json:"avatar_id,omitempty" valid:"avatar_id" form:"avatar_id"`
	Gender       uint8  `json:"gender,omitempty" valid:"gender" form:"gender"`
	Phone        string `json:"phone,omitempty" valid:"phone" form:"phone"`
	Province     string `json:"province,omitempty" valid:"province" form:"province"`
	City         string `json:"city,omitempty" valid:"city" form:"city"`
	District     string `json:"district,omitempty" valid:"district" form:"district"`
	Address      string `json:"address,omitempty" valid:"address" form:"address"`
	Introduction string `valid:"introduction" json:"introduction,omitempty" form:"introduction"`
	State        uint8  `json:"state,omitempty" valid:"state" form:"state"`
	Order        uint64 `json:"order,omitempty" valid:"order" form:"order"`
}

type UserUpdateEmailRequest struct {
	Email      string `json:"email,omitempty" valid:"email" form:"email"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code" form:"verify_code"`
}

type UserUpdatePhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone" form:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code" form:"verify_code"`
}

func UpdateUser(data interface{}, c *gin.Context) map[string][]string {

	// 查询用户名重复时，过滤掉当前用户 ID
	id := auth.CurrentUID(c)

	rules := govalidator.MapData{
		"last_name":    []string{"between:1,20"},
		"first_name":   []string{"between:1,20"},
		"last_kana":    []string{"between:1,20"},
		"first_kana":   []string{"between:1,20"},
		"birthday":     []string{"date"},
		"avatar_id":    []string{"numeric"},
		"gender":       []string{"in:1,2,3"},
		"phone":        []string{"digits:11", "not_exists:users,phone," + id},
		"province":     []string{"min_cn:2", "max_cn:20"},
		"city":         []string{"min_cn:2", "max_cn:20"},
		"district":     []string{"min_cn:2", "max_cn:20"},
		"address":      []string{"min_cn:2", "max_cn:50"},
		"introduction": []string{"min_cn:4", "max_cn:240"},
	}
	messages := govalidator.MapData{
		"last_name": []string{
			"between:姓必须在 1~20 个字符之间",
		},
		"first_name": []string{
			"between:名必须在 1~20 个字符之间",
		},
		"last_kana": []string{
			"between:姓（假名）必须在 1~20 个字符之间",
		},
		"first_kana": []string{
			"between:名（假名）必须在 1~20 个字符之间",
		},
		"birthday": []string{
			"date:生日格式不正确",
		},
		"avatar_id": []string{
			"numeric:头像 ID 必须是数字",
		},
		"gender": []string{
			"in:性别格式不正确",
		},
		"phone": []string{
			"digits:手机号码格式不正确",
			"not_exists:手机号码已经存在",
		},
		"province": []string{
			"min_cn:省份长度需至少 2 个字",
			"max_cn:省份长度不能超过 20 个字",
		},
		"city": []string{
			"min_cn:城市长度需至少 2 个字",
			"max_cn:城市长度不能超过 20 个字",
		},
		"district": []string{
			"min_cn:区县长度需至少 2 个字",
			"max_cn:区县长度不能超过 20 个字",
		},
		"address": []string{
			"min_cn:详细地址长度需至少 2 个字",
			"max_cn:详细地址长度不能超过 50 个字",
		},
		"introduction": []string{
			"min_cn:自我介绍长度需至少 4 个字",
			"max_cn:自我介绍长度不能超过 240 个字",
		},
	}
	return validate(data, rules, messages)
}

func UpdateUserEmail(data interface{}, c *gin.Context) map[string][]string {

	currentUser := auth.CurrentUser(c)
	rules := govalidator.MapData{
		"email": []string{
			"required", "min:4",
			"max:30",
			"email",
			"not_exists:users,email," + currentUser.GetStringID(),
			"not_in:" + currentUser.Email,
		},
		"verify_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被占用",
			"not_in:新的 Email 与老 Email 一致",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*UserUpdateEmailRequest)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)

	return errs
}

func UserUpdatePhone(data interface{}, c *gin.Context) map[string][]string {

	currentUser := auth.CurrentUser(c)

	rules := govalidator.MapData{
		"phone": []string{
			"required",
			"digits:11",
			"not_exists:users,phone," + currentUser.GetStringID(),
			"not_in:" + currentUser.Phone,
		},
		"verify_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
			"not_exists:手机号已被占用",
			"not_in:新的手机与老手机号一致",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*UserUpdatePhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}
