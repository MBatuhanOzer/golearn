package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", helloHandler)
	var err error = http.ListenAndServe(":8080", nil)
	fmt.Println("Server started at http://localhost:8080")
	if err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if (r.Method != "GET") {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	temp, err := template.ParseFiles("static/hello.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}