package uberfx

import "go.uber.org/fx"

func LoadOptions() []fx.Option {
	return []fx.Option{
		provideDependencies(),
		provideHandler(),
		provideService(),
		provideRepository(),
		InvokeRegisterRoutes(),
		InvokeHTTPServer(),
	}
}
