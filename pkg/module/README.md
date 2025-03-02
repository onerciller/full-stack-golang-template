# Module System

This package provides a module system for structuring applications using [Uber's fx](https://github.com/uber-go/fx) dependency injection framework.

## Module Interfaces

The module system defines several interfaces that modules can implement:

### `Module` Interface

The base interface that all modules must implement:

```go
type Module interface {
    // Providers returns the providers this module offers
    Provide() fx.Option
}
```

### `ModuleWithInvoke` Interface

For modules that need to perform actions during the invoke phase:

```go
type ModuleWithInvoke interface {
    Module
    Invoke() fx.Option
}
```

### `ModuleWithLifecycle` Interface

For modules that need to register lifecycle hooks that run after all modules have been registered:

```go
type ModuleWithLifecycle interface {
    Module
    RegisterLifecycle() fx.Option
}
```

## Registration Order

When registering modules using the `Register` function, the following order is observed:

1. All module providers are registered first (via `Provide()`)
2. Then all module invokes are registered (via `Invoke()`)
3. Finally, all lifecycle hooks are registered (via `RegisterLifecycle()`)

This ensures that lifecycle hooks have access to all dependencies provided by other modules.

## Creating a Module with Lifecycle Hooks

To create a module with lifecycle hooks:

```go
type MyModule struct {
    fx.In
}

// Provide implements the Module interface
func (m *MyModule) Provide() fx.Option {
    return fx.Module("my-module",
        fx.Provide(NewMyService),
    )
}

// RegisterLifecycle implements the ModuleWithLifecycle interface
func (m *MyModule) RegisterLifecycle() fx.Option {
    return fx.Invoke(func(lc fx.Lifecycle, service *MyService) {
        lc.Append(fx.Hook{
            OnStart: func(ctx context.Context) error {
                // Do something when the application starts
                return nil
            },
            OnStop: func(ctx context.Context) error {
                // Do something when the application stops
                return nil
            },
        })
    })
}
```

## Example Usage

See `example_lifecycle_module.go` for a complete example of a module that implements both `ModuleWithInvoke` and `ModuleWithLifecycle` interfaces.

## Default Modules

The `DefaultModules()` function returns a list of default modules that are included in the application. You can use `RegisterWithDefaultModules()` to register your own modules alongside the default ones.

```go
app := module.RegisterWithDefaultModules(
    &myapp.MyModule{},
    &myapp.AnotherModule{},
)
``` 