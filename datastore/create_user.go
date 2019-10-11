package datastore

import (
	"errors"
	"github.com/deli/user-service/commons/logs"
	"github.com/deli/user-service/model"
)

// SaveUSer saves a user in the DB, and returns the ID for that user.
func (u *UserRepository) SaveUser(user *model.User) (string, error) {

	tx, err := u.DS.Begin()

	if err != nil {
		return "", errors.New(err.Error())
	}

	stmt, err := tx.Prepare("INSERT INTO `user`" +
		" (id, name, last_name, password, username, city, country, email, created_at, user_type) " +
		"VALUES (?,?,?,?,?,?,?,?,now(),?)")
	if err != nil {
		return "", err
	}

	defer func() {
		err := stmt.Close()
		if err != nil {
			logs.Error("cannot close the statement")
		}
	}()

	id := user.Uid
	_, err = stmt.Exec(id, user.Name, user.LastName, user.Password, user.Username, user.City,
		user.Country, user.EmailAddress, user.UserType.Title)

	if err != nil {
		logs.Errorf("cannot execute prepared statement: %s", err.Error())
		return "", err
	}

	err = tx.Commit()

	if err != nil {
		return "", err
	}

	return id, nil
}
