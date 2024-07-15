//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/lzaun/iam/config"
	"github.com/lzaun/iam/internal/server"
)

var serverSet = wire.NewSet(server.NewServer)

var configSet = wire.NewSet(config.NewConfig)

func Init() *server.Server {
	wire.Build(
		serverSet,
		configSet,
	)

	return new(server.Server)
}
