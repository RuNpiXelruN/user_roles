package neo

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// Service type 
type Service interface{
	Conn() bolt.Conn
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