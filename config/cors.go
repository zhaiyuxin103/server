package config

import "server/pkg/config"

func init() {
	config.Add("cors", func() map[string]interface{} {
		return map[string]interface{}{

			"allowed_methods": config.Env("CORS_ALLOW_METHODS", "GET，POST，PUT，PATCH，UPDATE，DELETE，HEAD，OPTIONS"),

			"allowed_headers": config.Env("CORS_ALLOW_HEADERS", "Content-Type，Origin，Authorization，X-Requested-With，Accept-Language，Accept"),

			"exposed_headers": config.Env("CORS_EXPOSED_HEADERS，Content-Length，Access-Control-Allow-Origin，Access-Control-Allow-Headers，Cache-Control，Content-Language，Content-Type"),

			"allowed_origins": config.Env("CORS_ALLOW_ORIGINS", "*"),

			"allowed_origins_regexp": config.Env("CORS_ALLOW_ORIGINS_REGEXP", ""),

			"allow_credentials": config.Env("CORS_ALLOW_CREDENTIALS", true),

			"max_age": config.Env("CORS_MAX_AGE", 12),
		}
	})
}
