package datastore

import "github.com/deli/user-service/model"

type RepositorySpec interface {
	SaveUser(u *model.User) (string, error)
	GetUserById(id string) (*model.User, error)
	ValidateUserByPassword(username, password string) (string, error)
}

type UserRepository struct {
	DS *MysqlDS
}
