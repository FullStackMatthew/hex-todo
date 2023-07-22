package domain 

import "time"

type Todo struct {
	ID int
	Title string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TodoRepository interface {
	Store(todo Todo) (Todo, error)
	FindAll() ([]Todo, error)
	FindByID(id int) (Todo, error)
	Update(id int, todo Todo) (Todo, error)
	Delete(id int) error
}