// Package routes 注册路由
package routes

import (
	"net/http"
	api "server/app/http/controllers/api/v1"
	"server/app/http/controllers/api/v1/auth"
	"server/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册 API 相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		// 注册一个路由
		v1.GET("/", func(c *gin.Context) {
			// 以 JSON 格式响应
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"code":    http.StatusOK,
				"message": "Hello API.",
				"data":    gin.H{},
				"error":   gin.H{},
			})
		})

		// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
		// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
		// 测试时，可以调高一点。
		v1.Use(middlewares.LimitIP("200-H"))

		{
			authGroup := v1.Group("/auth")
			// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
			// 测试时，可以调高一点
			authGroup.Use(middlewares.LimitIP("1000-H"))
			{
				// 注册用户
				suc := new(auth.SignupController)
				// 判断手机是否已注册
				authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), suc.IsPhoneExist)
				// 判断 Email 是否已注册
				authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), suc.IsEmailExist)
				authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
				authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)

				// 发送验证码
				vcc := new(auth.VerifyCodeController)
				// 图片验证码，需要加限流
				authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("20-H"), vcc.ShowCaptcha)
				authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
				authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("50-H"), vcc.SendUsingEmail)

				// 登录
				lgc := new(auth.LoginController)
				// 使用手机号，短信验证码进行登录
				authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
				// 支持手机号 和 Email
				authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
				authGroup.POST("/login/refresh-token", middlewares.AuthJWT("api"), lgc.RefreshToken)

				// 重置密码
				pwc := new(auth.PasswordController)
				authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone)
				authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)

				uc := new(api.UsersController)

				// 获取当前用户
				v1.GET("/user", middlewares.AuthJWT("api"), uc.CurrentUser)
				usersGroup := v1.Group("/users")
				{
					usersGroup.GET("", uc.Index)
				}

				cgc := new(api.CategoriesController)
				cgcGroup := v1.Group("/categories")
				{
					cgcGroup.GET("", cgc.Index)
					cgcGroup.POST("", middlewares.AuthJWT("api"), cgc.Store)
					cgcGroup.PUT("/:id", middlewares.AuthJWT("api"), cgc.Update)
				}
			}
		}
	}
}
