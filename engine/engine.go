package engine

import (
	"github.com/deli/user-service/datastore"
	"github.com/deli/user-service/model"
	"github.com/deli/user-service/logs"
)

type GetUser func() *model.User

type Spec interface {
	Save(f GetUser) (string, error)
	GetById(uid string) *model.User
}

func New(saveRepo *datastore.SaveUserRepo) Spec {
	return &Engine{
		repo: saveRepo,
	}
}

type Engine struct {
	repo *datastore.SaveUserRepo
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
