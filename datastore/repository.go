package datastore

import "github.com/deli/user-service/models"

type RepositorySpec interface {
	SaveUser(u *models.User) (string, error)
	GetUserById(id string) (*models.User, error)
	ValidateUserByPassword(username, password string) (string, error)
}

type UserRepository struct {
	DS *MysqlDS
}
