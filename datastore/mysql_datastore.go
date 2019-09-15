package datastore

import (
	"database/sql"
	"github.com/labstack/gommon/log"
	_ "github.com/go-sql-driver/mysql"
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

	return &MysqlDS{DB: connection}, nil
}
