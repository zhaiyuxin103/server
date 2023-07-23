// Package config 站点配置信息
package config

import "server/pkg/config"

func init() {
	config.Add("mail", func() map[string]interface{} {
		return map[string]interface{}{
			// Default Mailer
			"default": config.Env("MAIL_MAILER", "smtp"),
			// Mailer Configurations
			"mailers": map[string]interface{}{
				"smtp": map[string]interface{}{
					"transport":  "smtp",
					"host":       config.Env("MAIL_HOST", "localhost"),
					"port":       config.Env("MAIL_PORT", 1025),
					"encryption": config.Env("MAIL_ENCRYPTION", "tls"),
					"username":   config.Env("MAIL_USERNAME"),
					"password":   config.Env("MAIL_PASSWORD"),
					"timeout":    nil,
					"auth_mode":  nil,
				},
				"ses": map[string]interface{}{
					"transport": "ses",
					"key":       config.Env("AWS_ACCESS_KEY_ID"),
					"secret":    config.Env("AWS_SECRET_ACCESS_KEY"),
					"region":    config.Env("AWS_SES_REGION"),
				},
				"mailgun": map[string]interface{}{
					"transport": "mailgun",
				},
				"postmark": map[string]interface{}{
					"transport": "postmark",
				},
				"sendmail": map[string]interface{}{
					"transport": "sendmail",
					"path":      "/usr/sbin/sendmail -bs",
				},
				"log": map[string]interface{}{
					"transport": "log",
					"channel":   config.Env("MAIL_LOG_CHANNEL"),
				},
				"array": map[string]interface{}{
					"transport": "array",
				},
				"failover": map[string]interface{}{
					"transport": "failover",
					"mailers": []string{
						"smtp",
						"log",
					},
				},
			},
			// Global "From" Address
			"from": map[string]interface{}{
				"address": config.Env("MAIL_FROM_ADDRESS", "hello@example.com"),
				"name":    config.Env("MAIL_FROM_NAME", "Example"),
			},
			// Markdown Mail Settings
			"markdown": map[string]interface{}{
				"theme": "default",
				"paths": []string{
					"resources/views/vendor/mail",
				},
			},
		}
	})
}
