package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"time"
)

type ContactList struct {
	id					 int
	name, email, address string
	phoneNumber          int
	createdAt            time.Time
	updatedAt			 time.Time
}

type ContactListInterface interface {
	CreateContactList(conn *pgxpool.Pool)
	GetContact(conn *pgxpool.Pool, contactId int)
	GetContactList(conn *pgxpool.Pool) pgx.Rows
	UpdateContact(conn *pgxpool.Pool, contactId int)
	DeleteContact(conn *pgxpool.Pool, contactId int)
}

func (c *ContactList) CreateContactList(conn *pgxpool.Pool) {
	_ , err := conn.Exec(context.Background(), "insert into contactlist (name, email, address, phonenumber, createdat, updated_at) VALUES($1, $2, $3, $4, $5, $6)",
		c.name, c.email, c.address, c.phoneNumber, c.createdAt, c.updatedAt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query create failed: %v\n", err)
	}
}

func (c *ContactList) GetContact(conn *pgxpool.Pool, contactId int) {
	err := conn.QueryRow(context.Background(), "select * from contactlist where id=$1", contactId).
		Scan(&c.id, &c.name, &c.email, &c.address, &c.phoneNumber, &c.createdAt, &c.updatedAt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query get failed: %v\n", err)
	}
}

func (c *ContactList) GetContactList(conn *pgxpool.Pool) pgx.Rows {
	rows, err := conn.Query(context.Background(), "select * from contactlist order by id")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query get list failed: %v\n", err)
	}
	return rows
}

func (c *ContactList) UpdateContact(conn *pgxpool.Pool, contactId int) {
	_, err := conn.Exec(context.Background(),
		"update contactlist set name=$2, email=$3, address=$4, phonenumber=$5, updated_at=$6 where id=$1",
		contactId, c.name, c.email, c.address, c.phoneNumber, c.updatedAt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query update failed: %v\n", err)
	}
}

func (c *ContactList) DeleteContact(conn *pgxpool.Pool, contactId int) {
	_, err := conn.Exec(context.Background(),
		"delete from contactlist where id=$1", contactId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query update failed: %v\n", err)
	}
}

type TaskList struct {
	contact   ContactList
	task      string
	completed bool
	createdAt time.Time
	updatedAt time.Time
}
