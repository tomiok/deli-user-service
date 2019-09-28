package datastore

import (
	"deli/user-service/model"
	"errors"
	"github.com/labstack/gommon/log"
)

type SaveUser interface {
	SaveUser(u *model.User) error
}

type SaveUserRepo struct {
	DS *MysqlDS
}

func (u SaveUserRepo) SaveUser(user *model.User) error {

	tx, err := u.DS.Begin()

	if err != nil {
		return errors.New(err.Error())
	}

	defer func() {
		err := tx.Rollback()
		if err != nil {
			log.Error("cannot rollback transaction")
		}
	}()

	stmt, err := tx.Prepare("INSERT INTO deli_user.`user`" +
		" (id, name, last_name, password, username, city, country, email, created_at, user_type) " +
		"VALUES (?,?,?,?,?,?,?,?,now(),?)")
	if err != nil {
		return err
	}

	defer func() {
		err := stmt.Close()
		if err != nil {
			log.Error("cannot close the statement")
		}
	}()

	_, err = stmt.Exec("idgen", user.Name, user.LastName, user.Password, user.Username, user.City,
		user.Country, user.EmailAddress, user.UserType.Title)

	if err != nil {
		log.Errorf("cannot execute prepared statement: %s", err.Error())
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}
