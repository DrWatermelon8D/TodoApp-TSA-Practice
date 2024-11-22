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
	writer.Header().Set("Content-Type", "text/html")

	task := req.FormValue("task")

	var location int

	if task != "" {
		for i := 0; i < len(tasks); i++ {
			if tasks[i] == "" {
				tasks[i] = task
				location = i
				break
			}
		}

	} else {
		location = 0
	}

	newTask := fmt.Sprintf("<p id='greeting'>%s, %s!</p>")

	writer.Write([]byte(newTask))

}
