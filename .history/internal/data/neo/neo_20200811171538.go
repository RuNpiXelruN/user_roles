package neo

import (
	"context"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// Service type 
type Service interface{
	Conn() bolt.Conn
	Seed(ctx context.Context) error
	GetSubordinates(userID string) ([]User, error)
}

// Client type 
type Client struct {
	conn bolt.Conn
}

// NewClient func 
func NewClient() *Client {
	return &Client{
		conn: Connect(),
	}
}

// Conn func 
func (c *Client) Conn() bolt.Conn {
	return c.conn
}

// User type 
type User struct {
	ID int `json:"Id"`
	Name string `json:"Name"`
	RoleID int `json:"Role"` 
}