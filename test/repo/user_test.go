package repo

import (
	"context"
	"testing"

	"github.com/lantonster/cerrors"
	"github.com/lantonster/ecodes"
	"github.com/stretchr/testify/assert"
)

func TestUserRepo_GetUserByUsername(t *testing.T) {
	repo := r.User()
	ctx := context.Background()

	// 测试存在用户的情况
	user, err := repo.GetUserByUsername(ctx, testUser.Username)
	assert.NoError(t, err)
	assert.NotNil(t, user)

	// 测试不存在用户的情况
	user, err = repo.GetUserByUsername(ctx, "non_existent_user")
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, ecodes.IAM_USERNAME_NOT_FOUND, cerrors.Code(err))
}
