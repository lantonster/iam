//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/lantonster/iam/config"
	"github.com/lantonster/iam/internal/handler"
	"github.com/lantonster/iam/internal/router"
	"github.com/lantonster/iam/internal/server"
)

var serverSet = wire.NewSet(server.NewServer)

var routerSet = wire.NewSet(router.NewRouter)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewAuthHandler,
)

var configSet = wire.NewSet(config.NewConfig)

func Init() *server.Server {
	wire.Build(
		serverSet,
		routerSet,
		handlerSet,
		configSet,
	)

	return new(server.Server)
}
