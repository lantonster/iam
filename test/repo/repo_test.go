package repo

import (
	"testing"

	"github.com/lantonster/iam/config"
	"github.com/lantonster/iam/internal/model"
	"github.com/lantonster/iam/internal/repo"
	"gorm.io/gorm"
)

var (
	r  repo.Repo
	db *gorm.DB
)

func TestMain(m *testing.M) {
	conf := config.NewConfig()
	r = repo.NewDefaultRepo(conf)

	// 清空数据库
	cleardb(conf)

	// 填充测试数据
	testdata()

	m.Run()
}

func cleardb(conf *config.Config) {
	var err error
	if db, err = conf.MySQL.Connect(); err != nil {
		panic(err)
	}

	// 获取所有表名
	tables := []string{}
	db.Raw("SHOW TABLES").Scan(&tables)

	// 循环使用 TRUNCATE TABLE 重置自增 ID 并清除数据
	for _, table := range tables {
		db.Exec("TRUNCATE TABLE " + table)
	}
}

var (
	testUser = model.User{Username: "test_user"}
)

func testdata() {
	db.Create(&testUser)
}
