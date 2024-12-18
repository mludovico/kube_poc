package controllers

import (
	"fmt"
	"net/http"
	"simple_go_app/controllers/interceptors"
)

func initializeHandlers() *http.ServeMux {
	fmt.Println("Initializing handlers...")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos", interceptors.LogInterceptorWrapper(
		interceptors.CorsInterceptorWrapper(GetAllTodos)))
	mux.HandleFunc("GET /todos/{ID}", interceptors.LogInterceptorWrapper(
		interceptors.CorsInterceptorWrapper(GetTodo)))
	mux.HandleFunc("POST /todos", interceptors.LogInterceptorWrapper(
		interceptors.CorsInterceptorWrapper(CreateTodo)))
	mux.HandleFunc("PUT /todos/{ID}", interceptors.LogInterceptorWrapper(
		interceptors.CorsInterceptorWrapper(UpdateTodo)))
	mux.HandleFunc("DELETE /todos/{ID}", interceptors.LogInterceptorWrapper(
		interceptors.CorsInterceptorWrapper(DeleteTodo)))
	mux.HandleFunc("OPTIONS /", interceptors.LogInterceptorWrapper(
		interceptors.CorsInterceptorWrapper(RespondToOptions)))
	return mux
}

func InitializeServer(host, port string) *http.Server {
	handlers := initializeHandlers()
	addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("Initializing server at: %s\n", addr)
	return &http.Server{
		Addr:    addr,
		Handler: handlers,
	}
}
