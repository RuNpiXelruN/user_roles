package pg

import (
	"testing"

	"github.com/ryboe/q"
)

func TestPostgres(t *testing.T) {
	conn, err := connTest()
	if err != nil {
		q.Q(err)
	} else {
		q.Q(conn)
	}
}
