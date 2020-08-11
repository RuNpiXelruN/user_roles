package db

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

// PGService type 
type PGService interface{}

// PGClient type 
type PGClient struct {}

// NewPGClient func 
func NewPGClient() *PGClient {
	return &PGClient{}
}

// Connect func 
func Connect(cfg Config) {
	pgCfg, err := pgxpool.ParseConfig(fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s", cfg.Host, cfg.Port, cfg.DBName, cfg.User, cfg.Password))
	if err != nil {
		return nil, err
	}
}