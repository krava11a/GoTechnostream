package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	Name string
	Done bool
}

func IsNotDone(todo Todo) bool {
	return !todo.Done
}

func main() {
	//создаем шаблон, ппередаем функци
	tmpl, err := template.New("template.html").Funcs(template.FuncMap{"IsNotDone": IsNotDone}).ParseFiles("template.html")
	if err != nil {
		log.Fatal("Can not expand template", err)
		return
	}

	todos := []Todo{
		{"Выучить GO", false},
		{"Посетить лекцию по вебу", false},
		{"...", false},
		{"Profit", false},
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			// читаем из urlencoded запроса
			param := request.FormValue("id")
			index, _ := strconv.ParseInt(param, 10, 0)
			todos[index].Done = true
		}

		err := tmpl.Execute(writer, todos)
		if err != nil {
			http.Error(writer,err.Error(),http.StatusInternalServerError)

		}


	})

	http.ListenAndServe(":8081",nil)

}
