package link

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"server/pkg/app"
	"server/pkg/database"
	"server/pkg/paginator"
)

func Get(id string, loading bool) (link Link) {
	if loading {
		database.DB.Preload(clause.Associations).Where("id", id).First(&link)
	} else {
		database.DB.Where("id", id).First(&link)
	}
	return
}

func GetBy(field, value string) (link Link) {
	database.DB.Where("? = ?", field, value).First(&link)
	return
}

func All() (links []Link) {
	database.DB.Find(&links)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Link{}).Where(fmt.Sprintf("%v = ?", field), value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
	query := database.DB.Model(Link{})
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
		&links,
		app.V1URL(database.TableName(&Link{})),
		perPage,
	)
	return
}
