package main

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"time"
)

type ContactList struct {
	name, email, address string
	phoneNumber          int
	createdAt            time.Time
	updatedAt			 time.Time
}

type ContactListInterface interface {
	CreateContactList(conn *sqlx.DB) error
	GetContact(conn *sqlx.DB, contactId int) (*sql.Rows, error)
	GetContactList(conn *sqlx.DB) (*sql.Rows, error)
	UpdateContact(conn *sqlx.DB, contactId int) (*sql.Rows, error)
	DeleteContact(conn *sqlx.DB, contactId int) error
	Values() ContactList
}

func (c ContactList) Values() ContactList{
	return c
}

func (c *ContactList) CreateContactList(conn *sqlx.DB) error {
	_ , err := conn.Exec("insert into contactlist (name, email, address, phonenumber, createdat, updated_at) VALUES($1, $2, $3, $4, $5, $6)",
		c.name, c.email, c.address, c.phoneNumber, c.createdAt, c.updatedAt)
	return err
}

func (c *ContactList) GetContact(conn *sqlx.DB, contactId int) (*sql.Rows, error){
	rows, err := conn.Query("select * from contactlist where id=$1", contactId)
	return rows, err
}

func (c *ContactList) GetContactList(conn *sqlx.DB) (*sql.Rows, error) {
	rows, err := conn.Query("select * from contactlist order by id")
	return rows, err
}

func (c *ContactList) UpdateContact(conn *sqlx.DB, contactId int) (*sql.Rows, error) {
	_, err := conn.Exec("update contactlist set name=$2, email=$3, address=$4, phonenumber=$5, updated_at=$6 where id=$1",
		contactId, c.name, c.email, c.address, c.phoneNumber, c.updatedAt)

	if err != nil {
		return nil, err
	}

	rows, err1 := conn.Query("select * from contactlist where id=$1", contactId)
	if err1 != nil {
		return nil, err1
	}
	return rows, nil
}

func (c *ContactList) DeleteContact(conn *sqlx.DB, contactId int) error {
	_, err := conn.Exec(		"delete from contactlist where id=$1", contactId)
	return err

}