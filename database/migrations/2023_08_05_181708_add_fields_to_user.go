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

		Province string `gorm:"column:province;type:varchar(255);default:null;comment:省份"`
		City     string `gorm:"column:city;type:varchar(255);default:null;comment:城市"`
		District string `gorm:"column:district;type:varchar(255);default:null;comment:区县"`
		Address  string `gorm:"column:address;type:varchar(255);default:null;comment:地址"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.AutoMigrate(&User{})
		if err != nil {
			return
		}
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.DropColumn(&User{}, "Province")
		if err != nil {
			return
		}
		err = migrator.DropColumn(&User{}, "City")
		if err != nil {
			return
		}
		err = migrator.DropColumn(&User{}, "District")
		if err != nil {
			return
		}
		err = migrator.DropColumn(&User{}, "Address")
		if err != nil {
			return
		}
	}

	migrate.Add("2023_08_05_181708_add_fields_to_user", up, down)
}
