package main

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
)

type TaskList struct {
	contactId int
	task      string
	completed bool
	createdAt time.Time
	updatedAt time.Time
}

type TaskListInterface interface {
	CreateTaskList(conn *sqlx.DB)
	GetTask(conn *sqlx.DB, taskID int) *sql.Rows
	GetTaskList(conn *sqlx.DB) *sql.Rows
	UpdateTask(conn *sqlx.DB, taskID int) *sql.Rows
	DeleteTask(conn *sqlx.DB, taskID int)
	Values() TaskList
}

func (c TaskList) Values() TaskList{
	return c
}

func (c *TaskList) CreateTaskList(conn *sqlx.DB) {
	_ , err := conn.Exec("insert into tasklist (task, completed, contact_id, createdat, updated_at) VALUES($1, $2, $3, $4, $5)",
		c.task, c.completed, c.contactId, c.createdAt, c.updatedAt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query create failed: %v\n", err)
	} else {
		fmt.Println("Record created")
	}
}

func (c *TaskList) GetTask(conn *sqlx.DB, taskID int) *sql.Rows{
	rows, err := conn.Query("select * from tasklist where id=$1", taskID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query get failed: %v\n", err)
	}
	return rows
}

func (c *TaskList) GetTaskList(conn *sqlx.DB) *sql.Rows {
	rows, err := conn.Query("select * from tasklist order by id")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query get list failed: %v\n", err)
	}
	return rows
}

func (c *TaskList) UpdateTask(conn *sqlx.DB, taskID int) *sql.Rows {
	_, err := conn.Exec("update tasklist set task=$2, updated_at=$3, contact_id=$4 where id=$1",
		taskID, c.task, c.updatedAt, c.contactId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query update failed: %v\n", err)
	} else {
		fmt.Println("Record updated")
	}
	rows, err := conn.Query("select * from tasklist where id=$1", taskID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query get failed: %v\n", err)
	}
	return rows
}

func (c *TaskList) DeleteTask(conn *sqlx.DB, taskID int) {
	_, err := conn.Exec("delete from tasklist where id=$1", taskID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query update failed: %v\n", err)
	}else {
		fmt.Println("Record deleted")
	}
}