package model

import (
	"testing"
)

var (
	admin = (*ut)(&userType{title: "admin"})
	writer = (*ut)(&userType{title: "writer"})
)

func TestMapForAdmin(t *testing.T) {
	adminUser := Map("tomas", "lingotti", "Rosario", "Argetina", "pass1", "", "tomi@msn.com", admin)

	if adminUser.UserType.title != "admin" {
		t.Error("should create an admin user")
		t.Fail()
	}
}

func TestMapForWriter(t *testing.T) {
	adminUser := Map("tomas", "lingotti", "Rosario", "Argetina", "pass1", "", "tomi@msn.com", writer)

	if adminUser.UserType.title != "writer" {
		t.Error("should create an admin user")
		t.Fail()
	}
}

func TestEncrypt(t *testing.T) {
	pass := "solidPassword"
	e := encryptPass(pass)
	ee := encryptPass(pass)

	if e == pass {
		t.Fail()
	}

	if e != ee {
		t.Fail()
	}
}