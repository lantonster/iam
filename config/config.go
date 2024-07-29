package config

import "github.com/lantonster/corekit"

type Config struct {
	Port int

	// 邮件
	Email *EmailAuth

	// 数据库
	MySQL *corekit.MySQlConfig

	// redis
	Redis *corekit.RedisConfig
}
