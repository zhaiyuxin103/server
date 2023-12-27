package middlewares

import (
	"server/pkg/config"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     strings.Split(config.Get("cors.allowed_origins"), "，"),
		AllowMethods:     strings.Split(config.Get("cors.allowed_methods"), "，"),
		AllowHeaders:     strings.Split(config.Get("cors.allowed_headers"), "，"),
		ExposeHeaders:    strings.Split(config.Get("cors.exposed_headers"), "，"),
		AllowCredentials: config.GetBool("cors.allow_credentials"),
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: time.Duration(config.GetInt("cors.max_age", 12)) * time.Hour,
	})
}
