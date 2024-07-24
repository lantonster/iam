package config

import "github.com/lantonster/corekit"

type Config struct {
	Port int

	MySQL *corekit.MySQlConfig
}
