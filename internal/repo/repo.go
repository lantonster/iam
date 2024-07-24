package repo

import (
	"github.com/lantonster/corekit"
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
	db := corekit.ConnectMySQL(conf.MySQL)
	dao.SetDefault(db)

	return &defaultRepo{
		user: newDefaultUserRepo(),
	}
}

func (r *defaultRepo) User() UserRepo {
	return r.user
}
