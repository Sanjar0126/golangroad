package main

import "time"

type ContactList struct {
	name, email, address string
	phoneNumber          int
	createdAt            time.Time
}

func (c ContactList) CreateContactList() {
	
}

type TaskList struct {
	contact   ContactList
	task      string
	completed bool
	createdAt time.Time
}
