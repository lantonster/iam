package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/lantonster/iam/pkg/utils"
)

const (
	// 邮箱验证码: %s 为邮箱, 值为验证码
	EmailVerificationCodeKey = "iam:verification_code:email:%s" // 邮箱验证码
	EmailVerificationExp     = 5 * 60 * time.Second             // 验证码过期时间: 5 分钟
)

type VerificationCodeRepo interface {
	// 生成验证码并保存到 redis
	GenerateCode(c context.Context, email string) (code string, err error)
}

type defaultVerificationCodeRepo struct{}

func newDefaultVerificationCodeRepo() *defaultVerificationCodeRepo {
	return &defaultVerificationCodeRepo{}
}

func (r *defaultVerificationCodeRepo) GenerateCode(c context.Context, email string) (code string, err error) {
	code = utils.GenerateVerificationCode()

	key := fmt.Sprintf(EmailVerificationCodeKey, email)
	if err := rdb.Set(c, key, code, EmailVerificationExp).Err(); err != nil {
		return "", err
	}

	return code, nil
}
