package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"server/pkg/app"
	"server/pkg/database"
	"server/pkg/paginator"
)

func Get(id string, loading bool) (category Category) {
	if loading {
		database.DB.Preload(clause.Associations).Where("id", id).First(&category)
	} else {
		database.DB.Where("id", id).First(&category)
	}
	return
}

func GetBy(field, value string) (category Category) {
	database.DB.Where("? = ?", field, value).First(&category)
	return
}

func All() (categories []Category) {
	database.DB.Find(&categories)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Category{}).Where(fmt.Sprintf("%v = ?", field), value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (categories []Category, paging paginator.Paging) {
	query := database.DB.Model(Category{})
	if c.Query("id") != "" {
		query = query.Where("id = ?", c.Query("id"))
	}
	if c.Query("name") != "" {
		query = query.Where("name like ?", "%"+c.Query("name")+"%")
	}
	if c.Query("state") != "" {
		query = query.Where("state = ?", c.Query("state"))
	}
	paging = paginator.Paginate(
		c,
		query,
		&categories,
		app.V1URL(database.TableName(&Category{})),
		perPage,
	)
	return
}

// HasChildren 检查是否含有子类目
func (category *Category) HasChildren() bool {
	var count int64
	database.DB.Model(Category{}).Where("parent_id = ?", category.ID).Count(&count)
	return count > 0
}
