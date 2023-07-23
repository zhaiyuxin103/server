// Package user 存放用户 Model 相关逻辑
package user

import (
	"github.com/golang-module/carbon/v2"
	"server/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel

	LastName          string          `json:"last_name"`
	FirstName         string          `json:"first_name"`
	LastKana          string          `json:"last_kana"`
	FirstKana         string          `json:"first_kana"`
	Birthday          carbon.Date     `json:"birthday"`
	AvatarID          uint            `json:"avatar_id"`
	Gender            uint8           `json:"gender"`
	Email             string          `json:"email"`
	EmailVerifiedAt   carbon.DateTime `json:"email_verified_at"`
	Phone             string          `json:"phone"`
	Password          string          `json:"-"`
	Introduction      string          `json:"introduction"`
	NotificationCount uint64          `json:"notification_count"`
	LastActivedAt     carbon.DateTime `json:"last_actived_at"`

	// TODO: 通过 avatar_id 关联用户
	// Avatar image.Image `gorm:"foreignKey:avatar_id" json:"image"`

	models.CommonTimestampsField
}
