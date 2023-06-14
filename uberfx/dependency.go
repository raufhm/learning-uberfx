package uberfx

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func NewGinEngine() *gin.Engine {
	return gin.Default()
}

func loadDbConn(v *viper.Viper) string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		v.GetString("DB_HOST"),
		v.GetString("DB_PORT"),
		v.GetString("DB_NAME"),
		v.GetString("DB_USER"),
		v.GetString("DB_PASSWORD"),
		v.GetString("DB_SSLMODE"),
	)
}

func NewDBConnection(v *viper.Viper) (*sql.DB, error) {
	db, err := sql.Open("postgres", loadDbConn(v))
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

//func NewProvideConfig(v *viper.Viper) (*config.Config, error) {
//	cfg := &config.Config{}
//
//	err := v.Unmarshal(cfg)
//	if err != nil {
//		return nil, err
//	}
//
//	return cfg, nil
//}

func provideDependencies() fx.Option {
	return fx.Provide(
		NewDBConnection,
		//NewProvideConfig,
		NewViperEngine,
		NewGinEngine,
	)
}
