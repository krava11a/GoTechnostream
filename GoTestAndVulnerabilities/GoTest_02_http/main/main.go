package main

import (
	"GoTest_02_http/handlers"
	"fmt"
	"net/http"
)

func main() {
	handler := handlers.Handler{
		Todos: &[]handlers.Todo{
			{"have learned go",false},
		},
	}

	http.HandleFunc("/",handler.HandleIndex)
	http.HandleFunc("/todos/",handler.HandleTodos)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Error: ",err.Error())
	}

}
