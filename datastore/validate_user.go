package datastore

import (
	"errors"
	"github.com/deli/user-service/model"
)

func (u *UserRepository) ValidateUserByPassword(username, password string) (string, error) {

	tx, err := u.DS.Begin()

	if err != nil {
		return "", err
	}

	var user model.User
	err = tx.QueryRow("select u.id from user u where u.username = ? and u.password = ?").
		Scan(&user.Uid, &user.Name, &user.LastName)

	if err != nil {
		return "", errors.New("cannot execute the query " + err.Error())
	}

	return "tokem", nil
}
