package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
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
	CreateTaskList(conn *pgxpool.Pool)
	GetTask(conn *pgxpool.Pool, taskID int) pgx.Rows
	GetTaskList(conn *pgxpool.Pool) pgx.Rows
	UpdateTask(conn *pgxpool.Pool, taskID int) pgx.Rows
	DeleteTask(conn *pgxpool.Pool, taskID int)
	Values() TaskList
}

func (c TaskList) Values() TaskList{
	return c
}

func (c *TaskList) CreateTaskList(conn *pgxpool.Pool) {
	_ , err := conn.Exec(context.Background(), "insert into tasklist (task, completed, contact_id, createdat, updated_at) VALUES($1, $2, $3, $4, $5)",
		c.task, c.completed, c.contactId, c.createdAt, c.updatedAt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query create failed: %v\n", err)
	} else {
		fmt.Println("Record created")
	}
}

func (c *TaskList) GetTask(conn *pgxpool.Pool, taskID int) pgx.Rows{
	rows, err := conn.Query(context.Background(), "select * from tasklist where id=$1", taskID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query get failed: %v\n", err)
	}
	return rows
}

func (c *TaskList) GetTaskList(conn *pgxpool.Pool) pgx.Rows {
	rows, err := conn.Query(context.Background(), "select * from tasklist order by id")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query get list failed: %v\n", err)
	}
	return rows
}

func (c *TaskList) UpdateTask(conn *pgxpool.Pool, taskID int) pgx.Rows {
	_, err := conn.Exec(context.Background(),
		"update tasklist set task=$2, updated_at=$3, contact_id=$4 where id=$1",
		taskID, c.task, c.updatedAt, c.contactId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query update failed: %v\n", err)
	} else {
		fmt.Println("Record updated")
	}
	rows, err := conn.Query(context.Background(), "select * from tasklist where id=$1", taskID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query get failed: %v\n", err)
	}
	return rows
}

func (c *TaskList) DeleteTask(conn *pgxpool.Pool, taskID int) {
	_, err := conn.Exec(context.Background(),
		"delete from tasklist where id=$1", taskID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query update failed: %v\n", err)
	}else {
		fmt.Println("Record deleted")
	}
}