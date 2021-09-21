package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elgael06/curso_1/controllers"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to API REST on Go")
}

func main() {
	var port int = 8000
	fmt.Println("server run on PORT:", port)
	http.HandleFunc("/", index)
	http.HandleFunc("/tasks", controllers.GetAllTask)
	http.HandleFunc("/getFile", controllers.FileRecivedCVS)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
