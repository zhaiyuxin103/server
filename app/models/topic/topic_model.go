// Package topic 模型
package topic

import (
	"server/app/models"
	"server/app/models/category"
	"server/app/models/user"
	"server/pkg/database"
)

type Topic struct {
	models.BaseModel

	Title    string `json:"title,omitempty"`
	SubTitle string `json:"sub_title,omitempty"`
	// TODO: 图片
	IsRelease       bool   `json:"is_release,omitempty"`
	ReleasedAt      string `json:"released_at,omitempty"`
	VoteCount       uint64 `json:"vote_count,omitempty"`
	UnvoteCount     uint64 `json:"unvote_count,omitempty"`
	ViewCount       uint64 `json:"view_count,omitempty"`
	FavoriteCount   uint64 `json:"favorite_count,omitempty"`
	ShareCount      uint64 `json:"share_count,omitempty"`
	ReplyCount      uint64 `json:"reply_count,omitempty"`
	LastReplyUserID uint64 `json:"last_reply_user_id,omitempty"`
	Excerpt         string `json:"excerpt,omitempty"`
	Slug            string `json:"slug,omitempty"`
	Content         string `json:"content,omitempty"`
	UserId          uint64 `json:"user_id,omitempty"`
	CategoryId      uint64 `json:"category_id,omitempty"`

	// 通过 user_id 关联用户
	User user.User `json:"user"`

	// 通过 category_id 关联分类
	Category category.Category `json:"category"`

	models.CommonTimestampsField
}

func (topic *Topic) Create() {
	database.DB.Create(&topic)
}

func (topic *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&topic)
	return result.RowsAffected
}

func (topic *Topic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topic)
	return result.RowsAffected
}
