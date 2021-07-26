package main

import (
	"bufio"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"time"
)

const DataBaseURL = "postgres://iman_user:iman_root@localhost:5432/golangdb"

func main() {
	dbConnection := connectDB()
	defer dbConnection.Close()

	var choice int

	for true {
		fmt.Print("\n1.Create Contact.\n2.Contact List.\n3.Get Contact by id.\n4.Update Contact (put) by id.\n5.Delete Contact.\n")
		fmt.Print("6.Create Task.\n7.Task List.\n8.Get Task by id.\n9.Update Task (put) by id.\n0.Delete Task.\n")
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Print(err)
			continue
		}
		fmt.Println(choice)
		switch choice {
		case 1:
			contact := ContactList{}
			fmt.Print("Enter Name: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan(){
				contact.name = scanner.Text()
			}
			fmt.Print("Enter E-mail: ")
			_, _ = fmt.Scanf("%s", &contact.email)
			fmt.Print("Enter Phone number: ")
			_, _ = fmt.Scanf("%d", &contact.phoneNumber)
			fmt.Print("Enter Address: ")
			scanner = bufio.NewScanner(os.Stdin)
			if scanner.Scan(){
				contact.address = scanner.Text()
			}
			contact.createdAt = time.Now()
			contact.updatedAt = time.Now()

			CreateContact(&contact, dbConnection)
		case 2:
			contact := ContactList{}
			GetContactList(&contact, dbConnection)
		case 3:
			contact := ContactList{}
			GetContact(&contact, dbConnection)
		case 4:
			contact := ContactList{}
			var contactId int
			fmt.Print("Enter ID of contact: ")
			_, err := fmt.Scanf("%d", &contactId)
			if err != nil {
				fmt.Print(err)
				return
			}
			fmt.Print("Enter Name: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan(){
				contact.name = scanner.Text()
			}
			fmt.Print("Enter E-mail: ")
			_, _ = fmt.Scanf("%s", &contact.email)
			fmt.Print("Enter Phone number: ")
			_, _ = fmt.Scanf("%d", &contact.phoneNumber)
			fmt.Print("Enter Address: ")
			scanner = bufio.NewScanner(os.Stdin)
			if scanner.Scan(){
				contact.address = scanner.Text()
			}
			contact.updatedAt = time.Now()

			UpdateContact(&contact, dbConnection, contactId)
		case 5:
			contact := ContactList{}
			var contactId int
			fmt.Print("Enter ID of contact: ")
			_, err := fmt.Scanf("%d", &contactId)
			if err != nil {
				fmt.Print(err)
				return
			}

			DeleteContact(&contact, dbConnection, contactId)
		case 6:
			task := TaskList{}
			fmt.Print("Enter Task: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan(){
				task.task = scanner.Text()
			}
			fmt.Print("Enter Contact ID: ")
			_, _ = fmt.Scanf("%d", &task.contactId)
			task.completed = false
			task.createdAt = time.Now()
			task.updatedAt = time.Now()

			CreateTask(&task, dbConnection)
		case 7:
			task := TaskList{}
			GetTaskList(&task, dbConnection)
		case 8:
			task := TaskList{}
			GetTask(&task, dbConnection)
		case 9:
			task := TaskList{}
			var taskID int
			fmt.Print("Enter ID of task: ")
			_, err := fmt.Scanf("%d", &taskID)
			if err != nil {
				fmt.Print(err)
				return
			}
			fmt.Print("Enter Task: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan(){
				task.task = scanner.Text()
			}
			fmt.Print("Enter Contact ID: ")
			_, _ = fmt.Scanf("%s", &task.contactId)
			task.updatedAt = time.Now()

			UpdateTask(&task, dbConnection, taskID)
		case 0:
			task := TaskList{}
			var taskID int
			fmt.Print("Enter ID of task: ")
			_, err := fmt.Scanf("%d", &taskID)
			if err != nil {
				fmt.Print(err)
				return
			}
			DeleteTask(&task, dbConnection, taskID)
		default:
			os.Exit(2)
		}
	}
}

func connectDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", DataBaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}else{
		fmt.Println("Connected to database")
	}
	return db
}