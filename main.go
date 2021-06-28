package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"
)

const DbUrl = "postgres://iman_user:iman_root@localhost:5432/golangdb"

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

func hello(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "Hello world\n")
	body := json.NewDecoder(req.Body)
	log.Print(body)
}

func main() {
	dbConnection := connectDB()
	defer dbConnection.Close()

	http.HandleFunc("/", hello)

	err := http.ListenAndServe(":8099", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//contact := ContactList{}

	//contact.CreateContactList(dbConnection)

	//contact.GetContact(dbConnection, 4)
	//fmt.Printf("ID\tName\tE-Mail\t\tAddress\t\tPhone Number\tCreated Date\n")
	//fmt.Printf("%d\t%s\t%s\t%s\t%d\t%s\t%s\n",
	//	contact.id, contact.name, contact.email, contact.address, contact.phoneNumber, contact.createdAt.String(), contact.updatedAt.String())

	//contacts := contact.GetContactList(dbConnection)
	//defer contacts.Close()
	//fmt.Printf("ID\tName\tE-Mail\t\tAddress\t\tPhone Number\tCreated Date\n")
	//for contacts.Next(){
	//	err := contacts.Scan(&contact.id, &contact.name, &contact.email, &contact.address, &contact.phoneNumber, &contact.createdAt, &contact.updatedAt)
	//	if err != nil {
	//		return
	//	}
	//fmt.Printf("%d\t%s\t%s\t%s\t%d\t%s\t%s\n",
	//	contact.id, contact.name, contact.email, contact.address, contact.phoneNumber, contact.createdAt.String(), contact.updatedAt.String())
	//}

	//contact.UpdateContact(dbConnection, 4)
	//contact.DeleteContact(dbConnection, 5)
}