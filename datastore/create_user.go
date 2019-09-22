package datastore

import (
	"deli/user-service/model"
	"errors"
	"github.com/labstack/gommon/log"
)

func (m *MysqlDS) SaveUser(user *model.User) error {
	tx, err := m.Begin()

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

	stmt.Exec(user)
}
