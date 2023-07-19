// Package routes 注册路由
package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *gin.Engine) {

	web := r.Group("/")
	{
		// 注册一个路由
		web.GET("/", func(c *gin.Context) {
			// 以 JSON 格式响应
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"code":    http.StatusOK,
				"message": "Hello Web.",
				"data":    gin.H{},
				"error":   gin.H{},
			})
		})
	}
}
