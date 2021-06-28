package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"time"
)

func main() {
	conn, err := pgxpool.Connect(context.Background(), "postgres://iman_user:iman_root@localhost:5432/golangdb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	contact := ContactList{"Sanjar", "dot@mail.ty",
						"Main Street", 989341233, time.Now()}

	contact.CreateContactList(conn)
}
