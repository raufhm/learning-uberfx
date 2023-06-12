package uberfx

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/raufhm/learning-uberfx/config"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func NewGinEngine() *gin.Engine {
	return gin.Default()
}

func NewDBConnection(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DBStringConn())
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func NewViperEngine() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(".env")
	v.SetConfigType("env")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return v, nil
}

func NewProvideConfig(v *viper.Viper) (*config.Config, error) {
	cfg := &config.Config{}

	err := v.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func provideDependencies() fx.Option {
	return fx.Provide(
		NewDBConnection,
		NewProvideConfig,
		NewViperEngine,
		NewGinEngine,
	)
}
