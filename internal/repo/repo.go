package repo

import (
	"github.com/lantonster/corekit"
	"github.com/lantonster/iam/config"
	"github.com/lantonster/iam/internal/dao"
)

type Repo struct {
	User UserRepo
}

func NewRepo(conf *config.Config) *Repo {
	db := corekit.ConnectMySQL(conf.MySQL)
	dao.SetDefault(db)

	return &Repo{
		User: newDefaultUserRepo(),
	}
}
