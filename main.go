package main

import (
	"context"
	"github.com/raufhm/learning-uberfx/uberfx"
	"go.uber.org/fx"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := fx.New(provider(), invoker())

	if err := app.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}

	gracefulShutdown(app)

	if err := app.Stop(context.Background()); err != nil {
		log.Printf("Failed to stop application gracefully: %v", err)
	}
}

func provider() fx.Option {
	return fx.Options(
		uberfx.ProvideViper(),
		uberfx.ProvideConfig(),
		uberfx.ProvideDBConnection(),
		uberfx.ProvideGinEngine(),
		uberfx.ProvideUserHandler(),
		uberfx.ProvideUserRepository(),
		uberfx.ProvideUserService(),
	)
}

func invoker() fx.Option {
	return fx.Options(uberfx.InvokeHTTPServer())
}

func gracefulShutdown(app *fx.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down gracefully...")
	app.Stop(context.Background())
	log.Println("Application stopped")
}
