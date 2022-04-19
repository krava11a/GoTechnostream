package main

import (
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

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		sessionID, err := request.Cookie("session_id")

		if err == http.ErrNoCookie {
			writer.Write([]byte(loginFormTempl))
			return
		} else if err != nil {
			//PanicOnErr(err)
		}
		//writer.Header().Set("Content-type","text/plain;charset=UTF-8")
		fmt.Fprint(writer, "Welcome, "+sessionID.Value)
	})
	http.HandleFunc("/get_cookie", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		inputLogin := request.Form["login"][0]
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{
			Name:    "session_id",
			Value:   inputLogin,
			Expires: expiration,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", http.StatusFound)
	})

	http.ListenAndServe(":8081", nil)

}
