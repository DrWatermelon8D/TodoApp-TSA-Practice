package main

import (
	"fmt"
	"html/template"
	"net/http"
)

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
	if task != "" {

		newTask := fmt.Sprintf("<p id='greeting'>%s, %s!</p>", "1.", task)

		writer.Write([]byte(newTask))
	} else {
		print("Aloe Vera")
	}

}
