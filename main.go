package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

func main() {
	dbpool, err := pgxpool.Connect(context.Background(), "postgres://iman_user:iman_root@localhost:5432/golangdb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	_, err = dbpool.Query(context.Background(), "insert into contactlist (name, email, address, phonenumber, createdat) values('Ivan', 'ivan@dot.com', 'Back Street', 942345422, now())")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
}
