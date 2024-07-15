//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/lzaun/iam/config"
	"github.com/lzaun/iam/internal/router"
	"github.com/lzaun/iam/internal/server"
)

var serverSet = wire.NewSet(server.NewServer)

var routerSet = wire.NewSet(router.NewRouter)

var configSet = wire.NewSet(config.NewConfig)

func Init() *server.Server {
	wire.Build(
		serverSet,
		routerSet,
		configSet,
	)

	return new(server.Server)
}
