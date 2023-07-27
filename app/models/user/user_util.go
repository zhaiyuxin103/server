package user

import (
	"gorm.io/gorm/clause"
	"server/pkg/database"
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
