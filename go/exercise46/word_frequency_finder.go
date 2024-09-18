package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var store = make(map[string]int)

func main() {
	server := http.NewServeMux()

	file, err := os.ReadFile("./macbeth.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	words := strings.Fields(string(file))

	for _, word := range words {
		trimmed := strings.Trim(word, ",;:_[]?!@\"'.—”“’)(&*%$#-|I")
		store[strings.ToLower(trimmed)] += 1
	}

	server.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		page := template.Must(template.ParseFiles("index.html"))

		page.Execute(response, struct {
			Title string
			Data  map[string]int
		}{
			Title: "Word Frequency Counter",
			Data:  store,
		})
	})

	http.ListenAndServe(":5000", server)
}
