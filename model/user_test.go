package model

import (
	"testing"
)

var admin = (*ut)(&userType{title: "admin"})

func TestMap(t *testing.T) {
	adminUser := Map("tomas", "lingotti", "Rosario", "Argetina", "pass1", "", "tomi@msn.com", admin)

	if adminUser.UserType.title != "admin" {
		t.Error("should create an admin user")
		t.Fail()
	}
}
