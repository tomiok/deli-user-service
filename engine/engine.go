package engine

import "deli/user-service/model"

type Spec interface {
	SaveUser() *model.User
	GetById(uid string) *model.User
}

type Engine struct {
}

func (e *Engine) SaveUser() *model.User {

	return nil
}

func (e *Engine) GetById(uid string) *model.User {
	return nil
}
