package uberfx

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/raufhm/learning-uberfx/config"
	"go.uber.org/fx"
	"log"
	"net/http"
	"time"
)

func InvokeHTTPServer() fx.Option {
	return fx.Invoke(StartHTTPServer)
}

func StartHTTPServer(lifecycle fx.Lifecycle, engine *gin.Engine, cfg *config.Config) {
	server := &http.Server{
		Addr:    cfg.Server.Port,
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

			return server.Shutdown(ctx)
		},
	})
}
