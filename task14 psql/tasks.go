package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func CreateContact(c ContactListInterface, conn *sqlx.DB) error {
	return c.CreateContactList(conn)
}

// GetContact - Get contact from db by ID pk
func GetContact(c ContactListInterface, conn *sqlx.DB, contactId int) error {
	values := c.Values()
	contact, err2 := c.GetContact(conn, contactId)
	if err2 != nil {
		return err2
	}
	fmt.Printf("ID\tName\tE-Mail\t\tAddress\t\tPhone Number\tCreated Date\t\t\tUpdated Date\n")
	for contact.Next(){
		err3 := contact.Scan(&contactId, &values.name, &values.email, &values.address, &values.phoneNumber, &values.createdAt, &values.updatedAt)
		if err3 != nil {
			return err3
		}
		fmt.Printf("%d\t%s\t%s\t%s\t%d\t%s\t%s\n",
			contactId, values.name, values.email, values.address, values.phoneNumber, values.createdAt.String(), values.updatedAt.String())
	}
	return nil
}

func GetContactList(contact ContactListInterface, conn *sqlx.DB) error {
	var contactId int
	values := contact.Values()
	contacts, err1 := contact.GetContactList(conn)
	if err1 != nil {
		return err1
	}
	fmt.Printf("ID\tName\tE-Mail\t\tAddress\t\tPhone Number\tCreated Date\t\t\tUpdated Date\n")
	for contacts.Next(){
		err := contacts.Scan(&contactId, &values.name, &values.email, &values.address, &values.phoneNumber, &values.createdAt, &values.updatedAt)
		if err != nil {
			return err
		}
		fmt.Printf("%d\t%s\t%s\t%s\t%d\t%s\t%s\n",
			contactId, values.name, values.email, values.address, values.phoneNumber, values.createdAt.String(), values.updatedAt.String())
	}
	return nil
}

func UpdateContact(c ContactListInterface, conn *sqlx.DB, contactId int) error {
	values := c.Values()
	contact, err1 := c.UpdateContact(conn, contactId)
	if err1 != nil {
		return err1
	}
	fmt.Printf("ID\tName\tE-Mail\t\tAddress\t\tPhone Number\tCreated Date\t\t\tUpdated Date\n")
	for contact.Next(){
		err := contact.Scan(&contactId, &values.name, &values.email, &values.address, &values.phoneNumber, &values.createdAt, &values.updatedAt)
		if err != nil {
			return err
		}
		fmt.Printf("%d\t%s\t%s\t%s\t%d\t%s\t%s\n",
			contactId, values.name, values.email, values.address, values.phoneNumber, values.createdAt.String(), values.updatedAt.String())
	}
	return nil
}

func DeleteContact(c ContactListInterface, conn *sqlx.DB, contactId int) error {
	return c.DeleteContact(conn, contactId)
}

//Task list

func CreateTask(t TaskListInterface, conn *sqlx.DB) error {
	return t.CreateTaskList(conn)
}

func GetTask(t TaskListInterface, conn *sqlx.DB) error {
	values := t.Values()
	var taskID int
	_, err := fmt.Scanf("%d", &taskID)
	if err != nil {
		return err
	}
	task, err1 := t.GetTask(conn, taskID)
	if err1 != nil {
		return err1
	}
	fmt.Printf("ID\tTask\t\tCompleted\tContact ID\tCreated Date\t\t\t\tUpdated Date\n")
	for task.Next(){
		err2 := task.Scan(&taskID, &values.task, &values.completed, &values.contactId, &values.createdAt, &values.updatedAt)
		if err2 != nil {
			return err2
		}
		fmt.Printf("%d\t%s\t%v\t\t%d\t\t%s\t%s\n",
			taskID, values.task, values.completed, values.contactId, values.createdAt.String(), values.updatedAt.String())
	}
	return nil
}

func GetTaskList(t TaskListInterface, conn *sqlx.DB) error {
	var taskID int
	values := t.Values()
	tasks, err1 := t.GetTaskList(conn)
	if err1 != nil {
		return err1
	}
	fmt.Printf("ID\tTask\t\tCompleted\tContact ID\tCreated Date\t\t\t\tUpdated Date\n")
	for tasks.Next(){
		err := tasks.Scan(&taskID, &values.task, &values.completed, &values.contactId, &values.createdAt, &values.updatedAt)
		if err != nil {
			return err
		}
		fmt.Printf("%d\t%s\t%v\t\t%d\t\t%s\t%s\n",
			taskID, values.task, values.completed, values.contactId, values.createdAt.String(), values.updatedAt.String())
	}
	return nil
}

func UpdateTask(c TaskListInterface, conn *sqlx.DB, taskID int) error {
	values := c.Values()
	task, err1 := c.UpdateTask(conn, taskID)
	if err1 != nil {
		return err1
	}
	fmt.Printf("ID\tTask\t\tCompleted\tContact ID\tCreated Date\t\t\t\tUpdated Date\n")
	for task.Next(){
		err := task.Scan(&taskID, &values.task, &values.completed, &values.contactId, &values.createdAt, &values.updatedAt)
		if err != nil {
			return err
		}
		fmt.Printf("%d\t%s\t%v\t\t%d\t\t%s\t%s\n",
			taskID, values.task, values.completed, values.contactId, values.createdAt.String(), values.updatedAt.String())
	}
	return nil
}

func DeleteTask(c TaskListInterface, conn *sqlx.DB, taskID int) error {
	return c.DeleteTask(conn, taskID)
}