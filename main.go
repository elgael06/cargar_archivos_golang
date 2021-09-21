package main

import (
	"fmt"
	"log"
	"net/http"
)

type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTask []task

var tasks = allTask{
	{
		ID:      1,
		Name:    "task one",
		Content: "some content",
	},
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to my api")
}

func main() {
	var port int = 3030
	fmt.Println("server run on PORT:", port)
	http.HandleFunc("/", index)
	http.HandleFunc("/tasks", getAllTask)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
