package main

import (
	"github.com/lantonster/iam/config"
	"github.com/lantonster/iam/internal/model"
	"github.com/lantonster/iam/internal/repo"
	"gorm.io/gen"
)

func main() {
	conf := config.NewConfig()
	db := repo.NewGormDB(conf)

	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/repo",
		Mode:    gen.WithDefaultQuery, // generate mode
	})

	g.UseDB(db)

	g.ApplyBasic(
		model.User{},
	)

	g.Execute()
}
