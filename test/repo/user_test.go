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

	t.Run("测试存在用户的情况", func(t *testing.T) {
		user, err := repo.GetUserByUsername(ctx, testUser.Username)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, testUser.Username, user.Username)
	})

	t.Run("测试不存在用户的情况", func(t *testing.T) {
		user, err := repo.GetUserByUsername(ctx, "non_existent_user")
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, ecodes.IAM_USERNAME_NOT_FOUND, cerrors.Code(err))
	})
}

func TestUserRepo_CheckUserNameDuplication(t *testing.T) {
	repo := r.User()
	ctx := context.Background()

	t.Run("测试用户名重复的情况", func(t *testing.T) {
		dup, err := repo.CheckUserNameDuplication(ctx, testUser.Username)
		assert.NoError(t, err)
		assert.True(t, dup)
	})

	t.Run("测试用户名不重复的情况", func(t *testing.T) {
		dup, err := repo.CheckUserNameDuplication(ctx, "non_existent_user")
		assert.NoError(t, err)
		assert.False(t, dup)
	})
}
