package topic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"server/pkg/app"
	"server/pkg/database"
	"server/pkg/paginator"
)

func Get(id string, loading bool) (topic Topic) {
	if loading {
		database.DB.Preload(clause.Associations).Where("id", id).First(&topic)
	} else {
		database.DB.Where("id", id).First(&topic)
	}
	return
}

func GetBy(field, value string) (topic Topic) {
	database.DB.Where("? = ?", field, value).First(&topic)
	return
}

func All() (topics []Topic) {
	database.DB.Find(&topics)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Topic{}).Where(fmt.Sprintf("%v = ?", field), value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (topics []Topic, paging paginator.Paging) {
	query := database.DB.Model(Topic{})
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
		&topics,
		app.V1URL(database.TableName(&Topic{})),
		perPage,
	)
	return
}
