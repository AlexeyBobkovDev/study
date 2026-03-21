package simple_sql

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	SELECT *
	FROM tasks
	ORDER BY id;
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id          int
			title       string
			description string
			completed   bool
			createdAt   time.Time
			completedAt *time.Time
		)

		err := rows.Scan(
			&id,
			&title,
			&description,
			&completed,
			&createdAt,
			&completedAt,
		)
		if err != nil {
			return err
		}
		fmt.Println(id, title, description, completed, createdAt, completedAt)
	}

	return nil
}
