package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Handler struct {
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	//Создаем http клиент. В структуру можно передать таймаут, куки и прочую инфу о запросе
	c := http.Client{}
	resp, err := c.Get("http://artii.herokuapp.com/make?text=" + path)
	if err != nil {
		log.Println(err)
	}
	//нужно закрыть тело, когда прочитаем то что нужно
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	//статус OK
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	path := r.URL.Path
	//	//Создаем http клиент. В структуру можно передать таймаут, куки и прочую инфу о запросе
	//	c := http.Client{}
	//	resp, err := c.Get("http://artii.herokuapp.com/make?text=" + path)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	//нужно закрыть тело, когда прочитаем то что нужно
	//	defer resp.Body.Close()
	//
	//	body, _ := ioutil.ReadAll(resp.Body)
	//
	//	//статус OK
	//	w.WriteHeader(http.StatusOK)
	//	w.Write(body)
	//})
	handler := Handler{}

	//http.Handle("/",handler)

	http.ListenAndServe(":8081", handler)

}
