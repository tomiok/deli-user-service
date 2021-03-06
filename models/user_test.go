package models

import (
	"github.com/deli/user-service/commons"
	"testing"
)

var (
	admin  = (*ut)(&UserType{Title: "admin"})
	writer = (*ut)(&UserType{Title: "writer"})
)

func TestMapForAdmin(t *testing.T) {
	adminUser := Map("tomas", "lingotti", "Rosario", "Argetina", "pass1", "", "tomi@msn.com", admin)

	if adminUser.UserType.Title != "admin" {
		t.Error("should create an admin user")
		t.Fail()
	}

	if adminUser.Password == "pass1" {
		t.Error("Password is not encrypted!")
		t.Fail()
	}
}

func TestMapForWriter(t *testing.T) {
	writerUser := Map("tomas", "lingotti", "Rosario", "Argetina", "pass1", "", "tomi@msn.com", writer)

	if writerUser.UserType.Title != "writer" {
		t.Error("should create an admin user")
		t.Fail()
	}

	if writerUser.Password == "pass1" {
		t.Error("password is not encrypted!")
		t.Fail()
	}
}

func TestEncrypt(t *testing.T) {
	pass := "solidPassword"
	e := commons.EncryptPass(pass)
	ee := commons.EncryptPass(pass)

	if e == pass {
		t.Fail()
	}

	if e != ee {
		t.Fail()
	}
}

func TestUsernameCreation(t *testing.T) {
	name := "Sherman"
	lastName := "Lewis"
	lastNameEmpty := ""

	username1 := createUserName(name, lastNameEmpty)
	username2 := createUserName(name, lastName)

	if username1 == username2 {
		t.Fail()
	}

	if "slewis" != username2 {
		t.Fail()
	}
}
