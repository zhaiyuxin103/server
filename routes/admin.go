// Package routes 注册路由
package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterAdminRoutes 注册网页相关路由
func RegisterAdminRoutes(r *gin.Engine) {

	admin := r.Group("/admin")
	{
		// 注册一个路由
		admin.GET("/", func(c *gin.Context) {
			// 以 JSON 格式响应
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"code":    http.StatusOK,
				"message": "Hello Admin.",
				"data":    gin.H{},
				"error":   gin.H{},
			})
		})
	}
}
