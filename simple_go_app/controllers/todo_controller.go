package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple_go_app/model"
	"simple_go_app/repositories"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	repo := repositories.NewTodoRepository()
	defer repo.DB.Close()
	todos, err := repo.GetAll()
	if err != nil {
		fmt.Printf("error getting all todos: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		fmt.Printf("error getting all todos: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	ID := r.PathValue("ID")
	repo := repositories.NewTodoRepository()
	defer repo.DB.Close()
	todo, err := repo.GetByID(ID)
	if err != nil {
		fmt.Printf("error getting todo: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if todo == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Todo not found"))
		return
	}
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		fmt.Printf("error getting todo: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := &model.Todo{}
	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		fmt.Printf("error creating todo: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}
	repo := repositories.NewTodoRepository()
	defer repo.DB.Close()
	err = repo.Create(todo)
	if err != nil {
		fmt.Printf("error creating todo: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	ID := r.PathValue("ID")
	todo := &model.Todo{}
	todo.ID = ID
	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		fmt.Printf("error updating todo: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}
	repo := repositories.NewTodoRepository()
	defer repo.DB.Close()
	rows, err := repo.Update(todo)
	if err != nil {
		fmt.Printf("error updating todo: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Todo not found"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	ID := r.PathValue("ID")
	repo := repositories.NewTodoRepository()
	defer repo.DB.Close()
	rows, err := repo.Delete(ID)
	if err != nil {
		fmt.Printf("error deleting todo: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Todo not found"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}
