package repo

import "gorm.io/gorm"

type Repo struct {
	*Query
}

func NewRepo(db *gorm.DB) *Repo {
	SetDefault(db)
	return &Repo{
		Query: Q,
	}
}
