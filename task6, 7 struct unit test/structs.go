package main

import (
	"fmt"
	"time"
)

type ContactList struct {
	id 					 int
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

type ContactManager struct {
	contacts []ContactList
}

func NewContactManager() ContactManager {
	manager := ContactManager{}
	manager.contacts = []ContactList{}
	return manager
}

func (manager *ContactManager) Create(contact *ContactList) *ContactList {
	contact.updatedAt = time.Now()
	contact.createdAt = time.Now()
	manager.contacts = append(manager.contacts, *contact)
	id := len(manager.contacts) - 1
	return &manager.contacts[id]
}

func (manager *ContactManager) Update(contact *ContactList) *ContactList {
	newContact := &manager.contacts[contact.id]
	newContact.name = contact.name
	newContact.email = contact.email
	newContact.address = contact.address
	newContact.phoneNumber = contact.phoneNumber
	newContact.updatedAt = time.Now()
	return newContact
}

func (manager *ContactManager) Delete(contactId int) {
	manager.contacts = append(manager.contacts[:contactId], manager.contacts[contactId+1:]...)
}

func (contact *ContactManager) GetAllContactList() []ContactList{
	return contact.contacts
}

func (manager *ContactManager) List() {
	for _, c := range manager.contacts {
		c.Detail()
	}
}

func (contact *ContactList) Detail() {
	fmt.Println("\nid:", contact.id)
	fmt.Println("name:", contact.name)
	fmt.Println("phone:", contact.phoneNumber)
	fmt.Println("email:", contact.email)
	fmt.Println("address:", contact.address)
	fmt.Println("created at:", contact.createdAt)
	fmt.Println("updated at:", contact.updatedAt)
	fmt.Println()
}

func main() {

}
