package datastore

import (
	"errors"
	"github.com/deli/user-service/model"
	"github.com/labstack/gommon/log"
)

type SaveUser interface {
	SaveUser(u *model.User) (string, error)
	GetUserById(id string) (*model.User, error)
}

type SaveUserRepo struct {
	DS *MysqlDS
}

// SaveUser saves a user in the DB, and returns the ID for that user.
func (u *SaveUserRepo) SaveUser(user *model.User) (string, error) {

	tx, err := u.DS.Begin()

	if err != nil {
		return "", errors.New(err.Error())
	}

	stmt, err := tx.Prepare("INSERT INTO deli_user.`user`" +
		" (id, name, last_name, password, username, city, country, email, created_at, user_type) " +
		"VALUES (?,?,?,?,?,?,?,?,now(),?)")
	if err != nil {
		return "", err
	}

	defer func() {
		err := stmt.Close()
		if err != nil {
			log.Error("cannot close the statement")
		}
	}()

	id := user.Uid
	_, err = stmt.Exec(id, user.Name, user.LastName, user.Password, user.Username, user.City,
		user.Country, user.EmailAddress, user.UserType.Title)

	if err != nil {
		log.Errorf("cannot execute prepared statement: %s", err.Error())
		return "", err
	}

	err = tx.Commit()

	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *SaveUserRepo) GetUserById(id string) (*model.User, error) {
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
			log.Error("cannot close the statement")
		}
	}()

	rows := stmt.QueryRow(id)

	var user model.User

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
