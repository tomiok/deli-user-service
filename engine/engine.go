package engine

import (
	"deli/user-service/datastore"
	"deli/user-service/model"
	"github.com/labstack/gommon/log"
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

func (e *Engine) GetById(uid string) *model.User {
	return nil
}

func (e *Engine) Save(f GetUser) (string, error) {
	id, err := e.repo.SaveUser(f())
	if err != nil {
		log.Errorf("cannot insert user")
		return "", err
	}

	return id, nil
}
