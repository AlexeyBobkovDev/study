package simple_connection

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	env := os.Getenv("conn_string")
	if env == "" {
		fmt.Println("Connection failed")
		return nil, errors.New("Conn string does not exist")
	}
	return pgx.Connect(ctx, env)
}
