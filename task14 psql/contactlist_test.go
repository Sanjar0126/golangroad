package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	"testing"
	"time"
)

var connection *sqlx.DB

func TestConnection(t *testing.T){
	connection = connectDB()
}

func TestCreateContact(t *testing.T) {
	var testContact = ContactList{
		name:        RandomString(10),
		email:       RandomString(7) + "@mail.com",
		address:     RandomString(25),
		phoneNumber: RandomInt(1000000, 9999999),
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}
	err := CreateContact(&testContact, connection)
	require.NoError(t, err)
	require.NotEmpty(t, testContact)
}

func TestGetContact(t *testing.T){

}

func RandomString(n int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}