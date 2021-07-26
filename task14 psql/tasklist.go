package main

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
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
	CreateTaskList(conn *sqlx.DB) error
	GetTask(conn *sqlx.DB, taskID int) (*sql.Rows, error)
	GetTaskList(conn *sqlx.DB) (*sql.Rows, error)
	UpdateTask(conn *sqlx.DB, taskID int) (*sql.Rows, error)
	DeleteTask(conn *sqlx.DB, taskID int) error
	Values() TaskList
}

func (c TaskList) Values() TaskList{
	return c
}

func (c *TaskList) CreateTaskList(conn *sqlx.DB) error {
	_ , err := conn.Exec("insert into tasklist (task, completed, contact_id, createdat, updated_at) VALUES($1, $2, $3, $4, $5)",
		c.task, c.completed, c.contactId, c.createdAt, c.updatedAt)
	return err
}

func (c *TaskList) GetTask(conn *sqlx.DB, taskID int) (*sql.Rows, error){
	rows, err := conn.Query("select * from tasklist where id=$1", taskID)
	return rows, err
}

func (c *TaskList) GetTaskList(conn *sqlx.DB) (*sql.Rows, error) {
	rows, err := conn.Query("select * from tasklist order by id")
	return rows, err
}

func (c *TaskList) UpdateTask(conn *sqlx.DB, taskID int) (*sql.Rows, error) {
	_, err := conn.Exec(
		"update tasklist set task=$2, updated_at=$3, contact_id=$4 where id=$1",
		taskID, c.task, c.updatedAt, c.contactId)
	if err != nil {
		return nil, err
	}
	rows, err1 := conn.Query("select * from tasklist where id=$1", taskID)
	if err1 != nil {
		return nil, err1
	}
	return rows, nil
}

func (c *TaskList) DeleteTask(conn *sqlx.DB, taskID int) error {
	_, err := conn.Exec(
		"delete from tasklist where id=$1", taskID)
	return err
}