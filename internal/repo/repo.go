package repo

import (
	"github.com/lantonster/iam/config"
	"github.com/lantonster/iam/internal/dao"
)

type Repo interface {
	User() UserRepo
}

type defaultRepo struct {
	user UserRepo
}

func NewDefaultRepo(conf *config.Config) Repo {
	db, err := conf.MySQL.Connect()
	if err != nil {
		panic(err)
	}
	dao.SetDefault(db)

	return &defaultRepo{
		user: newDefaultUserRepo(),
	}
}

func (r *defaultRepo) User() UserRepo {
	return r.user
}
