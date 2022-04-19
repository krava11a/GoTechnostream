package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func main() {
	todos := []Todo{
		{"Выучить GO", false},
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//Здесь надо отдать статический файл, который будет общаться с API из браузера
		//открываем файл
		fileContents, err := ioutil.ReadFile("index.html")
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		//и выводим сожержимое файла
		writer.Write(fileContents)

	})

	http.HandleFunc("/todos/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("request", request.URL.Path)
		defer request.Body.Close()

		//обработка разных методов
		switch request.Method {
		//GET
		case http.MethodGet:
			//конвертируем в JSON
			productJson, _ := json.Marshal(todos)
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write(productJson)
		//POST
		case http.MethodPost:
			decoder := json.NewDecoder(request.Body)
			todo := Todo{}
			//преобразуем json запрос в структуру
			err := decoder.Decode(&todo)
			if err != nil {
				log.Println(err)
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
		case http.MethodPut:
			id := request.URL.Path[len("/todos/"):]
			index, _ := strconv.ParseInt(id, 10, 0)
			todos[index].Done = true

		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8081", nil)
}
