package migrations

import (
	"database/sql"
	"gorm.io/gorm"
	"server/app/models"
	"server/pkg/migrate"
)

func init() {

	type Category struct {
		models.BaseModel

		Name     string `gorm:"column:name;type:varchar(255);not null;index;comment:名称"`
		ParentID uint64 `gorm:"column:parent_id;type:bigint(20) unsigned;default:null;index;comment:父类目ID"`
		// IsDirectory 是否为父类目
		IsDirectory bool `gorm:"column:is_directory;type:tinyint(1) unsigned;not null;default:0;comment:是否为父类目"`
		// Level 级别
		Level uint8 `gorm:"column:level;type:tinyint(4) unsigned;not null;default:0;comment:级别"`
		// Path 级别路径
		Path string `gorm:"column:path;type:varchar(255);not null;default:'-';comment:级别路径"`
		// Description 描述
		Description string `gorm:"column:description;type:varchar(255);default:null;comment:描述"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.AutoMigrate(&Category{})
		if err != nil {
			return
		}
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.DropTable(&Category{})
		if err != nil {
			return
		}
	}

	migrate.Add("2023_08_01_200552_add_categories_table", up, down)
}
