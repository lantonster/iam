package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

func NewConfig() *Config {
	v := viper.New()

	config := Config{}
	initDefaults(v, config)

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./configs")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
}

func initDefaults(v *viper.Viper, config Config) {
	initDefaultsRecursive(v, reflect.TypeOf(config), "")
}

func initDefaultsRecursive(v *viper.Viper, t reflect.Type, prefix string) {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		ft := f.Type

		if f.Anonymous {
			initDefaultsRecursive(v, ft, prefix)
		} else {
			name := strings.ToLower(f.Name)
			if value, ok := f.Tag.Lookup("mapstructure"); ok {
				name = value
			}
			if ft.Kind() == reflect.Struct {
				initDefaultsRecursive(v, ft, fmt.Sprintf("%s%s.", prefix, name))
			} else {
				if value, ok := f.Tag.Lookup("default"); ok {
					key := prefix + name
					v.SetDefault(key, value)
				}
			}
		}
	}
}
