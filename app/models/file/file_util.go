package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"server/pkg/app"
	"server/pkg/database"
	"server/pkg/paginator"
)

func Get(id string, loading bool) (file File) {
	if loading {
		database.DB.Preload(clause.Associations).Where("id", id).First(&file)
	} else {
		database.DB.Where("id", id).First(&file)
	}
	return
}

func GetBy(field, value string) (file File) {
	database.DB.Where("? = ?", field, value).First(&file)
	return
}

func All() (files []File) {
	database.DB.Find(&files)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(File{}).Where(fmt.Sprintf("%v = ?", field), value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (files []File, paging paginator.Paging) {
	query := database.DB.Model(File{})
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
		&files,
		app.V1URL(database.TableName(&File{})),
		perPage,
	)
	return
}
