package shared

import "go.uber.org/fx"

type Module struct {
	fx.In
}

func (m *Module) Provide() fx.Option {
	return fx.Module("shared",
		fx.Provide(NewBaseHandler),
	)
}
