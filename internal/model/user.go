package model

import (
	"github.com/lantonster/iam/pkg/utils"
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

// ComparePassword 验证密码是否正确
func (u *User) ComparePassword(password string) bool {
	return utils.VerifyPassword(password, u.Salt, u.Password)
}
