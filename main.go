package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

const DbUrl = "postgres://iman_user:iman_root@localhost:5432/golangdb"

func main() {
	dbConnection := connectDB()
	defer dbConnection.Close()

	var choice int

	for true {
		fmt.Println("1.Create Contact.\n2.Contact List.\n3.Get Contact by id.\n4.Update Contact (put) by id.\n5.Delete Contact.\n")
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Print(err)
			continue
		}
		fmt.Println(choice)
		switch choice {
		case 1:
			contact := ContactList{}

			CreateContact(&contact, dbConnection)
		case 2:
			contact := ContactList{}
			GetContactList(&contact, dbConnection)
		case 3:
			contact := ContactList{}
			GetContact(&contact, dbConnection)
		case 4:
			contact := ContactList{}
			UpdateContact(&contact, dbConnection)
		}
	}

	//contact.UpdateContact(dbConnection, 4)
	//contact.DeleteContact(dbConnection, 5)
}

func CreateContact(c ContactListInterface, conn *pgxpool.Pool) {
	c.CreateContactList(conn)
}

// GetContact - Get contact from db by ID pk
func GetContact(c ContactListInterface, conn *pgxpool.Pool) {
	values := c.Values()
	var contactId int
	_, err := fmt.Scanf("%d", &contactId)
	if err != nil {
		fmt.Print(err)
		return
	}
	contact := c.GetContact(conn, contactId)
	fmt.Printf("ID\tName\tE-Mail\t\tAddress\t\tPhone Number\tCreated Date\n")
	for contact.Next(){
		err := contact.Scan(&values.id, &values.name, &values.email, &values.address, &values.phoneNumber, &values.createdAt, &values.updatedAt)
		if err != nil {
			return
		}
		fmt.Printf("%d\t%s\t%s\t%s\t%d\t%s\t%s\n",
			values.id, values.name, values.email, values.address, values.phoneNumber, values.createdAt.String(), values.updatedAt.String())
	}
}

func GetContactList(contact ContactListInterface, conn *pgxpool.Pool) {
	values := contact.Values()
	contacts := contact.GetContactList(conn)
	defer contacts.Close()
	fmt.Printf("ID\tName\tE-Mail\t\tAddress\t\tPhone Number\tCreated Date\n")
	for contacts.Next(){
		err := contacts.Scan(&values.id, &values.name, &values.email, &values.address, &values.phoneNumber, &values.createdAt, &values.updatedAt)
		if err != nil {
			return
		}
		fmt.Printf("%d\t%s\t%s\t%s\t%d\t%s\t%s\n",
			values.id, values.name, values.email, values.address, values.phoneNumber, values.createdAt.String(), values.updatedAt.String())
	}
}

func UpdateContact(c ContactListInterface, conn *pgxpool.Pool) {
	//values := c.Values()
	var contactId int
	_, err := fmt.Scanf("%d", &contactId)
	if err != nil {
		fmt.Print(err)
		return
	}
	c.UpdateContact(conn, contactId)
}

func connectDB() *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), DbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}else{
		fmt.Println("Connected to database")
	}
	return pool
}