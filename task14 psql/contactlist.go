package main

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
)

type ContactList struct {
	name, email, address string
	phoneNumber          int
	createdAt            time.Time
	updatedAt			 time.Time
}

type ContactListInterface interface {
	CreateContactList(conn *sqlx.DB)
	GetContact(conn *sqlx.DB, contactId int) *sql.Rows
	GetContactList(conn *sqlx.DB) *sql.Rows
	UpdateContact(conn *sqlx.DB, contactId int) *sql.Rows
	DeleteContact(conn *sqlx.DB, contactId int)
	Values() ContactList
}

func (c ContactList) Values() ContactList{
	return c
}

func (c *ContactList) CreateContactList(conn *sqlx.DB) {
	_ , err := conn.Exec("insert into contactlist (name, email, address, phonenumber, createdat, updated_at) VALUES($1, $2, $3, $4, $5, $6)",
		c.name, c.email, c.address, c.phoneNumber, c.createdAt, c.updatedAt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query create failed: %v\n", err)
	} else {
		fmt.Println("Record created")
	}
}

func (c *ContactList) GetContact(conn *sqlx.DB, contactId int) *sql.Rows{
	rows, err := conn.Query("select * from contactlist where id=$1", contactId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query get failed: %v\n", err)
	}
	return rows
}

func (c *ContactList) GetContactList(conn *sqlx.DB) *sql.Rows {
	rows, err := conn.Query("select * from contactlist order by id")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query get list failed: %v\n", err)
	}
	return rows
}

func (c *ContactList) UpdateContact(conn *sqlx.DB, contactId int) *sql.Rows {
	_, err := conn.Exec(		"update contactlist set name=$2, email=$3, address=$4, phonenumber=$5, updated_at=$6 where id=$1",
		contactId, c.name, c.email, c.address, c.phoneNumber, c.updatedAt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query update failed: %v\n", err)
	} else {
		fmt.Println("Record updated")
	}
	rows, err := conn.Query("select * from contactlist where id=$1", contactId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query get failed: %v\n", err)
	}
	return rows
}

func (c *ContactList) DeleteContact(conn *sqlx.DB, contactId int) {
	_, err := conn.Exec(		"delete from contactlist where id=$1", contactId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query update failed: %v\n", err)
	}else {
		fmt.Println("Record deleted")
	}
}