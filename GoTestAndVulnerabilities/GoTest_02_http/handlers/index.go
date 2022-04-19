package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

func (h Handler) HandleIndex(writer http.ResponseWriter, request *http.Request) {
	//Здесь надо отдать статический файл, который будет общаться с API из браузера
	//открываем файл
	fileContents, err := ioutil.ReadFile("../static/index.html")
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	//и выводим сожержимое файла
	writer.Write(fileContents)

}
