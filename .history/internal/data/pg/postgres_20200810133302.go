package pg

import (
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Config type 
type Config struct {
	User                  string
	Password              string
	Host                  string
	Port                  int
	DBName                  string
}

// Service type 
type Service interface{}

// Client type 
type Client struct {}

// NewClient func 
func NewClient() *Client {
	return &Client{}
}

// Connect func 
func Connect(cfg Config) ((*pgxpool.Config, error)) {
	pgCfg, err := pgxpool.ParseConfig(fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s", cfg.Host, cfg.Port, cfg.DBName, cfg.User, cfg.Password))
	if err != nil {
		return nil, err
	}
	return pgCfg, nil
}