package usecases

import (
	"example.com/hex-todo/backend/domain"
)

type TodoInteractor struct {
	Repository domain.TodoRepository
}

func (interactor *TodoInteractor) AddTodo(todo domain.Todo) (domain.Todo, error) {
	return interactor.Repository.Store(todo)
}

func (interactor *TodoInteractor) FindByID(id int) (domain.Todo, error) {
	return interactor.Repository.FindByID(id)
}

func (interactor *TodoInteractor) FindAll() ([]domain.Todo, error) {
	return interactor.Repository.FindAll()
}

func (interactor *TodoInteractor) UpdateTodo(id int, todo domain.Todo) (domain.Todo, error) {
	return interactor.Repository.Update(id, todo)
}

func (interactor *TodoInteractor) DeleteTodo(id int) error {
	return interactor.Repository.Delete(id)
}
