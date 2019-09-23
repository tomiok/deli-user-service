package model

import (
	"crypto/sha256"
	"github.com/google/uuid"
	"math/rand"
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
			Uid:          genUUID(),
			Name:         name,
			LastName:     lastName,
			Username:     createUserName(name, lastName),
			Password:     encryptPass(password),
			City:         city,
			Country:      country,
			EmailAddress: email,
			CreatedAt:    time.Now(),
			UserType:     &UserType{Title: ut.Title},
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
			UserType:  &UserType{Title: ut.Title},
		}

	default:
		panic("Cannot enter another type rather than admin, writer or user")
	}
}

func createUserName(name string, lastName string) string {
	if lastName == ""{
		return string(name[0]) + string(rand.Intn(100))
	}
	return string(name[0]) + lastName
}

func encryptPass(password string) string {
	sum := sha256.Sum256([]byte(password))
	return string(sum[:])
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
	UserType     *UserType
}

func genUUID() string {
	return uuid.New().String()
}
