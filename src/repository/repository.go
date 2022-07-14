package repository

import "api-controle/src/model"

type IRepository interface {
	Insert(model model.IModel) (uint64, error)

	Update(model model.IModel) error

	Save(model model.IModel) (uint64, error)

	SaveAll(models interface{}) error

	FindById(receiver model.IModel, id interface{}) error

	FindFirst(receiver model.IModel, where string, args ...interface{}) error

	FindAll(models interface{}, where string, args ...interface{}) (err error)

	Delete(model model.IModel, where string, args ...interface{}) error
}
