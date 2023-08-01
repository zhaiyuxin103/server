package category

import (
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"server/pkg/helpers"
	"server/pkg/logger"
)

// BeforeSave GORM 的模型钩子，在创建和更新模型前调用
func (category *Category) BeforeSave(tx *gorm.DB) (err error) {

	if helpers.Empty(category.ParentID) {
		// 将层级设为 0
		category.Level = 0
		// 将 path 设为 -
		category.Path = "-"
	} else {
		var parent Category
		err := tx.First(&parent, category.ParentID).Error
		if err != nil {
			logger.LogIf(err)
		}
		// 将层级设为父类目的层级 + 1
		category.Level = parent.Level + 1
		// 将 path 值设为父类目的 path 追加父类目 ID 以及最后跟上一个 - 分隔符
		category.Path = parent.Path + cast.ToString(category.ParentID) + "-"
	}
	return
}

// func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {}
// func (category *Category) AfterCreate(tx *gorm.DB) (err error) {}
// func (category *Category) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (category *Category) AfterUpdate(tx *gorm.DB) (err error) {}
// func (category *Category) AfterSave(tx *gorm.DB) (err error) {}
// func (category *Category) BeforeDelete(tx *gorm.DB) (err error) {}
// func (category *Category) AfterDelete(tx *gorm.DB) (err error) {}
// func (category *Category) AfterFind(tx *gorm.DB) (err error) {}
