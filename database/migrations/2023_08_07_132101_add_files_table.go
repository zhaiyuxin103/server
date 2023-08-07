package migrations

import (
	"database/sql"
	"gorm.io/gorm"
	"server/app/models"
	"server/pkg/migrate"
)

func init() {

	type File struct {
		models.BaseModel

		Type   string `gorm:"column:type;type:varchar(255);not null;index;comment:类型"`
		UserID uint64 `gorm:"column:user_id;type:bigint unsigned;not null;index;comment:用户 ID"`
		Folder string `gorm:"column:folder;type:varchar(255);not null;index;comment:文件夹"`
		Path   string `gorm:"column:path;type:varchar(255);not null;index;comment:路径"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.AutoMigrate(&File{})
		if err != nil {
			return
		}
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.DropTable(&File{})
		if err != nil {
			return
		}
	}

	migrate.Add("2023_08_07_132101_add_files_table", up, down)
}
