package main

import "time"

type ContactList struct {
	name, email, address string
	phoneNumber         int
	createdAt time.Time
}

func (c ContactList) CreateContactList() {
}