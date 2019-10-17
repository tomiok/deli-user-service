package datastore

import (
	"errors"
	"github.com/deli/user-service/models"
	"github.com/deli/user-service/token"
)

func (u *UserRepository) ValidateUserByPassword(username, password string) (string, error) {

	tx, err := u.DS.Begin()

	if err != nil {
		return "", err
	}

	var user models.User
	err = tx.QueryRow("select u.id from user u where u.username = ? and u.password = ?", username, password).
		Scan(&user.Uid)

	if err != nil {
		return "", errors.New("query errors:: " + err.Error())
	}

                                 	jsonToken, err := validateDBQuery(user.Uid)

	return jsonToken, err
}

func validateDBQuery(id string) (string, error) {
	if id == "" {
		return "", errors.New("user not found in database")
	}

	return token.Encode(id)
}
