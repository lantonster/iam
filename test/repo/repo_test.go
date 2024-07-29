package repo

import (
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/lantonster/iam/config"
	"github.com/lantonster/iam/internal/model"
	"github.com/lantonster/iam/internal/repo"
	"gorm.io/gorm"
)

var (
	r   repo.Repo
	db  *gorm.DB
	rdb *redis.Client
)

func TestMain(m *testing.M) {
	conf := config.NewConfig()
	r = repo.NewDefaultRepo(conf)

	var err error
	if db, err = conf.MySQL.Connect(); err != nil {
		panic(err)
	}
	if rdb, err = conf.Redis.Connect(); err != nil {
		panic(err)
	}

	// 清空数据库
	cleardb()
	clearrdb()

	// 填充测试数据
	testdata()

	m.Run()
}

func cleardb() {
	// 获取所有表名
	tables := []string{}
	db.Raw("SHOW TABLES").Scan(&tables)

	// 循环使用 TRUNCATE TABLE 重置自增 ID 并清除数据
	for _, table := range tables {
		db.Exec("TRUNCATE TABLE " + table)
	}
}

func clearrdb() {
	rdb.FlushDB(rdb.Context())
}

var (
	testUser = model.User{Username: "test_user"}
)

func testdata() {
	db.Create(&testUser)
}
