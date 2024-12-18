package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"simple_go_app/model"
)

func NewTodoRepository() *Repository {
	repo := newUserRepository()
	return repo
}

func (repo *Repository) GetAll() ([]model.Todo, error) {
	rows, err := repo.DB.Query("SELECT * FROM todo ORDER BY created_at DESC")
	if err != nil {
		return nil, fmt.Errorf("error getting all todos: %s\n", err)
	}

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Done,
			&todo.CreatedAt,
			&todo.UpdatedAt); err != nil {
			return nil, fmt.Errorf("could not parse todo: %s\n", err)
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (repo *Repository) GetByID(id string) (*model.Todo, error) {
	row := repo.DB.QueryRow("SELECT * FROM todo WHERE id = $1", id)
	var todo model.Todo
	if err := row.Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.Done,
		&todo.CreatedAt,
		&todo.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not parse todo: %s\n", err)
	}
	return &todo, nil
}

func (repo *Repository) Create(todo *model.Todo) error {
	_, err := repo.DB.Exec("INSERT INTO todo (title, description) VALUES ($1, $2)",
		todo.Title,
		todo.Description)
	if err != nil {
		return fmt.Errorf("error executing query todo: %s\n", err)
	}
	return nil
}

func (repo *Repository) Update(todo *model.Todo) (int, error) {
	result, err := repo.DB.Exec("UPDATE todo SET title = $1, description = $2, is_completed = $3, created_at = $4 WHERE id = $5",
		todo.Title,
		todo.Description,
		todo.Done,
		todo.CreatedAt,
		todo.ID)
	if err != nil {
		return 0, fmt.Errorf("error executing query todo: %s\n", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error getting rows affected: %s\n", err)
	}
	return int(rowsAffected), nil
}

func (repo *Repository) Delete(id string) (int, error) {
	result, err := repo.DB.Exec("DELETE FROM todo WHERE id = $1", id)
	if err != nil {
		return 0, fmt.Errorf("error executing query todo: %s\n", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error getting rows affected: %s\n", err)
	}
	return int(rowsAffected), nil
}
