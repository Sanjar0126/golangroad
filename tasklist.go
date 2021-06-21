package main

import "time"

type TaskList struct {
	contact   ContactList
	task      string
	completed bool
	createdAt time.Time
}
