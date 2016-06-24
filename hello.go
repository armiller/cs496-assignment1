package hello

import (
    "net/http"
    "time"
    "html/template"
)

type Hello struct {
	Time string
	Ip string
}

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	response := Hello{
		time.Now().String(),
		r.RemoteAddr,
	}
	t, _ := template.ParseFiles("templates/hello.html")
	t.Execute(w, response)
}
