package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tasks [10]string

func main() {
	http.HandleFunc("/", homeHandle)

	http.HandleFunc("/task", taskHandle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func homeHandle(writer http.ResponseWriter, req *http.Request) {

	template := template.Must(template.ParseFiles("index.html"))

	template.Execute(writer, nil)
}

func taskHandle(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "button/html")

	task := req.FormValue("task")

	if task != "" {
		newTask := fmt.Sprintf("<button type=submit>%s</button>", task)

		writer.Write([]byte(newTask))
		writer.Write([]byte("<p></p>"))
	}

}
