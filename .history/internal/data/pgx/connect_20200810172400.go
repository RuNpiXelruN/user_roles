package pgx

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// NewConn func 
func NewConn() (*pgx.Conn, error) {
	conCfg, err := pgx.ParseConfig("postgres://postgres:password@127.0.0.1:5432/deputychallenge?search_path=deputy")
	if err != nil {
		fmt.Println("pgx.ParseConfig error", err)
		return nil, err
	}

	conn, err := pgx.ConnectConfig(context.Background(), conCfg)
	if err != nil {
		fmt.Println("pgx.Connect error", err)
		return nil, err
	}

	return conn, nil
}