package neo

import (
	"context"
	"encoding/json"

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
	ID int `json:"id"`
	Name string `json:"name"`
	Role int `json:"role_id"`
}

// IncomingUser type 
type IncomingUser struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Role int `json:"role_id"`
}

func (i *IncomingUser) marshalJSON() ([]byte, error) {
	i := User{
		ID: u.ID,
		Name: u.Name,
		Role: u.Role,
	}

	return json.Marshal(o)
}