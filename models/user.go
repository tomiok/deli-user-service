package models

import (
	"github.com/deli/user-service/commons"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type ut UserType

type UserType struct {
	Title string
}

func (u *User) DoMap() *User {
	return Map(u.Name, u.LastName, u.City, u.Country, u.Password, u.Username, u.EmailAddress, (*ut)(u.UserType))
}

func Map(name string, lastName string, city string, country string, password string, username string,
	email string, ut *ut) *User {
	uType := ut.Title

	switch uType {
	case "admin", "writer":
		return &User{
			Uid:          commons.StringUUID(),
			Name:         name,
			LastName:     lastName,
			Username:     createUserName(name, lastName),
			Password:     commons.EncryptPass(password),
			City:         city,
			Country:      country,
			EmailAddress: email,
			CreatedAt:    time.Now(),
			UserType:     &UserType{Title: ut.Title},
		}
	case "user":
		return &User{
			Uid:       commons.StringUUID(),
			Name:      name,
			LastName:  lastName,
			Username:  username,
			Password:  commons.EncryptPass(password),
			City:      city,
			Country:   country,
			CreatedAt: time.Now(),
			UserType:  &UserType{Title: ut.Title},
		}

	default:
		panic("cannot enter another type rather than admin, writer or user")
	}
}

func createUserName(name string, lastName string) string {
	rand.Seed(time.Now().UnixNano())
	var username string
	if lastName == "" {
		username = name + strconv.Itoa(rand.Intn(999))
	} else {
		username = string(name[0]) + lastName
	}

	return strings.ToLower(username)
}

type User struct {
	Uid          string    `json:"uid,omitempty"`
	Name         string    `json:"name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	Username     string    `json:"username,omitempty"`
	Password     string    `json:"password,omitempty"`
	City         string    `json:"city,omitempty"`
	Country      string    `json:"country,omitempty"`
	EmailAddress string    `json:"email,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UserType     *UserType `json:"user_type,omitempty"`
}