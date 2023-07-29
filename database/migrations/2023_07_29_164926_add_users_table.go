package migrations

import (
	"database/sql"
	"gorm.io/gorm"
	"server/app/models"
	"server/pkg/migrate"
)

func init() {

	type User struct {
		models.BaseModel

		LastName          string `gorm:"column:last_name;type:varchar(255);not null;comment:姓"`
		FirstName         string `gorm:"column:first_name;type:varchar(255);not null;comment:名"`
		LastKana          string `gorm:"column:last_kana;type:varchar(255);not null;comment:姓(カナ)"`
		FirstKana         string `gorm:"column:first_kana;type:varchar(255);not null;comment:名(カナ)"`
		Birthday          string `gorm:"column:birthday;type:date;not null;comment:生年月日"`
		AvatarID          uint64 `gorm:"column:avatar_id;type:bigint;not null;comment:头像ID"`
		Gender            uint8  `gorm:"column:gender;type:tinyint(1);not null;index;comment:性别 1:男性 2:女性"`
		Email             string `gorm:"column:email;type:varchar(255);index;default:null;comment:邮箱"`
		EmailVerifiedAt   string `gorm:"column:email_verified_at;type:datetime;default:null;comment:邮箱验证时间"`
		Phone             string `gorm:"column:phone;type:varchar(20);index;default:null;comment:手机号"`
		Password          string `gorm:"column:password;type:varchar(255);comment:密码"`
		Introduction      string `gorm:"column:introduction;type:varchar(255);default:null;comment:自我介绍"`
		NotificationCount uint64 `gorm:"column:notification_count;type:bigint;not null;default:0;comment:通知数量"`
		LastActivedAt     string `gorm:"column:last_actived_at;type:datetime;index;default:null;comment:最后活跃时间"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.AutoMigrate(&User{})
		if err != nil {
			return
		}
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.DropTable(&User{})
		if err != nil {
			return
		}
	}

	migrate.Add("2023_07_29_164926_add_users_table", up, down)
}
