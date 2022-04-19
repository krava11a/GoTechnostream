package main

import (
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.New("main")
	tmpl, _ = tmpl.Parse(
		`<div style="display: inline-block;
			border: 1px solid #aaa;
			border-radius: 3px;
			padding:30px;
			margin:20px;">
			{{ if ne . "str"}}
				not str
			{{end}}
			<pre>{{.}}</pre>
			</div>`)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		//Создаем http клиент. В структуру можно передать таймаут, куки и прочую инфу о запросе
		//c := http.Client{}
		//resp, err := c.Get("http://artii.herokuapp.com/make?text=" + path)
		//if err != nil {
		//	log.Println(err)
		//	writer.WriteHeader(http.StatusInternalServerError)
		//	writer.Write([]byte("error"))
		//	return
		//}
		////нужно закрыть тело, когда прочитаем то что нужно
		//defer resp.Body.Close()

		//body, _ := ioutil.ReadAll(resp.Body)

		//статус OK
		tmpl.Execute(writer, path)
	})

	http.ListenAndServe(":8081", nil)
}
