package datastore

import (
	"errors"
	"github.com/deli/user-service/commons/logs"
	"github.com/deli/user-service/models"
)

func (u *UserRepository) GetUserById(id string) (*models.User, error) {
	tx, err := u.DS.Begin()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	stmt, err := tx.Prepare("select u.id, u.name, u.last_name, u.username, u.email, u.city, u.country, " +
		"u.created_at from user u where id = ?")

	if err != nil {
		return nil, err
	}

	defer func() {
		err := stmt.Close()
		if err != nil {
			logs.Error("cannot close the statement")
		}
	}()

	rows := stmt.QueryRow(id)

	var user models.User

	//TODO add user type
	err = rows.Scan(&user.Uid, &user.Name, &user.LastName,
		&user.Username, &user.EmailAddress, &user.City, &user.Country, &user.CreatedAt)

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	if user.Uid == "" {
		return nil, errors.New("cannot found user")
	}

	return &user, nil
}
