package engine

import (
	"deli/user-service/datastore"
	"deli/user-service/model"
)

type Spec interface {
	Save(u *model.User)
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

func (e *Engine) Save(u *model.User) {
	err := e.repo.SaveUser(u)
	if err != nil {
		panic(err)
	}
}
