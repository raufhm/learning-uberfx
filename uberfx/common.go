package uberfx

import "go.uber.org/fx"

func invoker() fx.Option {
	return fx.Options(
		InvokeRegisterRoutes(),
		InvokeHTTPServer(),
	)
}

func LoadOptions() []fx.Option {
	return []fx.Option{
		provideDependencies(),
		provideHandler(),
		provideService(),
		provideRepository(),
		invoker(),
	}
}
