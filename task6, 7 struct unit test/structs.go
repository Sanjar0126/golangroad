package main

import "time"

type ContactList struct {
	name, email, address string
	phoneNumber          int
	createdAt            time.Time
	updatedAt			 time.Time
}

type TaskList struct {
	contactId int
	task      string
	completed bool
	createdAt time.Time
	updatedAt time.Time
}

func main() {

}
