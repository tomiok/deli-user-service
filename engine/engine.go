package engine

import (
	"github.com/deli/user-service/commons/logs"
	"github.com/deli/user-service/datastore"
	"github.com/deli/user-service/model"
)

type GetUser func() *model.User

type Spec interface {
	Save(f GetUser) (string, error)
	GetById(uid string) *model.User
}

func New(userRepo *datastore.UserRepository) Spec {
	return &Engine{
		repo: userRepo,
	}
}

type Engine struct {
	repo *datastore.UserRepository
}

func (e *Engine) GetById(id string) *model.User {
	u, err := e.repo.GetUserById(id)

	if err != nil {

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
