package engine

import (
	"github.com/deli/user-service/commons/logs"
	"github.com/deli/user-service/datastore"
	"github.com/deli/user-service/models"
)

type GetUser func() *models.User

type Spec interface {
	Save(f GetUser) (string, error)
	GetById(uid string) *models.User
	ValidateUser(username, password string) (string, error)
}

func New(userRepo *datastore.UserRepository) Spec {
	return &Engine{
		repo: userRepo,
	}
}

type Engine struct {
	repo datastore.RepositorySpec
}

func (e *Engine) GetById(id string) *models.User {
	u, err := e.repo.GetUserById(id)

	if err != nil {
		return nil
	}

	return u
}

func (e *Engine) Save(f GetUser) (string, error) {
	id, err := e.repo.SaveUser(f())
	if err != nil {
		logs.Errorf("cannot insert user")
		return "", err
	}

	return id, nil
}

func (e *Engine) ValidateUser(username, password string) (string, error) {
	return e.repo.ValidateUserByPassword(username, password)
}
