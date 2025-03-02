package module

import (
	"context"
	"log"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

// AppLifecycleModule provides a way to register lifecycle hooks that run
// after all other modules have been registered
type AppLifecycleModule struct {
	AutoMigrate func(db *gorm.DB) error
}

// Provide satisfies the Module interface
func (m *AppLifecycleModule) Provide() fx.Option {
	// No need to provide anything as AppLifecycleHooks is automatically constructed by fx
	return fx.Options()
}

// RegisterLifecycle satisfies the ModuleWithLifecycle interface
// This will be called after all other modules are registered
func (m *AppLifecycleModule) RegisterLifecycle() fx.Option {
	return fx.Invoke(func(lc fx.Lifecycle, hooks AppLifecycleHooks, db *gorm.DB) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Println("AppLifecycle: All modules have been registered and started")

				if m.AutoMigrate != nil {
					if err := m.AutoMigrate(db); err != nil {
						return err
					}
				}

				return hooks.OnApplicationStart(ctx)
			},
			OnStop: func(ctx context.Context) error {
				log.Println("AppLifecycle: Application is shutting down")
				return hooks.OnApplicationStop(ctx)
			},
		})
	})
}
