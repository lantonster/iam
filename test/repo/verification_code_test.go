package repo

import (
	"context"
	"fmt"
	"testing"

	"github.com/lantonster/iam/internal/repo"
	"github.com/stretchr/testify/assert"
)

func TestVerificationCodeRepo_GenerateCode(t *testing.T) {
	re := r.VerificationCode()

	t.Run("测试验证码是否写入 redis", func(t *testing.T) {
		email := "example@example.com"

		key := fmt.Sprintf(repo.EmailVerificationCodeKey, email)
		code, err := re.GenerateCode(context.Background(), email)
		getCode := rdb.Get(context.Background(), key).Val()

		assert.NoError(t, err)
		assert.NotEqual(t, "", code)
		assert.Equal(t, getCode, code)
	})
}
