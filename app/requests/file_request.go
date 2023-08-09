package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"mime/multipart"
)

type FileRequest struct {
	Type   string                `valid:"type" json:"type" form:"type"`
	Folder string                `valid:"folder" json:"folder" form:"folder"`
	File   *multipart.FileHeader `valid:"file" form:"file" form:"file"`
	State  bool                  `valid:"state" json:"state" form:"state"`
	Order  uint64                `valid:"order" json:"order" form:"order"`
}

func StoreFile(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"type":   []string{"required", "in:image"},
		"folder": []string{"required", "in:avatar,topic"},
		// size 的单位为 bytes
		// - 1024 bytes 为 1kb
		// - 1048576 bytes 为 1mb
		// - 20971520 bytes 为 20mb
		"file:file": []string{"ext:png,jpg,jpeg,pdf", "size:20971520", "required"},
	}
	messages := govalidator.MapData{
		"type": []string{
			"required:类型为必填项",
			"in:类型必须为 image",
		},
		"folder": []string{
			"required:文件夹为必填项",
			"in:文件夹必须为 avatar 或 topic",
		},
		"file:file": []string{
			"ext:ext头像只能上传 png, jpg, jpeg 任意一种的图片",
			"size:头像文件最大不能超过 20MB",
			"required:必须上传图片",
		},
	}
	return validateFile(c, data, rules, messages)
}

func UpdateFile(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// "name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:files,name"},
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
	return validateFile(c, data, rules, messages)
}
