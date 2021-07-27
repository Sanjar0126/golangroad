package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var manager = NewContactManager()

func TestContactManager_Create(t *testing.T) {
	var contact = &ContactList{
		id: 0,
		name: "Ivan",
		email: "ivan@mail.po",
		address: "Main Street, 10-4",
		phoneNumber: 43123412,
	}
	var newContact = manager.Create(contact)
	require.Equal(t, newContact.name, contact.name)
	require.Equal(t, newContact.email, contact.email)
	require.Equal(t, newContact.address, contact.address)
	require.Equal(t, newContact.phoneNumber, contact.phoneNumber)
	require.NotZero(t, newContact.createdAt)
	require.NotZero(t, newContact.updatedAt)

	contact = &ContactList{
		id: 1,
		name: "Gregoriy",
		email: "greg@mail.po",
		address: "Main Street, 1-43",
		phoneNumber: 54241234,
	}
	newContact = manager.Create(contact)
	require.Equal(t, newContact.name, contact.name)
	require.Equal(t, newContact.email, contact.email)
	require.Equal(t, newContact.address, contact.address)
	require.Equal(t, newContact.phoneNumber, contact.phoneNumber)
	require.NotZero(t, newContact.createdAt)
	require.NotZero(t, newContact.updatedAt)
}

func TestContactManager_Update(t *testing.T) {
	contact := &manager.contacts[0]
	contact.name = "Ivanko"
	contact.email = "ivanko@mail.po"
	newContact := manager.Update(contact)
	require.Equal(t, newContact.name, contact.name)
	require.Equal(t, newContact.email, contact.email)
	require.Equal(t, newContact.address, contact.address)
	require.Equal(t, newContact.phoneNumber, contact.phoneNumber)
	require.NotZero(t, newContact.createdAt)
	require.NotZero(t, newContact.updatedAt)

	contact = &manager.contacts[1]
	contact.email = "gregory@mail.by"
	contact.phoneNumber = 54241344
	newContact = manager.Update(contact)
	require.Equal(t, newContact.name, contact.name)
	require.Equal(t, newContact.email, contact.email)
	require.Equal(t, newContact.address, contact.address)
	require.Equal(t, newContact.phoneNumber, contact.phoneNumber)
	require.NotZero(t, newContact.createdAt)
	require.NotZero(t, newContact.updatedAt)
}

func TestContactManager_List(t *testing.T) {
	contacts := manager.GetAllContactList()
	require.Equal(t, contacts[0].name, "Ivanko")
	require.Equal(t, contacts[1].phoneNumber, 54241344)
}

func TestContactManager_Delete(t *testing.T) {
	manager.Delete(0)
	if len(manager.contacts) != 1 {
		t.Error("Error while deleting contact")
	}
}