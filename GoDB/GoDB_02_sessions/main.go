package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"
)

var loginFormTempl = `
	<!DOCTYPE html>	
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>BD</title>
	</head>
	<body>
	<form action="/get_cookie" method="post">
		Login: <input type="text" name="login">
		Password: <input type="password" name="password">
		<input type="submit" value="Login">
	</form>
	</body>
	</html>
`

var sessions = map[string]string{}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		sessionID, err := request.Cookie("session_id")

		if err == http.ErrNoCookie {
			writer.Write([]byte(loginFormTempl))
			return
		} else if err != nil {
			//PanicOnErr(err)
		}
		username, ok := sessions[sessionID.Value]
		if !ok {
			fmt.Fprint(writer, "Session not found!")
		} else {
			fmt.Fprint(writer, "Welcome, "+username)
		}

	})
	http.HandleFunc("/get_cookie", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		inputLogin := request.Form["login"][0]
		pass := request.Form["password"][0]
		expiration := time.Now().Add(365 * 24 * time.Hour)
		sessionID := RandStringRunes(inputLogin+pass+time.Now().String())
		sessions[sessionID] = inputLogin

		cookie := http.Cookie{
			Name:    "session_id",
			Value:   sessionID,
			Expires: expiration,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", http.StatusFound)
	})

	http.ListenAndServe(":8081", nil)
}

func RandStringRunes(login string) (hash string) {
	h:=sha1.New()
	h.Write([]byte(login))
	hash = hex.EncodeToString(h.Sum(nil))
	return
}
