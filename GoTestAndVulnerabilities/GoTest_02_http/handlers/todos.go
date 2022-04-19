package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (h Handler) HandleTodos(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("request", request.URL.Path)
	defer request.Body.Close()

	//обработка разных методов
	switch request.Method {
	//GET
	case http.MethodGet:
		//конвертируем в JSON
		productJson, _ := json.Marshal(h.Todos)
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
		*h.Todos = append(*h.Todos, todo)
	case http.MethodPut:
		id := request.URL.Path[len("/todos/"):]
		index, _ := strconv.ParseInt(id, 10, 0)
		(*h.Todos)[index].Done = true

	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}
