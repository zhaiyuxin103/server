// Package config 站点配置信息
package config

import "server/pkg/config"

func init() {
	config.Add("file", func() map[string]interface{} {
		return map[string]interface{}{

			// Default Filesystem Disk
			"default": config.Env("FILE_DISK", "local"),

			// Filesystem Disks
			"disks": map[string]interface{}{
				"local": map[string]interface{}{
					"driver": "local",
					"root":   "public",
					"throw":  false,
				},

				"s3": map[string]interface{}{
					"driver":                  "s3",
					"key":                     config.Env("AWS_ACCESS_KEY_ID"),
					"secret":                  config.Env("AWS_SECRET_ACCESS_KEY"),
					"region":                  config.Env("AWS_DEFAULT_REGION"),
					"bucket":                  config.Env("AWS_BUCKET"),
					"url":                     config.Env("AWS_URL"),
					"endpoint":                config.Env("AWS_ENDPOINT"),
					"use_path_style_endpoint": config.Env("AWS_USE_PATH_STYLE_ENDPOINT", false),
					"throw":                   false,
				}},
		}
	})
}
