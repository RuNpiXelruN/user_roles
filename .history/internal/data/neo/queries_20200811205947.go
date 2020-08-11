package neo

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
)

func (c *Client) prepareStatement(query string, conn bolt.Conn) (bolt.Stmt, error) {

	st, err := conn.PrepareNeo(query)
	if err != nil {
		log.Println("conn.PrepareNeo error", err)
		return nil, err
	}

	return st, nil
}

// Seed func
func (c *Client) Seed(ctx context.Context) error {

	st, err := c.prepareStatement(seedDB, c.conn)
	if err != nil {
		return err
	}

	_, err = st.ExecNeo(map[string]interface{}{})
	if err != nil {
		log.Println("st.ExecNeo error", err)
		return err
	}

	return nil
}

// GetSubordinates func
func (c *Client) GetSubordinates(ctx context.Context, userID string) ([]byte, error) {
// func (c *Client) GetSubordinates(ctx context.Context, userID string) ([]User, error) {

	idInt, err := strconv.Atoi(userID)
	if err != nil {
		log.Println("strconv.Atoi error", err)
		return nil, err
	}

	st, err := c.prepareStatement(getSubordinates, c.conn)
	if err != nil {
		return nil, err
	}

	rows, err := st.QueryNeo(map[string]interface{}{
		"userID": idInt,
	})

	if err != nil {
		log.Println("st.QueryNeo error", err)
		return nil, err
	}

	defer rows.Close()
	data, _, err := rows.All()

	responseBytes, err := c.MapResponseToUsers(data)
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}

// MapResponseToUsers func
func (c *Client) MapResponseToUsers(data [][]interface{}) ([]byte, error) {
// func (c *Client) MapResponseToUsers(data [][]interface{}) ([]User, error) {

	users := []User{}

	for _, vals := range data {
		for _, val := range vals {
			user, err := c.mapUserNode(val.(graph.Node).Properties)
			if err != nil {
				return nil, err
			}

			users = append(users, *user)
		}
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		log.Println("json.Marshal error", err)
		return nil, err
	}

	return bytes, nil
}
