package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type CategoryRequest struct {
	Name        string `valid:"name" json:"name" form:"name"`
	ParentID    uint64 `valid:"parent_id" json:"parent_id" form:"parent_id"`
	IsDirectory bool   `valid:"is_directory" json:"is_directory" form:"is_directory"`
	Description string `valid:"description" json:"description,omitempty" form:"description"`
	State       bool   `valid:"state" json:"state" form:"state"`
	Order       uint64 `valid:"order" json:"order" form:"order"`
}

func StoreCategory(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":         []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name"},
		"parent_id":    []string{"numeric", "exists:categories,id"},
		"is_directory": []string{"bool"},
		"description":  []string{"min_cn:3", "max_cn:255"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:分类名称为必填项",
			"min_cn:分类名称长度需至少 2 个字",
			"max_cn:分类名称长度不能超过 8 个字",
			"not_exists:分类名称已存在",
		},
		"parent_id": []string{
			"numeric:父级分类格式不正确",
			"exists:父级分类不存在",
		},
		"is_directory": []string{
			"bool:是否为目录格式不正确",
		},
		"description": []string{
			"min_cn:分类描述长度需至少 3 个字",
			"max_cn:分类描述长度不能超过 255 个字",
		},
	}
	return validate(data, rules, messages)
}

func UpdateCategory(data interface{}, c *gin.Context) map[string][]string {

	id := c.Param("id")

	rules := govalidator.MapData{
		"name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name," + id},
		"description": []string{"min_cn:3", "max_cn:255"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:分类名称为必填项",
			"min_cn:分类名称长度需至少 2 个字",
			"max_cn:分类名称长度不能超过 8 个字",
			"not_exists:分类名称已存在",
		},
		"description": []string{
			"min_cn:分类描述长度需至少 3 个字",
			"max_cn:分类描述长度不能超过 255 个字",
		},
	}
	return validate(data, rules, messages)
}
