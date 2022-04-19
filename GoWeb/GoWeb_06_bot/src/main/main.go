package main

import (
	"encoding/json"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//для вендоринга используется GB
//сборка проекта осуществляется с помощью gb build
//установка зависимостей - gb vendor fetch gopkg.in/telegram-bot-api.v4
//установка зависимостей из манифеста - gb vendor restore
//Запустить GB не удалось, похоже он умер

type Joke struct {
	ID   uint32 `json:"id"`
	Joke string `json:"joke"`
}

type JokeResponse struct {
	Value Joke   `json:"value"`
	Type  string `json:"type"`
}

var buttons = []tgbotapi.KeyboardButton{
	tgbotapi.KeyboardButton{Text: "Get Joke"},
}

//Этот URL уже мертв
const WebhookURL = "https://msu-go-2017.herokuapp.com"

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
	//Heroku прокидывает порт для приложения в переменную окружения PORT
	port := os.Getenv("PORT")
	bot , err := tgbotapi.NewBotAPI("token to access the HTTP AP")
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	//Install webhook
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe(":"+port, nil)

	//получаем все обновления из канала updates
	for update := range updates{
		var message tgbotapi.MessageConfig
		log.Println("recieved text:", update.Message.Text)

		switch update.Message.Text {
		case "Get Joke":
			//Если нажать на кнопку в боте придет сообщение GetJoke
			message = tgbotapi.NewMessage(update.Message.Chat.ID,getJoke())
		default:
			message = tgbotapi.NewMessage(update.Message.Chat.ID, `Press "GetJoke" to receive joke`)
		}

		//В ответном сообщениии просим показать клавиатуру
		message.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)

		bot.Send(message)
	}




}
