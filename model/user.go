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
			Uid:          genUUID(),
			Name:         name,
			LastName:     lastName,
			Username:     createUserName(name, lastName),
			Password:     password,
			City:         city,
			Country:      country,
			EmailAddress: email,
			CreatedAt:    time.Now(),
		}
	case "user":
		return &User{
			Uid:       genUUID(),
			Name:      name,
			LastName:  lastName,
			Username:  username,
			Password:  password,
			City:      city,
			Country:   country,
			CreatedAt: time.Now(),
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
	Uid          string
	Name         string
	LastName     string
	Username     string
	Password     string
	City         string
	Country      string
	EmailAddress string
	CreatedAt    time.Time
}

func genUUID() string {
	return uuid.New().String()
}
