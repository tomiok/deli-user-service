package engine

import (
	"deli/user-service/datastore"
	"deli/user-service/model"
	"github.com/labstack/gommon/log"
)

type Spec interface {
	Save(u *model.User) (string, error)
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

func (e *Engine) Save(u *model.User) (string, error) {
	id, err := e.repo.SaveUser(u)
	if err != nil {
		log.Errorf("cannot insert user")
		return "", err
	}

	return id, nil
}
