package main

import (
	"github.com/lzaun/iam/config"
	"github.com/lzaun/iam/internal/model"
	"github.com/lzaun/iam/internal/repo"
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
