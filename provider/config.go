package provider

import (
	"github.com/onerciller/fullstack-golang-template/pkg/config"
)

// @Configuration
type AppConfig struct{}

// @Bean
func (c *AppConfig) Provide() config.ConfigProvider {
	return config.New(
		config.WithPath("./configs"),
		config.WithConfigName("config"),
		config.WithConfigType("yaml"),
	)
}
