package model

import "gorm.io/plugin/soft_delete"

type Model struct {
	Id        int64                 `json:"id" gorm:"primarykey"`
	CreatedAt int64                 `json:"createdAt"`
	UpdatedAt int64                 `json:"updatedAt"`
	DeletedAt soft_delete.DeletedAt `json:"-"`
}
