// Package requests 处理请求数据和表单验证
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"server/app/requests/validators"
	"time"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone" form:"phone"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email" form:"email"`
}

// SignupUsingPhoneRequest 通过手机注册的请求信息
type SignupUsingPhoneRequest struct {
	LastName        string    `json:"last_name,omitempty" valid:"last_name" form:"last_name"`
	FirstName       string    `json:"first_name,omitempty" valid:"first_name" form:"first_name"`
	LastKana        string    `json:"last_kana,omitempty" valid:"last_kana" form:"last_kana"`
	FirstKana       string    `json:"first_kana,omitempty" valid:"first_kana" form:"first_kana"`
	Birthday        time.Time `json:"birthday,omitempty" valid:"birthday" form:"birthday" time_format:"2006-01-02"`
	AvatarID        uint64    `json:"avatar_id,omitempty" valid:"avatar_id" form:"avatar_id"`
	Gender          uint8     `json:"gender,omitempty" valid:"gender" form:"gender"`
	Email           string    `json:"email,omitempty" valid:"email" form:"email"`
	Phone           string    `json:"phone,omitempty" valid:"phone" form:"phone"`
	VerifyCode      string    `json:"verify_code,omitempty" valid:"verify_code" form:"verify_code"`
	Password        string    `valid:"password" json:"password,omitempty" form:"password"`
	PasswordConfirm string    `valid:"password_confirm" json:"password_confirm,omitempty" form:"password_confirm"`
	Introduction    string    `valid:"introduction" json:"introduction,omitempty" form:"introduction"`
	State           uint8     `json:"state,omitempty" valid:"state" form:"state"`
	Order           uint64    `json:"order,omitempty" valid:"order" form:"order"`
}

func SignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}
	return validate(data, rules, messages)
}

func SignupEmailExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}
	return validate(data, rules, messages)
}

func SignupUsingPhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"last_name":        []string{"required", "between:1,20"},
		"first_name":       []string{"required", "between:1,20"},
		"last_kana":        []string{"between:1,20"},
		"first_kana":       []string{"between:1,20"},
		"birthday":         []string{"date"},
		"avatar_id":        []string{"numeric"},
		"gender":           []string{"in:1,2,3"},
		"email":            []string{"required", "email", "not_exists:users,email"},
		"phone":            []string{"required", "digits:11", "not_exists:users,phone"},
		"verify_code":      []string{"required", "digits:6"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
	}

	messages := govalidator.MapData{
		"last_name": []string{
			"required:姓氏为必填项",
			"between:姓氏长度需在 1~20 之间",
		},
		"first_name": []string{
			"required:名字为必填项",
			"between:名字长度需在 1~20 之间",
		},
		"last_kana": []string{
			"between:姓氏长度需在 1~20 之间",
		},
		"first_kana": []string{
			"between:名字长度需在 1~20 之间",
		},
		"birthday": []string{
			"date:生日格式错误，正确格式为 2006-01-02",
		},
		"avatar_id": []string{
			"numeric:头像 ID 必须为数字",
		},
		"gender": []string{
			"in:性别参数错误，只允许 1,2,3",
		},
		"email": []string{
			"required:Email 为必填项",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被注册",
		},
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*SignupUsingPhoneRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}
