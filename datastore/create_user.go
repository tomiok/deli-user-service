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
		return errors.New("cannot start transaction")
	}

	defer func() {
		err := tx.Rollback()
		if err != nil {
			log.Error("cannot rollback transaction")
		}
	}()

	stmt, err := tx.Prepare("INSERT INTO `user` VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	defer func() {
		err := stmt.Close()
		if err != nil {
			log.Error("cannot close the statement")
		}
	}()

	_, err = stmt.Exec(user)

	if err != nil {
		log.Fatal("cannot execute prepared statement")
		return err
	}

	err = tx.Commit()

	if err != nil{
		return err
	}

	return nil

}
