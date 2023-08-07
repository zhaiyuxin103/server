// Package file 模型
package file

import (
	"server/app/models"
	"server/pkg/database"
)

type File struct {
	models.BaseModel

	Type   string `json:"type"`
	UserID uint64 `json:"user_id"`
	Folder string `json:"folder"`
	Path   string `json:"path"`

	models.CommonTimestampsField
}

func (file *File) Create() {
	database.DB.Create(&file)
}

func (file *File) Save() (rowsAffected int64) {
	result := database.DB.Save(&file)
	return result.RowsAffected
}

func (file *File) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&file)
	return result.RowsAffected
}
