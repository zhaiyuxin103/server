package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"server/pkg/app"
	"server/pkg/database"
	"server/pkg/paginator"
)

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// GetByPhone 通过手机号来获取用户
func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}

// GetByMulti 通过 手机号 / Email / 用户名 来获取用户
func GetByMulti(username string) (userModel User) {
	database.DB.
		Where("phone = ?", username).
		Or("email = ?", username).
		First(&userModel)
	return
}

// Get 通过 ID 获取用户
func Get(id string, loading bool) (user User) {
	if loading {
		database.DB.Preload(clause.Associations).Where("id", id).First(&user)
	} else {
		database.DB.Where("id", id).First(&user)
	}
	return
}

// GetByEmail 通过 Email 来获取用户
func GetByEmail(email string) (userModel User) {
	database.DB.Where("email = ?", email).First(&userModel)
	return
}

// All 获取所有用户数据
func All() (users []User) {
	database.DB.Find(&users)
	return
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	query := database.DB.Model(User{})
	if c.Query("id") != "" {
		query = query.Where("id = ?", c.Query("id"))
	}
	if c.Query("last_name") != "" {
		query = query.Where("last_name like ?", "%"+c.Query("last_name")+"%")
	}
	if c.Query("first_name") != "" {
		query = query.Where("first_name like ?", "%"+c.Query("first_name")+"%")
	}
	if c.Query("last_kana") != "" {
		query = query.Where("last_kana like ?", "%"+c.Query("last_kana")+"%")
	}
	if c.Query("first_kana") != "" {
		query = query.Where("first_kana like ?", "%"+c.Query("first_kana")+"%")
	}
	// TODO: 生日
	if c.Query("gender") != "" {
		query = query.Where("gender = ?", c.Query("gender"))
	}
	if c.Query("email") != "" {
		query = query.Where("email like ?", "%"+c.Query("email")+"%")
	}
	if c.Query("phone") != "" {
		query = query.Where("phone like ?", "%"+c.Query("phone")+"%")
	}
	if c.Query("introduction") != "" {
		query = query.Where("introduction like ?", "%"+c.Query("introduction")+"%")
	}
	if c.Query("state") != "" {
		query = query.Where("state = ?", c.Query("state"))
	}
	paging = paginator.Paginate(
		c,
		query,
		&users,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)
	return
}
