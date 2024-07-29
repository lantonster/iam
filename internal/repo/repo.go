package repo

import (
	"github.com/go-redis/redis/v8"
	"github.com/lantonster/iam/config"
	"github.com/lantonster/iam/internal/dao"
)

var rdb *redis.Client

type Repo interface {
	User() UserRepo
	VerificationCode() VerificationCodeRepo
}

type defaultRepo struct {
	user             UserRepo
	verificationCode VerificationCodeRepo
}

// NewDefaultRepo 函数根据配置创建默认的仓储实例
func NewDefaultRepo(conf *config.Config) Repo {
	// 尝试根据配置中的 MySQL 信息连接数据库
	db, err := conf.MySQL.Connect()
	if err != nil {
		panic(err)
	}
	dao.SetDefault(db)

	// 尝试根据配置中的 Redis 信息连接 Redis
	rdb, err = conf.Redis.Connect()
	if err != nil {
		panic(err)
	}

	return &defaultRepo{
		user:             newDefaultUserRepo(),
		verificationCode: newDefaultVerificationCodeRepo(),
	}
}

func (r *defaultRepo) User() UserRepo {
	return r.user
}

func (r *defaultRepo) VerificationCode() VerificationCodeRepo {
	return r.verificationCode
}
