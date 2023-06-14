package uberfx

import (
	"context"
	"fmt"
	"github.com/raufhm/learning-uberfx/internal/route"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func StartHTTPServer(lifecycle fx.Lifecycle, engine *gin.Engine, v *viper.Viper) {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", v.GetString("SERVER_PORT")),
		Handler: engine,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("Failed to start HTTP server: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			fmt.Println("Shutting down the server...")

			if err := server.Shutdown(ctx); err != nil {
				return fmt.Errorf("failed to gracefully shut down the server: %v", err)
			}

			fmt.Println("Server gracefully stopped")

			return nil
		},
	})
}

func InvokeHTTPServer() fx.Option {
	return fx.Invoke(StartHTTPServer)
}

func InvokeRegisterRoutes() fx.Option {
	return fx.Invoke(route.RegisterRoutes)
}
