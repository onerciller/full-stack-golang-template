package config

import "github.com/samber/do"

func Provide(di *do.Injector) (ConfigProvider, error) {
	return New(
		WithPath("./configs"),
		WithConfigName("config"),
		WithConfigType("yaml"),
	), nil
}
