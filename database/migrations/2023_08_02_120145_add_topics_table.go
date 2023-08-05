package migrations

import (
	"database/sql"
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
	"server/app/models"
	"server/pkg/migrate"
)

func init() {

	type User struct {
		models.BaseModel
	}
	type Category struct {
		models.BaseModel
	}

	type Topic struct {
		models.BaseModel

		Title           string      `gorm:"column:title;type:varchar(255);not null;comment:标题"`
		SubTitle        string      `gorm:"column:sub_title;type:varchar(255);default:null;comment:副标题"`
		UserID          uint64      `gorm:"column:user_id;type:bigint(20) unsigned;not null;index;comment:用户 ID"`
		CategoryID      uint64      `gorm:"column:category_id;type:bigint(20) unsigned;not null;index;comment:类目 ID"`
		Content         string      `gorm:"column:content;type:text;not null;comment:内容"`
		IsRelease       uint8       `gorm:"column:is_release;type:tinyint(1) unsigned;default:0;comment:是否发布"`
		ReleasedAt      carbon.Date `gorm:"column:released_at;type:date;default:null;comment:发布时间"`
		VoteCount       uint64      `gorm:"column:vote_count;type:int(11) unsigned;default:0;comment:点赞数"`
		UnvoteCount     uint64      `gorm:"column:unvote_count;type:int(11) unsigned;default:0;comment:踩数"`
		ViewCount       uint64      `gorm:"column:view_count;type:int(11) unsigned;default:0;comment:浏览数"`
		FavoriteCount   uint64      `gorm:"column:favorite_count;type:int(11) unsigned;default:0;comment:收藏数"`
		ShareCount      uint64      `gorm:"column:share_count;type:int(11) unsigned;default:0;comment:分享数"`
		ReplyCount      uint64      `gorm:"column:reply_count;type:int(11) unsigned;default:0;comment:回复数"`
		LastReplyUserID uint64      `gorm:"column:last_reply_user_id;type:bigint(20) unsigned;default:null;comment:最后回复用户 ID"`
		Excerpt         string      `gorm:"column:excerpt;type:varchar(255);default:null;comment:摘要"`
		Slug            string      `gorm:"column:slug;type:varchar(255);default:null;comment:SEO 友好的 URI"`

		// 会创建 user_id 和 category_id 外键的约束
		User     User
		Category Category

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.AutoMigrate(&Topic{})
		if err != nil {
			return
		}
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.DropTable(&Topic{})
		if err != nil {
			return
		}
	}

	migrate.Add("2023_08_02_120145_add_topics_table", up, down)
}
