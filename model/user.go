package model

import (
	"github.com/google/uuid"
	"time"
)

type ut userType

type userType struct {
	title string
}

func Map(name string, lastName string, city string, country string, password string, username string,
	email string, ut ut) *User {
	uType := ut.title

	switch uType {
	case "admin", "writer":
		return &User{
			uid:          genUUID(),
			name:         name,
			lastName:     lastName,
			username:     createUserName(name, lastName),
			password:     password,
			city:         city,
			country:      country,
			emailAddress: email,
			createdAt:    time.Now(),
		}
	case "user":
		return &User{
			uid:       genUUID(),
			name:      name,
			lastName:  lastName,
			username:  username,
			password:  password,
			city:      city,
			country:   country,
			createdAt: time.Now(),
		}

	default:
		panic("Cannot enter another type rather than admin, writer or user")
	}
}

// TODO fix this and short the first name to one character
func createUserName(name string, lastName string) string {
	return name + lastName
}

type User struct {
	uid          string
	name         string
	lastName     string
	username     string
	password     string
	city         string
	country      string
	emailAddress string
	createdAt    time.Time
}

func genUUID() string {
	return uuid.New().String()
}
