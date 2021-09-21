package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elgael06/curso_1/controllers"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to my api")
}

func main() {
	var port int = 3040
	fmt.Println("server run on PORT:", port)
	http.HandleFunc("/", index)
	http.HandleFunc("/tasks", controllers.GetAllTask)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
