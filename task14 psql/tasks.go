package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func CreateContact(c ContactListInterface, conn *sqlx.DB) {
	c.CreateContactList(conn)
}

// GetContact - Get contact from db by ID pk
func GetContact(c ContactListInterface, conn *sqlx.DB) {
	values := c.Values()
	var contactId int
	_, err := fmt.Scanf("%d", &contactId)
	if err != nil {
		fmt.Print(err)
		return
	}
	contact := c.GetContact(conn, contactId)
	fmt.Printf("ID\tName\tE-Mail\t\tAddress\t\tPhone Number\tCreated Date\t\t\tUpdated Date\n")
	for contact.Next(){
		err := contact.Scan(&contactId, &values.name, &values.email, &values.address, &values.phoneNumber, &values.createdAt, &values.updatedAt)
		if err != nil {
			return
		}
		fmt.Printf("%d\t%s\t%s\t%s\t%d\t%s\t%s\n",
			contactId, values.name, values.email, values.address, values.phoneNumber, values.createdAt.String(), values.updatedAt.String())
	}
}

func GetContactList(contact ContactListInterface, conn *sqlx.DB) {
	var contactId int
	values := contact.Values()
	contacts := contact.GetContactList(conn)
	fmt.Printf("ID\tName\tE-Mail\t\tAddress\t\tPhone Number\tCreated Date\t\t\tUpdated Date\n")
	for contacts.Next(){
		err := contacts.Scan(&contactId, &values.name, &values.email, &values.address, &values.phoneNumber, &values.createdAt, &values.updatedAt)
		if err != nil {
			return
		}
		fmt.Printf("%d\t%s\t%s\t%s\t%d\t%s\t%s\n",
			contactId, values.name, values.email, values.address, values.phoneNumber, values.createdAt.String(), values.updatedAt.String())
	}
}

func UpdateContact(c ContactListInterface, conn *sqlx.DB, contactId int) {
	values := c.Values()
	contact := c.UpdateContact(conn, contactId)
	fmt.Printf("ID\tName\tE-Mail\t\tAddress\t\tPhone Number\tCreated Date\t\t\tUpdated Date\n")
	for contact.Next(){
		err := contact.Scan(&contactId, &values.name, &values.email, &values.address, &values.phoneNumber, &values.createdAt, &values.updatedAt)
		if err != nil {
			return
		}
		fmt.Printf("%d\t%s\t%s\t%s\t%d\t%s\t%s\n",
			contactId, values.name, values.email, values.address, values.phoneNumber, values.createdAt.String(), values.updatedAt.String())
	}
}

func DeleteContact(c ContactListInterface, conn *sqlx.DB, contactId int){
	c.DeleteContact(conn, contactId)
}

//Task list

func CreateTask(t TaskListInterface, conn *sqlx.DB) {
	t.CreateTaskList(conn)
}

func GetTask(t TaskListInterface, conn *sqlx.DB) {
	values := t.Values()
	var taskID int
	_, err := fmt.Scanf("%d", &taskID)
	if err != nil {
		fmt.Print(err)
		return
	}
	task := t.GetTask(conn, taskID)
	fmt.Printf("ID\tTask\t\tCompleted\tContact ID\tCreated Date\t\t\t\tUpdated Date\n")
	for task.Next(){
		err := task.Scan(&taskID, &values.task, &values.completed, &values.contactId, &values.createdAt, &values.updatedAt)
		if err != nil {
			return
		}
		fmt.Printf("%d\t%s\t%v\t\t%d\t\t%s\t%s\n",
			taskID, values.task, values.completed, values.contactId, values.createdAt.String(), values.updatedAt.String())
	}
}

func GetTaskList(t TaskListInterface, conn *sqlx.DB) {
	var taskID int
	values := t.Values()
	tasks := t.GetTaskList(conn)
	fmt.Printf("ID\tTask\t\tCompleted\tContact ID\tCreated Date\t\t\t\tUpdated Date\n")
	for tasks.Next(){
		err := tasks.Scan(&taskID, &values.task, &values.completed, &values.contactId, &values.createdAt, &values.updatedAt)
		if err != nil {
			return
		}
		fmt.Printf("%d\t%s\t%v\t\t%d\t\t%s\t%s\n",
			taskID, values.task, values.completed, values.contactId, values.createdAt.String(), values.updatedAt.String())
	}
}

func UpdateTask(c TaskListInterface, conn *sqlx.DB, taskID int) {
	values := c.Values()
	task := c.UpdateTask(conn, taskID)
	fmt.Printf("ID\tTask\t\tCompleted\tContact ID\tCreated Date\t\t\t\tUpdated Date\n")
	for task.Next(){
		err := task.Scan(&taskID, &values.task, &values.completed, &values.contactId, &values.createdAt, &values.updatedAt)
		if err != nil {
			return
		}
		fmt.Printf("%d\t%s\t%v\t\t%d\t\t%s\t%s\n",
			taskID, values.task, values.completed, values.contactId, values.createdAt.String(), values.updatedAt.String())
	}
}

func DeleteTask(c TaskListInterface, conn *sqlx.DB, taskID int){
	c.DeleteTask(conn, taskID)
}