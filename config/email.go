package config

type EmailAuth struct {
	SmtpPort         int              `mapstructure:"smtp_port"`
	SmtpHost         string           `mapstructure:"smtp_host"`
	SmtpUsername     string           `mapstructure:"smtp_username"`
	SmtpPassword     string           `mapstructure:"smtp_password"`
	VerificationCode VerificationCode `mapstructure:"verification_code"`
}

type VerificationCode struct {
	Sender     string `mapstructure:"sender"`
	Subject    string `mapstructure:"subject"`
	Content    string `mapstructure:"content"`
	ExpireTime int    `mapstructure:"expire_time"` // 验证码过期时间 单位：秒
}
