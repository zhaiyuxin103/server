package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type TopicRequest struct {
	Title      string `json:"title,omitempty" valid:"title" form:"title"`
	SubTitle   string `json:"sub_title,omitempty" valid:"sub_title" form:"sub_title"`
	CategoryID string `json:"category_id,omitempty" valid:"category_id" form:"category_id"`
	Content    string `json:"content,omitempty" valid:"content" form:"content"`
	State      bool   `valid:"state" json:"state" form:"state"`
	Order      uint64 `valid:"order" json:"order" form:"order"`
}

func StoreTopic(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"title":       []string{"required", "min_cn:3", "max_cn:40"},
		"sub_title":   []string{"required", "min_cn:3", "max_cn:255"},
		"category_id": []string{"required", "exists:categories,id"},
		"content":     []string{"required", "min_cn:10", "max_cn:50000"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"required:帖子标题为必填项",
			"min_cn:标题长度需大于 3",
			"max_cn:标题长度需小于 40",
		},
		"sub_title": []string{
			"required:帖子副标题为必填项",
			"min_cn:副标题长度需大于 3",
			"max_cn:副标题长度需小于 40",
		},
		"content": []string{
			"required:帖子内容为必填项",
			"min_cn:长度需大于 10",
		},
		"category_id": []string{
			"required:帖子分类为必填项",
			"exists:帖子分类未找到",
		},
	}
	return validate(data, rules, messages)
}

func UpdateTopic(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// "name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:topics,name"},
		// "description": []string{"min_cn:3", "max_cn:255"},
	}
	messages := govalidator.MapData{
		// "name": []string{
		//     "required:名称为必填项",
		//     "min_cn:名称长度需至少 2 个字",
		//     "max_cn:名称长度不能超过 8 个字",
		//     "not_exists:名称已存在",
		// },
		// "description": []string{
		//     "min_cn:描述长度需至少 3 个字",
		//     "max_cn:描述长度不能超过 255 个字",
		// },
	}
	return validate(data, rules, messages)
}
