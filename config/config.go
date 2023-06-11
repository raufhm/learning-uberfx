package config

import "fmt"

type Config struct {
	Server *Server
	DB     *DB
}

type Server struct {
	Port string
}

type DB struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg *Config) DBStringConn() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.User, cfg.DB.Password, cfg.DB.SSLMode)
}
