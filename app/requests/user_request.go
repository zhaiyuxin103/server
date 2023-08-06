package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
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
