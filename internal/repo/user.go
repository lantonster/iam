package repo

import (
	"context"
	"errors"

	"github.com/lantonster/cerrors"
	"github.com/lantonster/ecodes"
	"github.com/lantonster/iam/internal/dao"
	"github.com/lantonster/iam/internal/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	// 通过用户名获取用户
	GetUserByUsername(c context.Context, username string) (user *model.User, err error)

	// 检查用户名是否重复
	CheckUserNameDuplication(c context.Context, username string) (bool, error)
}

type defaultUserRepo struct{}

func newDefaultUserRepo() *defaultUserRepo {
	return &defaultUserRepo{}
}

// GetUserByUsername 通过用户名获取用户
func (r *defaultUserRepo) GetUserByUsername(c context.Context, username string) (user *model.User, err error) {
	u := dao.User

	if user, err = u.WithContext(c).Where(u.Username.Eq(username)).First(); err != nil {
		// ErrRecordNotFound 查询不到记录即用户名不存在
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cerrors.WithCode(ecodes.IAM_USERNAME_NOT_FOUND, "username %s not found", username)
		}

		return nil, cerrors.Wrap(err, "failed to get user by username")
	}

	return user, nil
}

// CheckUserNameDuplication 检查用户名是否重复
func (r *defaultUserRepo) CheckUserNameDuplication(c context.Context, username string) (bool, error) {
	u := dao.User

	count, err := u.WithContext(c).Where(u.Username.Eq(username)).Count()
	if err != nil {
		return false, cerrors.Wrap(err, "failed to check user name duplication")
	}

	return count > 0, nil
}
