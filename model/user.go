package model

import (
	"crypto/sha256"
	"deli/user-service/commons"
	"encoding/hex"
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
			Password:     encryptPass(password),
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
	rand.Seed(time.Now().UnixNano())
	var username string
	if lastName == "" {
		username = name + strconv.Itoa(rand.Intn(999))
	} else {
		username = string(name[0]) + lastName
	}

	return strings.ToLower(username)
}

func encryptPass(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	sha256Hash := hex.EncodeToString(h.Sum(nil))
	return sha256Hash
}

type User struct {
	Uid          string    `json:"uid"`
	Name         string    `json:"name"`
	LastName     string    `json:"last_name"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	City         string    `json:"city"`
	Country      string    `json:"country"`
	EmailAddress string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
	UserType     *UserType `json:"user_type"`
}
