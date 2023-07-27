// Package middlewares Gin 中间件
package middlewares

import (
	"fmt"
	"github.com/spf13/cast"
	"server/app/models/user"
	"server/pkg/config"
	"server/pkg/jwt"
	"server/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthJWT(guard string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParserToken(c)

		// JWT 解析失败，有错误发生
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}

		// JWT 解析成功，设置用户信息
		switch guard {
		case "api":
			userModel := user.Get(cast.ToString(claims.User.ID), false)
			if userModel.ID == 0 {
				response.Unauthorized(c, "找不到对应用户，用户可能已删除")
				return
			}
			// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
			c.Set("user_id", userModel.GetStringID())
			c.Set("user_name", userModel.LastName+userModel.FirstName)
			c.Set("user", userModel)
		default:
			response.Unauthorized(c, "无效的守卫")
		}

		c.Next()
	}
}
