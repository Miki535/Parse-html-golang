package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		url := r.FormValue("url")
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		result := string(body)
		data := struct {
			Result string
		}{
			Result: result,
		}
		tpl.Execute(w, data)
	}
	tpl.Execute(w, nil)
}
