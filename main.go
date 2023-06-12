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
	app := fx.New(
		uberfx.LoadOptions()...,
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}

	gracefulShutdown(app)

	if err := app.Stop(context.Background()); err != nil {
		log.Printf("Failed to stop application gracefully: %v", err)
	}
}

func gracefulShutdown(app *fx.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down gracefully...")
	app.Stop(context.Background())
	log.Println("Application stopped")
}
