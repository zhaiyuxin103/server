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
		"title":       []string{"min_cn:3", "max_cn:40"},
		"sub_title":   []string{"min_cn:3", "max_cn:255"},
		"category_id": []string{"exists:categories,id"},
		"content":     []string{"min_cn:10", "max_cn:50000"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"min_cn:标题长度需大于 3",
			"max_cn:标题长度需小于 40",
		},
		"sub_title": []string{
			"min_cn:副标题长度需大于 3",
			"max_cn:副标题长度需小于 40",
		},
		"content": []string{
			"min_cn:长度需大于 10",
		},
		"category_id": []string{
			"exists:帖子分类未找到",
		},
	}
	return validate(data, rules, messages)
}
