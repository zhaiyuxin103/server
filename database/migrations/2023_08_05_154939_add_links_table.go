package migrations

import (
	"database/sql"
	"gorm.io/gorm"
	"server/app/models"
	"server/pkg/migrate"
)

func init() {

	type Link struct {
		models.BaseModel

		Name        string `gorm:"column:name;type:varchar(255);not null"`
		ImageID     uint64 `gorm:"column:image_id;type:bigint(20);not null"`
		Description string `gorm:"column:description;type:text;default:null"`
		URI         string `gorm:"column:uri;type:varchar(255);default:null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.AutoMigrate(&Link{})
		if err != nil {
			return
		}
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.DropTable(&Link{})
		if err != nil {
			return
		}
	}

	migrate.Add("2023_08_05_154939_add_links_table", up, down)
}
