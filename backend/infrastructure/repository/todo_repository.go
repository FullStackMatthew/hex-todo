package repository

import (
	"database/sql"
	"time"
	"github.com/sirupsen/logrus"
	"example.com/hex-todo/backend/domain"
)

type TodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(DB *sql.DB) *TodoRepository {
	return &TodoRepository{DB: DB}
}

func (r *TodoRepository) Store(todo domain.Todo) (domain.Todo, error) {
	query := `
		INSERT INTO todos (title, completed, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	
	err := r.DB.QueryRow(query, todo.Title, todo.Completed, time.Now(), time.Now()).Scan(&todo.ID)
	if err != nil {
		logrus.Error(err)
		return domain.Todo{}, err
	}
	return todo, nil
}

func (r *TodoRepository) FindAll() ([]domain.Todo, error) {
	todos := []domain.Todo{}
	query := `SELECT * FROM todos`
	rows, err := r.DB.Query(query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo domain.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			logrus.Error(err)
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepository) FindByID(id int) (domain.Todo, error) {
	var todo domain.Todo
	query := `SELECT * FROM todos WHERE id = $1`
	if err := r.DB.QueryRow(query, id).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		logrus.Error(err)
		return domain.Todo{}, err
	}
	return todo, nil
}

func (r *TodoRepository) Update(id int, todo domain.Todo) (domain.Todo, error) {
	query := `
		UPDATE todos
		SET title = $1, completed = $2, updated_at = $3
		WHERE id = $4
		RETURNING id, created_at`

	err := r.DB.QueryRow(query, todo.Title, todo.Completed, time.Now(), id).Scan(&todo.ID, &todo.CreatedAt)
	if err != nil {
		logrus.Error(err)
		return domain.Todo{}, err
	}
	return todo, nil
}

func (r *TodoRepository) Delete(id int) error {
	query := `DELETE FROM todos WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		logrus.Error(err)
	}
	return err
}