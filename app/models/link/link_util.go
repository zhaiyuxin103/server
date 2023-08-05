package link

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"server/pkg/app"
	"server/pkg/cache"
	"server/pkg/database"
	"server/pkg/helpers"
	"server/pkg/paginator"
	"time"
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

func AllCached() (links []Link) {
	// 设置缓存 key
	cacheKey := "links:all"
	// 设置过期时间
	expireTime := 120 * time.Minute
	// 取数据
	cache.GetObject(cacheKey, &links)

	// 如果数据为空
	if helpers.Empty(links) {
		// 查询数据库
		links = All()
		if helpers.Empty(links) {
			return links
		}
		// 设置缓存
		cache.Set(cacheKey, links, expireTime)
	}
	return
}
