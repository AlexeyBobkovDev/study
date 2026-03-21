package main

import (
	"context"
	"fmt"
	"study/feature_postgres/simple_connection"
	"study/simple_sql"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	// if err := simple_sql.InsertRow(
	// 	ctx,
	// 	conn,
	// 	"Dinner",
	// 	"Need to eat",
	// 	true,
	// 	time.Now(),
	// ); err != nil {
	// 	panic(err)
	// }

	if err := simple_sql.UpdateRow(ctx, conn); err != nil {
		panic(err)
	}

	if err := simple_sql.DeleteRow(ctx, conn); err != nil {
		panic(err)
	}

	if err := simple_sql.SelectRows(ctx, conn); err != nil{
		panic(err)
	}

	fmt.Println("Succeed!")
}
