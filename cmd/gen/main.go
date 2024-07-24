package main

import (
	"github.com/lantonster/corekit"
	"github.com/lantonster/iam/config"
	"github.com/lantonster/iam/internal/model"
	"gorm.io/gen"
)

func main() {
	conf := config.NewConfig()
	db := corekit.ConnectMySQL(conf.MySQL)

	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/dao",
		Mode:    gen.WithDefaultQuery, // generate mode
	})

	g.UseDB(db)

	g.ApplyBasic(
		model.User{},
	)

	g.Execute()
}
