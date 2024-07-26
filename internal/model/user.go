package model

import (
	"github.com/lantonster/cerrors"
	"github.com/lantonster/ecodes"
	"github.com/lantonster/iam/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	Id        int64                 `json:"id" gorm:"primarykey"`
	CreatedAt int64                 `json:"createdAt"`
	UpdatedAt int64                 `json:"updatedAt"`
	DeletedAt soft_delete.DeletedAt `json:"-"`

	Username string // 用户名
	Password string // 密码
	Salt     string // 密码盐
}

// BeforeCreate 在创建用户之前，检查用户名是否合法
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return IsUsernameValid(u.Username)
}

// ComparePassword 验证密码是否正确
func (u *User) ComparePassword(password string) bool {
	return utils.VerifyPassword(password, u.Salt, u.Password)
}

// IsUsernameValid 检查用户名是否合法: 长度在 1 到 30 之间，由字母、数字和下划线组成
func IsUsernameValid(username string) error {
	// 检查用户名长度，如果小于 1 或大于 30，则返回错误
	if l := len(username); l < 1 || 30 < l {
		return cerrors.WithCode(ecodes.IAM_INVALID_USERNAME_LENGTH, "the length of username should within the range of 1 to 30")
	}

	// 检查用户名格式，如果不符合要求（不是由字母、数字和下划线组成），则返回错误
	if !utils.ValidateUsernameFormat(username) {
		return cerrors.WithCode(ecodes.IAM_INVALID_USERNAME_FORMAT, "the username can only be composed of letters, numbers and underscores")
	}

	return nil
}
