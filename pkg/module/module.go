package module

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/fx"
)

// Module defines the interface that all modules must implement
type Module interface {
	// Providers returns the providers this module offers
	Provide() fx.Option
}

// ModuleWithInvoke interface for modules that need to register invokes
type ModuleWithInvoke interface {
	Module
	// Invoke returns the invokes this module offers
	Invoke() fx.Option
}

// ModuleWithLifecycle interface for modules that need to register lifecycle hooks
// These hooks will be executed after all modules have been registered
type ModuleWithLifecycle interface {
	Module
	// RegisterLifecycle returns the lifecycle hooks this module offers
	RegisterLifecycle() fx.Option
}

type Manager struct {
	app *fx.App
}

func (m *Manager) Start(ctx context.Context) error {
	return m.app.Start(ctx)
}

func (m *Manager) GracefulStop(ctx context.Context) error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("Received termination signal, stopping application...")
	return m.app.Stop(ctx)
}

// Register registers a module with the application
func Register(modules ...Module) *Manager {
	var options []fx.Option

	// Add module providers
	for _, module := range modules {
		options = append(options, module.Provide())
	}

	// Add module invokes
	for _, module := range modules {
		if invokeModule, ok := module.(ModuleWithInvoke); ok {
			options = append(options, invokeModule.Invoke())
		}
	}

	// Add lifecycle hooks (registered last, after all modules are loaded)
	for _, module := range modules {
		if lifecycleModule, ok := module.(ModuleWithLifecycle); ok {
			options = append(options, lifecycleModule.RegisterLifecycle())
		}
	}

	return &Manager{app: fx.New(options...)}

}

func DefaultModules() []Module {
	return []Module{
		&ConfigModule{},
		&HttpServerModule{},
		&DatabaseModule{},
		&SecurityModule{},
	}
}

func RegisterWithDefault(modules ...Module) *Manager {
	return Register(append(DefaultModules(), modules...)...)
}
