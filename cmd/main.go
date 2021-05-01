package main

import (
	"log"

	"github.com/developer2912/todo-app"
	"github.com/developer2912/todo-app/pkg/handler"
	"github.com/developer2912/todo-app/pkg/repository"
	"github.com/developer2912/todo-app/pkg/service"
)

func main() {

	repositories := repository.NewRepository()
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run("8080", handlers.InitRouts()); err != nil {
		log.Fatalf("error occured while runnig http server: %s", err.Error())
	}
}
