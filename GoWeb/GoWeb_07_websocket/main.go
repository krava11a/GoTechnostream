package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Joke struct {
	ID   uint32 `json:"id"`
	Joke string `json:"joke"`
}

type JokeResponse struct {
	Value Joke   `json:"value"`
	Type  string `json:"type"`
}

var upgrader = websocket.Upgrader{

	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//Браузеры за безопасность. Это для того чтобы CORS origin не мучился.
	//ниже нам подходит любой Origin
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Bus struct {
	register  chan *websocket.Conn
	broadcast chan []byte
	clients   map[*websocket.Conn]bool
}

func (b *Bus) Run() {
	for {
		select {
		case message := <-b.broadcast:
			//всем кто зарегался шлем месагу
			for client := range b.clients {
				w, err := client.NextWriter(websocket.TextMessage)
				if err != nil {
					//если достучаться до клиента не удалось, нафиг его
					delete(b.clients, client)
					continue
				}
				w.Write(message)
			}
		case client := <-b.register:
			//регаем клиентов в мапе клиентов
			log.Println("User registered")
			b.clients[client] = true
		}
	}
}

//Constructor
func NewBus() *Bus {
	return &Bus{
		register:  make(chan *websocket.Conn),
		broadcast: make(chan []byte),
		clients:   make(map[*websocket.Conn]bool),
	}
}

func runJoker(b *Bus) {
	for {
		//ходим за шутками каждые 6 секунд
		<-time.After(6 * time.Second)
		log.Println("Its joke time!!")
		b.broadcast <- []byte(getJoke())
	}
}

func getJoke() string {
	c:=http.Client{}
	resp, err := c.Get("http://api.icndb.com/jokes/random?limitTo=[nerdy]")
	if err != nil {
		return "Jokes API not responding"
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	joke := JokeResponse{}

	err = json.Unmarshal(body, &joke)
	if err != nil {
		return "Joke error"
	}
	return joke.Value.Joke
}

func main() {

	bus := NewBus()
	go bus.Run()
	go runJoker(bus)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//upgrade connection
		ws, err := upgrader.Upgrade(writer,request,nil)
		if err!=nil{
			log.Fatal(err)
		}

		bus.register <- ws
	})

	http.ListenAndServe(":8081",nil)
}
