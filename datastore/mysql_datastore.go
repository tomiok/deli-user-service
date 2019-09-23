package datastore

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

type MysqlDS struct {
	*sql.DB
}

func NewMysqlDS(source string) (*MysqlDS, error) {

	connection, err := sql.Open("mysql", source)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Infof("Connection mysql %s", source)

	return &MysqlDS{DB: connection}, nil
}
