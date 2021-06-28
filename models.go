package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"time"
)

type ContactList struct {
	name, email, address string
	phoneNumber          int
	createdAt            time.Time
}

func (c *ContactList) CreateContactList(conn *pgxpool.Pool) {
	sqlQuery := fmt.Sprintf("insert into contactlist (name, email, address, phonenumber, createdat) VALUES('%s', '%s', '%s', '%d', '%s')", c.name, c.email, c.address, c.phoneNumber, c.createdAt)
	rows, err := conn.Query(context.Background(), sqlQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}
	rows.Close()
}

type TaskList struct {
	contact   ContactList
	task      string
	completed bool
	createdAt time.Time
}
