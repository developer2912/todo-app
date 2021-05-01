package main

import (
	"log"

	"github.com/developer2912/todo-app"
	"github.com/developer2912/todo-app/pkg/handler"
)

func main() {

	handlers := new(handler.Handler)

	server := new(todo.Server)
	if err := server.Run("8080", handlers.InitRouts()); err != nil {
		log.Fatalf("error occured while runnig http server: %s", err.Error())
	}
}
