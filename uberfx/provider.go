package uberfx

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/raufhm/learning-uberfx/config"
	"github.com/raufhm/learning-uberfx/handler"
	"github.com/raufhm/learning-uberfx/repository"
	"github.com/raufhm/learning-uberfx/service"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func NewGinEngine() *gin.Engine {
	return gin.Default()
}

func NewUserHandler(userService *service.UserService) *handler.UserHandler {
	return &handler.UserHandler{
		UserService: userService,
	}
}

func NewUserRepository() repository.UserRepository {
	return repository.NewUserRepository()
}

func NewUserService(userRepo repository.UserRepository) *service.UserService {
	return &service.UserService{
		UserRepository: userRepo,
	}
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

func ProvideGinEngine() fx.Option {
	return fx.Provide(NewGinEngine)
}

func ProvideUserHandler() fx.Option {
	return fx.Provide(NewUserHandler)
}

func ProvideUserRepository() fx.Option {
	return fx.Provide(NewUserRepository)
}

func ProvideUserService() fx.Option {
	return fx.Provide(NewUserService)
}

func ProvideViper() fx.Option {
	return fx.Provide(NewViperEngine)
}

func ProvideConfig() fx.Option {
	return fx.Provide(NewProvideConfig)
}

func ProvideDBConnection() fx.Option {
	return fx.Provide(NewDBConnection)
}
