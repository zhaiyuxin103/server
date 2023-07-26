// Package models 模型通用属性和方法
package models

import (
	"github.com/golang-module/carbon/v2"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;comment:ID" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	State     uint8           `gorm:"column:state;default:0;comment:状态" json:"state"`
	Order     uint64          `gorm:"column:order;default:0;comment:排序" json:"order"`
	CreatedAt carbon.DateTime `gorm:"column:created_at;type:timestamp;index;comment:创建时间" json:"created_at,omitempty"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at;type:timestamp;index;comment:最后编辑时间" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at;type:timestamp;index;comment:删除时间" json:"deleted_at,omitempty"`
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}
