//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/lantonster/iam/config"
	"github.com/lantonster/iam/internal/handler"
	"github.com/lantonster/iam/internal/repo"
	"github.com/lantonster/iam/internal/router"
	"github.com/lantonster/iam/internal/server"
	"github.com/lantonster/iam/internal/service"
)

var serverSet = wire.NewSet(server.NewServer)

var routerSet = wire.NewSet(router.NewRouter)

var handlerSet = wire.NewSet(handler.NewHandler)

var serviceSet = wire.NewSet(service.NewDefaultService)

var repoSet = wire.NewSet(repo.NewDefaultRepo)

var configSet = wire.NewSet(config.NewConfig)

func Init() *server.Server {
	wire.Build(
		serverSet,
		routerSet,
		handlerSet,
		serviceSet,
		repoSet,
		configSet,
	)

	return new(server.Server)
}
